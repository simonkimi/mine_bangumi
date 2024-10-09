package job

import "context"

type ExecFn func(ctx context.Context, progress *Progress) error

type Exec interface {
	Execute(ctx context.Context, progress *Progress) error
}
