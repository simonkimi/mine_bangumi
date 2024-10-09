package job

import "context"

type simpleJob struct {
	fn ExecFn
}

func (s *simpleJob) Execute(ctx context.Context, progress *Progress) error {
	return s.fn(ctx, progress)
}

func NewSimpleJob(fn ExecFn) Exec {
	return &simpleJob{fn: fn}
}
