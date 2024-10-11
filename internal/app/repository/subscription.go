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
