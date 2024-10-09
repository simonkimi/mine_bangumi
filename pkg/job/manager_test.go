package job

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

const tinyTime = 10 * time.Millisecond

type testJobExec struct {
	started   chan struct{}
	finish    chan struct{}
	t         *testing.T
	cancelled bool
	progress  *Progress
}

func newTestJobExec(t *testing.T, finish chan struct{}) *testJobExec {
	return &testJobExec{
		started: make(chan struct{}),
		finish:  finish,
		t:       t,
	}
}

func (t *testJobExec) Execute(ctx context.Context, progress *Progress) error {
	t.progress = progress
	close(t.started)
	t.t.Logf("executing job")
	if t.finish != nil {
		<-t.finish
		select {
		case <-ctx.Done():
			t.cancelled = true
		default:
		}
	}
	return nil
}

func TestAdd(t *testing.T) {
	req := require.New(t)

	m := NewManager()
	exec1 := newTestJobExec(t, make(chan struct{}))
	job1Id := m.AddJob(context.Background(), exec1)

	req.Equal(1, job1Id)
	time.Sleep(100 * time.Millisecond)

	select {
	case <-exec1.started:
	default:
		t.Error("exec was not started")
	}

	job1 := m.GetJob(job1Id)
	req.NotNil(job1)

	req.Equal(StatusRunning, job1.Status)
	req.NotNil(job1.StartTime)
	req.Nil(job1.EndTime)

	// Job2
	exec2 := newTestJobExec(t, make(chan struct{}))
	job2Id := m.AddJob(context.Background(), exec2)

	req.Equal(2, job2Id)
	time.Sleep(tinyTime)
	job2 := m.GetJob(job2Id)
	req.NotNil(job2)

	req.Equal(StatusReady, job2.Status)
	req.Nil(job2.StartTime)
	req.Nil(job2.EndTime)

	// 完成Job1, Job2开始执行
	close(exec1.finish)

	time.Sleep(tinyTime)
	req.Equal(StatusFinished, job1.Status)
	req.NotNil(job1.EndTime)

	req.Equal(StatusRunning, job2.Status)
	req.NotNil(job2.StartTime)

	select {
	case <-exec2.started:
	default:
		t.Error("exec2 was not started")
	}
}

func TestCancel(t *testing.T) {
	req := require.New(t)
	m := NewManager()

	exec1 := newTestJobExec(t, make(chan struct{}))
	job1Id := m.AddJob(context.Background(), exec1)
	job1 := m.GetJob(job1Id)

	exec2 := newTestJobExec(t, make(chan struct{}))
	job2Id := m.AddJob(context.Background(), exec2)
	job2 := m.GetJob(job2Id)

	// cancel job2
	time.Sleep(tinyTime)
	m.CancelJob(job2Id)

	req.Equal(StatusCancelled, job2.Status)
	req.Nil(job2.EndTime)
	req.Len(m.queue, 1)

	// cancel job1
	m.CancelJob(job1Id)
	time.Sleep(tinyTime)
	req.Equal(StatusStopping, job1.Status)
	req.Len(m.queue, 1)
	close(exec1.finish)
	time.Sleep(tinyTime)
	req.Equal(StatusCancelled, job1.Status)
	req.NotNil(job1.EndTime)
	req.Len(m.queue, 0)
	req.True(exec1.cancelled)
}

func TestListener(t *testing.T) {
	m := NewManager()
	ctx, cancel := context.WithCancel(context.Background())
	notifier := m.Subscribe(ctx)

	exec1 := newTestJobExec(t, make(chan struct{}))
	job1Id := m.AddJob(context.Background(), exec1)

	req := require.New(t)
	req.Len(m.listeners, 1)

	select {
	case newJob := <-notifier.NewJob:
		req.Equal(job1Id, newJob.ID)
		req.Equal(StatusReady, newJob.Status)
	case <-time.After(time.Second):
		t.Error("no new job notification")
	}

	select {
	case newJob := <-notifier.UpdatedJob:
		req.Equal(job1Id, newJob.ID)
		req.Equal(StatusRunning, newJob.Status)
	case <-time.After(time.Second):
		t.Error("no updated job notification")
	}

	select {
	case <-exec1.started:
	case <-time.After(time.Second):
		t.Error("exec was not started")
	}

	exec1.progress.SetPercent(0.1)
	select {
	case updatedJob := <-notifier.UpdatedJob:
		req.Equal(job1Id, updatedJob.ID)
		req.Equal(0.1, updatedJob.Progress)
	case <-time.After(time.Second):
		t.Error("no updated job notification")
	}

	exec1.progress.SetPercent(0.5)
	select {
	case updatedJob := <-notifier.UpdatedJob:
		req.Equal(job1Id, updatedJob.ID)
		req.Equal(0.5, updatedJob.Progress)
	case <-time.After(time.Second):
		t.Error("no updated job notification")
	}

	close(exec1.finish)

	select {
	case removeJob := <-notifier.RemovedJob:
		req.Equal(job1Id, removeJob.ID)
		req.Equal(StatusFinished, removeJob.Status)
	case <-time.After(time.Second):
		t.Error("no removed job notification")
	}

	select {
	case <-notifier.UpdatedJob:
		t.Error("received an additional updatedJob")
	case <-time.After(100 * time.Millisecond):
	}

	// cancel job
	exec2 := newTestJobExec(t, make(chan struct{}))
	job2Id := m.AddJob(context.Background(), exec2)
	m.CancelJob(job2Id)

	select {
	case removeJob := <-notifier.RemovedJob:
		req.Equal(job2Id, removeJob.ID)
		req.Equal(StatusCancelled, removeJob.Status)
	case <-time.After(time.Second):
		t.Error("no removed job notification")
	}

	// cancel subscription
	cancel()
	time.Sleep(tinyTime)
	req.Len(m.listeners, 0)
}
