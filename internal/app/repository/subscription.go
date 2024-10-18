package repository

import (
	"github.com/pkg/errors"
	"github.com/simonkimi/minebangumi/internal/app/model"
)

func (r *Repo) AddSubscription(subscription *model.Subscription) error {
	err := r.Db.Create(subscription).Error
	if err != nil {
		return errors.Wrap(err, "add subscription failed")
	}
	return nil
}

func (r *Repo) GetAllSubscriptions() ([]*model.Subscription, error) {
	var subscriptions []*model.Subscription
	err := r.Db.Find(&subscriptions).Error
	if err != nil {
		return nil, errors.Wrap(err, "get all subscriptions failed")
	}
	return subscriptions, nil
}

func (r *Repo) GetAllActiveSubscriptions() ([]*model.Subscription, error) {
	var subscriptions []*model.Subscription
	err := r.Db.Where("is_enabled = ? AND is_ended = ?", true, false).Find(&subscriptions).Error
	if err != nil {
		return nil, errors.Wrap(err, "get all active subscriptions failed")
	}
	return subscriptions, nil
}
