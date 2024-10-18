package service

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/simonkimi/minebangumi/api"
	"github.com/simonkimi/minebangumi/internal/app/model"
	"github.com/simonkimi/minebangumi/pkg/job"
	"github.com/simonkimi/minebangumi/pkg/mikan"
	"golang.org/x/sync/semaphore"
	"sync"
)

type FetchSubscriptionTask struct {
	mgr Manager
}

func (f *FetchSubscriptionTask) Execute(ctx context.Context, progress *job.Progress) error {
	activeSubscriptions, err := f.mgr.GetRepo().GetAllActiveSubscriptions()
	if err != nil {
		return err
	}
	if len(activeSubscriptions) == 0 {
		return nil
	}
	progress.SetTotal(len(activeSubscriptions))

	var wg sync.WaitGroup
	sem := semaphore.NewWeighted(5)

	for _, subscription := range activeSubscriptions {
		wg.Add(1)
		go func(sub *model.Subscription) {
			defer wg.Done()
			if err := sem.Acquire(ctx, 1); err != nil {
				progress.NextStep()
				return
			}
			defer sem.Release(1)
			if err := f.refreshSubscription(ctx, sub); err != nil {
				progress.NextStepWithError(err)
				return
			}
			progress.NextStep()
		}(subscription)
	}
	wg.Wait()
	return nil
}

func (f *FetchSubscriptionTask) refreshSubscription(ctx context.Context, sub *model.Subscription) error {
	switch sub.Source {
	case api.SourceEnumMikan.String():
		return f.refreshMikanSubscription(ctx, sub)
	default:
		return errors.New(fmt.Sprintf(""))
	}
}

func (f *FetchSubscriptionTask) refreshMikanSubscription(ctx context.Context, sub *model.Subscription) error {
	bangumi, err := mikan.ParseUrl(ctx, f.mgr.GetHttpX().GetClient(), sub.Link)
	if err != nil {
		return err
	}

}
