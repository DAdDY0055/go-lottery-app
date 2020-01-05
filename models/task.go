package models

import "github.com/jinzhu/gorm"

type Task struct {
	gorm.Model
	Text string
}
