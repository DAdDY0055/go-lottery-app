package db

import (
	"github.com/DAdDY0055/go-lottery-app/models"
	"github.com/jinzhu/gorm"
	"os"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Use PostgreSQL in gorm
)

var (
	db  *gorm.DB
	err error
)

func Initialize() {
	// connection := "host=127.0.0.1 port=15432 user=postgres password=postgres dbname=postgres sslmode=disable"
	connection := os.Getenv("DATABASE_URL")
	db, err = gorm.Open("postgres", connection)

	db.LogMode(true)

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.User{}, &models.Prize{})
}

func Get() *gorm.DB {
	return db
}

func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}
