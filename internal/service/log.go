package service

import (
	"github.com/imniynaiy/ticket-system/internal/database"
	"github.com/imniynaiy/ticket-system/internal/model"
	"gorm.io/gorm"
)

func getLogRepo() *gorm.DB {
	return database.GlobalDB.Model(&model.Log{})
}
