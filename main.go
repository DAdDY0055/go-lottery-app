package main

import (
	"github.com/DAdDY0055/go-gin-gorm-todo-app/db"
	"github.com/DAdDY0055/go-gin-gorm-todo-app/router"
)

func main() {
	db.Initialize()
	defer db.Close()

	router.Router()
}
