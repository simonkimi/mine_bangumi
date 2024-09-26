package repository

import (
	"github.com/simonkimi/minebangumi/internal/app/model"
)

func (r *Repo) InsertBangumi(model *model.Bangumi) error {
	return r.Db.Create(model).Error
}
