package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string   `gorm:"not null; unique; index"`
	Email    string   `gorm:"not null; unique; index"`
	Password string   `gorm:"not null"`
	Trackrs  []Trackr `gorm:"foreignKey:UserID;"`
}
