package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Password string
	Todos    []Todo
}

type Todo struct {
	gorm.Model
	Title     string
	Completed bool `gorm:"default:false"`
	UserID    uint
}
