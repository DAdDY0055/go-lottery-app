package models

import "github.com/jinzhu/gorm"

type Prize struct {
	gorm.Model
	Name string
	WinUserName string
}
