package job

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

const tinyTime = 100 * time.Millisecond

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
