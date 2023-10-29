package seeds

import (
	"github.com/joker0/renvalmart/internal/app/models"
	"gorm.io/gorm"
)

func SeedItems(db *gorm.DB) {
	items := []models.Item{
		{Name: "Item 1", Description: "Description for Item 1", SellingPrice: 100, BuyingPrice: 80, Stock: 100},
		{Name: "Item 2", Description: "Description for Item 2", SellingPrice: 150, BuyingPrice: 120, Stock: 50},
		{Name: "Item 3", Description: "Description for Item 3", SellingPrice: 200, BuyingPrice: 150, Stock: 75},
		{Name: "Item 4", Description: "Description for Item 4", SellingPrice: 75, BuyingPrice: 60, Stock: 90},
		{Name: "Item 5", Description: "Description for Item 5", SellingPrice: 120, BuyingPrice: 100, Stock: 60},
		{Name: "Item 6", Description: "Description for Item 6", SellingPrice: 80, BuyingPrice: 70, Stock: 110},
		{Name: "Item 7", Description: "Description for Item 7", SellingPrice: 90, BuyingPrice: 80, Stock: 95},
		{Name: "Item 8", Description: "Description for Item 8", SellingPrice: 140, BuyingPrice: 120, Stock: 55},
		{Name: "Item 9", Description: "Description for Item 9", SellingPrice: 110, BuyingPrice: 90, Stock: 70},
		{Name: "Item 10", Description: "Description for Item 10", SellingPrice: 95, BuyingPrice: 80, Stock: 85},
	}

	for _, item := range items {
		db.Create(&item)
	}
}
