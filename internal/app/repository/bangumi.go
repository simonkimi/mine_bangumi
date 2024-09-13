package repository

import (
	"github.com/simonkimi/minebangumi/internal/app/model"
	"github.com/simonkimi/minebangumi/internal/pkg/database"
)

func InsertBangumi(model *model.Bangumi) error {
	return database.Db.Create(model).Error
}
