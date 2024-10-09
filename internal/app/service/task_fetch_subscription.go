package service

import "context"

type FetchSubscriptionTask struct {
}

func (t *FetchSubscriptionTask) Start(ctx context.Context) error {
	return nil
}

func (t *FetchSubscriptionTask) GetDescription() string {
	return "FetchSubscriptionTask"
}
