package model

import (
	"time"

	"gorm.io/gorm"
)

type EthTxHashStatus struct {
	gorm.Model
	ChainId   int    `gorm:"int;not null"`
	TxHash    string `gorm:"uniqueIndex;not null"`
	Status    string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
