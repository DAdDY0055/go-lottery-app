package db

import (
	"github.com/DAdDY0055/go-gin-gorm-todo-app/models"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

var db *gorm.DB

func Initialize() {
	db, _ = gorm.Open("sqlite3", "task.db")

	db.LogMode(true)

	db.AutoMigrate(&models.Task{})
}

func Get() *gorm.DB {
	return db
}

func Close() {
	db.Close()
}
