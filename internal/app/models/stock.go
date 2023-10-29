package models

import (
	"time"

	"gorm.io/gorm"
)

// Stock represents the stock model in the database
type Stock struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	InitialStock int            `json:"initial_stock"`
	ItemID       uint           `json:"item_id"`
	SupplierID   uint           `json:"supplier_id"`
	Item         Item           `gorm:"foreignKey:ItemID"`
	Supplier     Supplier       `gorm:"foreignKey:SupplierID"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`
}
