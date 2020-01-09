package models

type User struct {
	ID   int
	Name string
	Win  string `gorm:"default:'未済'"`
}
