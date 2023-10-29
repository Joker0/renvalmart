package models

import (
	"gorm.io/gorm"
)

// Supplier represents the supplier model in the database
type Supplier struct {
	gorm.Model
	Name          string `json:"name"`
	ContactPerson string `json:"contact_person"`
	Phone         string `json:"phone"`
	Email         string `json:"email"`
	Address       string `json:"address"`
}
