package models

import (
	"gorm.io/gorm"
)

// Stock represents the stock model in the database
type Stock struct {
	gorm.Model
	InitialStock int      `json:"initial_stock"`
	ItemID       uint     `json:"item_id"`
	SupplierID   uint     `json:"supplier_id"`
	Item         Item     `json:"item" gorm:"foreignKey:ItemID"`
	Supplier     Supplier `json:"supplier" gorm:"foreignKey:SupplierID"`
}
