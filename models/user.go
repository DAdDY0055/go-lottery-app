package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name string
	Department string
	IsWin      string `gorm:"default:'false'"`
	Win        string `gorm:"default:'false'"`
}
