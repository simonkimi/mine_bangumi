package job

import (
	"context"
	"time"
)

type Status string

const (
	StatusReady     Status = "READY"
	StatusRunning   Status = "RUNNING"
	StatusStopping  Status = "STOPPING"
	StatusFinished  Status = "FINISHED"
	StatusCancelled Status = "CANCELLED"
	StatusFailed    Status = "FAILED"
)

type Job struct {
	ID        int
	Status    Status
	Progress  float64
	StartTime *time.Time
	EndTime   *time.Time
	errMsg    *string

	exec       Exec
	outerCtx   context.Context
	cancelFunc context.CancelFunc
}

func (j *Job) TimeElapsed() time.Duration {
	if j.EndTime == nil {
		return time.Since(*j.StartTime)
	}
	return j.EndTime.Sub(*j.StartTime)
}

func (j *Job) cancel() {
	if j.Status == StatusReady {
		j.Status = StatusCancelled
	} else if j.Status == StatusRunning {
		j.Status = StatusStopping
	}

	if j.cancelFunc != nil {
		j.cancelFunc()
	}
}

func (m *Manager) onJobFinish(j *Job) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	if j.Status == StatusStopping {
		j.Status = StatusCancelled
	} else if j.Status != StatusFailed {
		j.Status = StatusFinished
	}
	t := time.Now()
	j.EndTime = &t
}

func (m *Manager) onJobFailed(j *Job) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	j.Status = StatusFailed
	t := time.Now()
	j.EndTime = &t
}

func (m *Manager) onJobError(j *Job, err error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	errStr := err.Error()
	j.errMsg = &errStr
	j.Status = StatusFailed
	t := time.Now()
	j.EndTime = &t
}
