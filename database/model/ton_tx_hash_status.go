package model

import (
	"time"

	"gorm.io/gorm"
)

type TonTxHashStatus struct {
	gorm.Model
	TxHash    string `gorm:"uniqueIndex;not null"`
	Status    string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
