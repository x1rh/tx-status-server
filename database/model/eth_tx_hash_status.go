package model

import (
	"gorm.io/gorm"
)

type EthTxHashStatus struct {
	gorm.Model
	TxHash    string `gorm:"uniqueIndex;not null"`
	Status    string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

