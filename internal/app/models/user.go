package models

import (
	"gorm.io/gorm"
)

// User represents the user model in the database
type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"Not null"`
	Password string `json:"password"`
	Email    string `json:"email" gorm:"unique"`
	Role     string `json:"role"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}
