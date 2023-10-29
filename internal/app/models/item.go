package models

import (
	"gorm.io/gorm"
)

// Item represents the item model in the database
type Item struct {
	gorm.Model
	Name         string `gorm:"type:varchar(255);not null" json:"name"`
	Description  string `gorm:"type:text" json:"description"`
	SellingPrice int    `json:"selling_price"`
	BuyingPrice  int    `json:"buying_price"`
	Stock        int    `json:"stock"`
}
