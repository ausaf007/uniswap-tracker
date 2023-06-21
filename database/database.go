package database

import (
	"github.com/ausaf007/uniswap-tracker/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDatabase(databaseName string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(databaseName), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&models.Pool{}, &models.PoolData{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
