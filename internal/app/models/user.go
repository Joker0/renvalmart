package models

import (
	"gorm.io/gorm"
)

// User represents the user model in the database
type User struct {
	gorm.Model
	Name     string `json:"name"`
	Password string `json:"-"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}
