package models

import "github.com/jinzhu/gorm"

type SessionInfo struct {
	gorm.Model
	AdminUserName  interface{}
	IsSessionAlive bool
}
