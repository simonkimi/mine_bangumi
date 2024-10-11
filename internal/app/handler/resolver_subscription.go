package handler

import (
	"context"
	"github.com/simonkimi/minebangumi/api"
	"github.com/simonkimi/minebangumi/internal/app/model"
	"github.com/simonkimi/minebangumi/internal/app/service"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func (q *queryResolver) ParseAcgSubscription(ctx context.Context, input api.ParseAcgSubscriptionInput) (*api.ParseAcgSubscriptionResult, error) {
	client := q.mgr.GetHttpX().GetClient()
	return service.ParseAcgSubscriptionSource(ctx, client, input.URL, input.Source)
}

func (r *mutationResolver) AddAcgSubscription(_ context.Context, input api.AddSubscriptionInput) (*bool, error) {
	repo := r.mgr.GetRepo()
	err := repo.AddSubscription(&model.Subscription{
		Model:           gorm.Model{},
		Link:            input.URL,
		IsAggregate:     false,
		Source:          input.Source.String(),
		BlackListFilter: input.BlackListFilter,
		WhiteListFilter: input.WhiteListFilter,
		IsEnabled:       true,
	})
	if err != nil {
		logrus.WithError(err).Error("AddAcgSubscription")
		return nil, api.NewInternalServerError(err)
	}
	result := true
	return &result, nil
}
