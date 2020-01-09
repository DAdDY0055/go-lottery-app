package models

import "github.com/jinzhu/gorm"

type Prize struct {
	gorm.Model
	Name string
	Winner string
}
