package main

import (
	"github.com/DAdDY0055/go-lottery-app/db"
	"github.com/DAdDY0055/go-lottery-app/router"
)

func main() {
	db.Initialize()
	defer db.Close()

	router.Router()
}
