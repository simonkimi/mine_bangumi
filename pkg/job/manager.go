package job

import (
	"context"
	"github.com/simonkimi/minebangumi/tools/xarray"
	"github.com/sirupsen/logrus"
	"runtime/debug"
	"sync"
)

type Manager struct {
	queue []*Job
	jobId int

	mutex    sync.Mutex
	notEmpty *sync.Cond

	listeners []*ChangeNotifier

	stop chan struct{}
}

func NewManager() *Manager {
	m := &Manager{
		stop: make(chan struct{}),
	}
	m.notEmpty = sync.NewCond(&m.mutex)
	go m.loop()
	return m
}

func (m *Manager) Subscribe(ctx context.Context) *ChangeNotifier {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	notifier := newChangeNotifier()
	m.listeners = append(m.listeners, notifier)
	go func() {
		<-ctx.Done()
		m.mutex.Lock()
		defer m.mutex.Unlock()
		m.listeners, _ = xarray.RemoveFirst(m.listeners, notifier)
		notifier.close()
	}()

	return notifier
}

func (m *Manager) AddJob(ctx context.Context, exec Exec) int {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.jobId += 1
	j := &Job{
		ID:       m.jobId,
		Status:   StatusReady,
		outerCtx: ctx,
		exec:     exec,
	}

	m.queue = append(m.queue, j)
	if len(m.queue) == 1 {
		m.notEmpty.Broadcast()
	}
	return j.ID
}

func (m *Manager) CancelJob(id int) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	j := m.getJob(id)
	if j == nil {
		return
	}
	j.cancel()
	if j.Status == StatusCancelled {
		m.removeJob(j)
	}
}

func (m *Manager) GetJob(id int) *Job {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	for _, j := range m.queue {
		if j.ID == id {
			return j
		}
	}
	return nil
}

func (m *Manager) getJob(id int) *Job {
	for _, j := range m.queue {
		if j.ID == id {
			return j
		}
	}
	return nil
}

func (m *Manager) getReadyJob() *Job {
	for _, j := range m.queue {
		if j.Status == StatusReady {
			return j
		}
	}

	return nil

}

func (m *Manager) loop() {
	m.mutex.Lock()
	for {
		j := m.getReadyJob()
		if j == nil {
			m.notEmpty.Wait()
			select {
			case <-m.stop:
				m.mutex.Unlock()
				return
			default:
				continue
			}
		}
		done := m.startJob(j.outerCtx, j)
		m.mutex.Unlock()
		<-done
		m.mutex.Lock()
		m.removeJob(j)
	}
}

func (m *Manager) startJob(ctx context.Context, j *Job) chan struct{} {
	ctx, cancel := context.WithCancel(ctx)
	j.onJobStart(cancel)
	done := make(chan struct{})
	go m.execJob(ctx, j, done)
	return done
}

func (m *Manager) execJob(ctx context.Context, j *Job, done chan struct{}) {
	defer close(done)
	defer j.onJobFinish()
	defer func() {
		if r := recover(); r != nil {
			logrus.Errorf("panic in job %d: %v", j.ID, r)
			logrus.Error(string(debug.Stack()))
			j.onJobFailed()
		}
	}()

	progress := m.newProgress(j)
	err := j.exec.Execute(ctx, progress)
	if err != nil {
		j.onJobError(err)
		logrus.WithError(err).Errorf("task failed")
	}
}

func (m *Manager) removeJob(j *Job) {
	queue, exist := xarray.RemoveFirst(m.queue, j)
	if !exist {
		return
	}
	m.queue = queue
	m.notifyJobRemove(j)
}
