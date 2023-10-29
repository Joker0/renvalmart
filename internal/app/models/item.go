package models

import (
	"time"

	"gorm.io/gorm"
)

// Item represents the item model in the database
type Item struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	Name         string         `gorm:"type:varchar(255);not null" json:"name"`
	Description  string         `gorm:"type:text" json:"description"`
	SellingPrice int            `json:"selling_price"`
	BuyingPrice  int            `json:"buying_price"`
	Stock        int            `json:"stock"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`
}

// TableName sets the name of the database table for the Item model
func (Item) TableName() string {
	return "items"
}
