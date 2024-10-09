package service

import "context"

type Tasks interface {
	Start(ctx context.Context) error
	GetDescription() string
}
