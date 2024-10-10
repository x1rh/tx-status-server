package model

import (
	"gorm.io/gorm"
)

type TonTxHashStatus struct {
	gorm.Model
	TxHash    string `gorm:"uniqueIndex;not null"` // 交易哈希，唯一索引
	Status    string `gorm:"not null"`             // 交易状态
	CreatedAt time.Time
	UpdatedAt time.Time
}
