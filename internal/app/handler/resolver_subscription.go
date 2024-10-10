package handler

import (
	"context"
	"github.com/simonkimi/minebangumi/api"
	"github.com/simonkimi/minebangumi/internal/app/service"
)

func (q *queryResolver) ParseAcgSubscription(ctx context.Context, input api.ParseAcgSubscriptionInput) (*api.ParseAcgSubscriptionResult, error) {
	client := q.mgr.GetHttpX().GetClient()
	return service.ParseAcgSubscriptionSource(ctx, client, input.URL, input.Source)
}

func (r *mutationResolver) AddAcgSubscription(ctx context.Context, input api.AddSubscriptionInput) (*bool, error) {
	//TODO implement me
	panic("implement me")
}
