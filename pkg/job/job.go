package job

import (
	"context"
	"sync"
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

	mux sync.Mutex

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

func (j *Job) onJobStart(cancel context.CancelFunc) {
	j.mux.Lock()
	defer j.mux.Unlock()
	t := time.Now()
	j.StartTime = &t
	j.Status = StatusRunning
	j.cancelFunc = cancel
}

func (j *Job) onJobFinish() {
	j.mux.Lock()
	defer j.mux.Unlock()
	if j.Status == StatusStopping {
		j.Status = StatusCancelled
	} else if j.Status != StatusFailed {
		j.Status = StatusFinished
	}
	t := time.Now()
	j.EndTime = &t
}

func (j *Job) onJobFailed() {
	j.mux.Lock()
	defer j.mux.Unlock()
	j.Status = StatusFailed
	t := time.Now()
	j.EndTime = &t
}

func (j *Job) onJobError(err error) {
	j.mux.Lock()
	defer j.mux.Unlock()
	errStr := err.Error()
	j.errMsg = &errStr
	j.Status = StatusFailed
	t := time.Now()
	j.EndTime = &t
}
