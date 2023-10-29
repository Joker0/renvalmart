package models

import (
	"time"

	"gorm.io/gorm"
)

// Supplier represents the supplier model in the database
type Supplier struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	Name          string         `json:"name"`
	ContactPerson string         `json:"contact_person"`
	Phone         string         `json:"phone"`
	Email         string         `json:"email"`
	Address       string         `json:"address"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at"`
}
