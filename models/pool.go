package models

import "gorm.io/gorm"

type Pool struct {
	gorm.Model
	PoolAddress string     `gorm:"type:text;"`
	PoolData    []PoolData `gorm:"foreignKey:PoolID"`
}
