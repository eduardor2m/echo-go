package utils

import (
	"github.com/eduardor2m/echo-go/domain"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDb() *gorm.DB {

	dbPath := "sqlite.db"

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&domain.User{})

	return db

}
