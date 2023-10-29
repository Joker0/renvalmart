package seeds

import (
	"math/rand"

	"github.com/joker0/renvalmart/internal/app/models"
	"gorm.io/gorm"
)

func SeedStocks(db *gorm.DB) {
	// Seed 10 stock items for each of the first 10 items with random suppliers
	for itemID := 1; itemID <= 10; itemID++ {
		// Randomly select a supplier ID between 1 and 10 (adjust as needed)
		supplierID := uint(rand.Intn(10) + 1)
		initialStock := rand.Intn(100) + 1 // Random initial stock (1 to 100)

		stock := models.Stock{
			InitialStock: initialStock,
			ItemID:       uint(itemID),
			SupplierID:   supplierID,
		}

		db.Create(&stock)
	}
}
