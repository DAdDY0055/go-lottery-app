package models

import "github.com/jinzhu/gorm"

type AdminUser struct {
	gorm.Model
	Name string
	Password string
}
