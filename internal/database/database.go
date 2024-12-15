package database

import (
	"fmt"

	"github.com/imniynaiy/ticket-system/internal/config"
	"github.com/imniynaiy/ticket-system/internal/log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var GlobalDB *gorm.DB

func InitDB(dc *config.DatabaseConfig) {
	// "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		dc.Username, dc.Password, dc.Address, dc.Port, dc.DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Init DB failed", log.String("err", err.Error()))
	}
	GlobalDB = db
}

func CloseDB() {
	sqlDB, err := GlobalDB.DB()
	if err != nil {
		log.Error("Failed to get sqlDB", log.String("err", err.Error()))
	}
	sqlDB.Close()
}
