package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	var err error
	DB, err = gorm.Open(sqlite.Open("transactions.db"), &gorm.Config{})
	return err
}

func Migration() error {
	return DB.AutoMigrate(
		&model.EthTxHashStatus{}, 
		&model.SolTxHashStatus{}, 
		&model.TonTxHashStatus{},
	)
}

