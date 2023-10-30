package repositories

import (
	"github.com/joker0/renvalmart/internal/app/models"
	"gorm.io/gorm"
)

// StockRepository represents a repository for stock data access
type StockRepository struct {
	DB *gorm.DB // Initialize this with your database connection
}

// NewStockRepository creates a new StockRepository instance
func NewStockRepository(db *gorm.DB) *StockRepository {
	return &StockRepository{DB: db}
}

// CreateStock creates a new stock in the database
func (sr *StockRepository) CreateStock(stock *models.Stock) error {
	result := sr.DB.Create(stock)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetStocks retrieves a list of stocks from the database
func (sr *StockRepository) GetStocks() ([]models.Stock, error) {
	var stocks []models.Stock
	result := sr.DB.Preload("Item").Preload("Supplier").Find(&stocks)
	if result.Error != nil {
		return nil, result.Error
	}
	return stocks, nil
}

// GetStockByID retrieves a stock by its ID
func (sr *StockRepository) GetStockByID(id int) (*models.Stock, error) {
	stock := new(models.Stock)
	result := sr.DB.Where("id = ?", id).Preload("Item").Preload("Supplier").First(stock)
	if result.Error != nil {
		return nil, result.Error
	}
	return stock, nil
}

// UpdateStock updates a stock in the database
func (sr *StockRepository) UpdateStock(stock *models.Stock, itemid, supllierid int) error {
	stock.ItemID = itemid
	stock.SupplierID = supllierid
	result := sr.DB.Save(stock)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// DeleteStock deletes a stock from the database
func (sr *StockRepository) DeleteStock(stock *models.Stock) error {
	result := sr.DB.Delete(stock)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
