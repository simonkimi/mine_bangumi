package repository

import (
	"github.com/simonkimi/minebangumi/internal/app/database"
	"github.com/simonkimi/minebangumi/internal/app/model"
)

func InsertBangumi(model *model.Bangumi) error {
	return database.Db.Create(model).Error
}
