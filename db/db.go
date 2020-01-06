package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Use PostgreSQL in gorm
	"github.com/DAdDY0055/go-lottery-app/models"
)

var (
	db  *gorm.DB
	err error
)

func Initialize() {
	connection := "host=0.0.0.0 port=5432 user=postgres password=postgres dbname=postgres sslmode=disable"
	db, err = gorm.Open("postgres", connection)

	db.LogMode(true)

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Task{})
}

func Get() *gorm.DB {
	return db
}

func Close() {
	if err := db.Close(); err != nil {
		panic(err)
  }
}
