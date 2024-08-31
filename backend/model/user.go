package model

import "gorm.io/gorm"

// User struct to model user data
type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"unique"`
	Password string `json:"password"`
	Notes    []Note `json:"notes"`
}
