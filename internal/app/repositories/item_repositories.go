package repositories

import (
	"github.com/joker0/renvalmart/internal/app/models"

	"gorm.io/gorm"
)

// ItemRepository represents a repository for item data access
type ItemRepository struct {
	DB *gorm.DB // Initialize this with your database connection
}

// NewItemRepository creates a new ItemRepository instance
func NewItemRepository(db *gorm.DB) *ItemRepository {
	return &ItemRepository{DB: db}
}

// CreateItem creates a new item in the database
func (ir *ItemRepository) CreateItem(item *models.Item) error {
	result := ir.DB.Create(item)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetItems retrieves a list of items from the database
func (ir *ItemRepository) GetItems() ([]models.Item, error) {
	var items []models.Item
	result := ir.DB.Find(&items)
	if result.Error != nil {
		return nil, result.Error
	}
	return items, nil
}

// GetItemByID retrieves an item by its ID
func (ir *ItemRepository) GetItemByID(id int) (*models.Item, error) {
	item := new(models.Item)
	result := ir.DB.Where("id = ?", id).First(item)
	if result.Error != nil {
		return nil, result.Error
	}
	return item, nil
}

// UpdateItem updates an item in the database
func (ir *ItemRepository) UpdateItem(item *models.Item) error {
	result := ir.DB.Save(item)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// DeleteItem deletes an item from the database
func (ir *ItemRepository) DeleteItem(item *models.Item) error {
	result := ir.DB.Delete(item)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
