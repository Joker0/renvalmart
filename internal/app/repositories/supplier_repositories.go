package repositories

import (
	"github.com/joker0/renvalmart/internal/app/models"
	"gorm.io/gorm"
)

// SupplierRepository represents a repository for supplier data access
type SupplierRepository struct {
	DB *gorm.DB // Initialize this with your database connection
}

// NewSupplierRepository creates a new SupplierRepository instance
func NewSupplierRepository(db *gorm.DB) *SupplierRepository {
	return &SupplierRepository{DB: db}
}

// CreateSupplier creates a new supplier in the database
func (sr *SupplierRepository) CreateSupplier(supplier *models.Supplier) error {
	result := sr.DB.Create(supplier)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetSuppliers retrieves a list of suppliers from the database
func (sr *SupplierRepository) GetSuppliers() ([]models.Supplier, error) {
	var suppliers []models.Supplier
	result := sr.DB.Find(&suppliers)
	if result.Error != nil {
		return nil, result.Error
	}
	return suppliers, nil
}

// GetSupplierByID retrieves a supplier by its ID
func (sr *SupplierRepository) GetSupplierByID(id int) (*models.Supplier, error) {
	supplier := new(models.Supplier)
	result := sr.DB.Where("id = ?", id).First(supplier)
	if result.Error != nil {
		return nil, result.Error
	}
	return supplier, nil
}

// UpdateSupplier updates a supplier in the database
func (sr *SupplierRepository) UpdateSupplier(supplier *models.Supplier) error {
	result := sr.DB.Save(supplier)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// DeleteSupplier deletes a supplier from the database
func (sr *SupplierRepository) DeleteSupplier(supplier *models.Supplier) error {
	result := sr.DB.Delete(supplier)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
