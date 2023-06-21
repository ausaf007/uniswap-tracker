package models

import "gorm.io/gorm"

type PoolData struct {
	gorm.Model
	PoolID        uint   `gorm:"uniqueIndex:idx_pool_block"`
	BlockNumber   int64  `gorm:"uniqueIndex:idx_pool_block;type:integer;"`
	Token0Balance string `gorm:"type:text;"`
	Token1Balance string `gorm:"type:text;"`
	Tick          string `gorm:"type:text;"`
	Token0Delta   string `gorm:"type:text;"`
	Token1Delta   string `gorm:"type:text;"`
}
