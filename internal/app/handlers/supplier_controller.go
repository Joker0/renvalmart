package controllers

import (
	"net/http"
	"strconv"

	"github.com/joker0/renvalmart/internal/app/models"
	"github.com/joker0/renvalmart/internal/app/repositories"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// SupplierController handles CRUD operations for suppliers
type SupplierController struct {
	DB *gorm.DB
}

// CreateSupplier creates a new supplier
// @Summary Create a new supplier
// @Description Create a new supplier with the given data
// @Accept json
// @Produce json
// @Param input body CreateSupplierRequest true "Supplier data"
// @Success 201 {object} models.Supplier
// @Router /suppliers [post]
func (sc *SupplierController) CreateSupplier(c echo.Context) error {
	supplier := new(models.Supplier)
	if err := c.Bind(supplier); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}
	if err := repositories.NewSupplierRepository(sc.DB).CreateSupplier(supplier); err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to create supplier")
	}
	return c.JSON(http.StatusCreated, supplier)
}

// GetSupplier retrieves a supplier by its ID
// @Summary Get a supplier by ID
// @Description Get a supplier by its ID
// @Accept json
// @Produce json
// @Param id path int true "Supplier ID"
// @Success 200 {object} models.Supplier
// @Router /suppliers/{id} [get]
func (sc *SupplierController) GetSupplier(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	supplier, err := repositories.NewSupplierRepository(sc.DB).GetSupplierByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Supplier not found")
	}
	return c.JSON(http.StatusOK, supplier)
}

// GetSuppliers retrieves a list of all suppliers
// @Summary Get all suppliers
// @Description Get a list of all suppliers
// @Accept json
// @Produce json
// @Success 200 {array} models.Supplier
// @Router /suppliers [get]
func (sc *SupplierController) GetSuppliers(c echo.Context) error {
	suppliers, err := repositories.NewSupplierRepository(sc.DB).GetSuppliers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to get suppliers")
	}
	return c.JSON(http.StatusOK, suppliers)
}

// UpdateSupplier updates an existing supplier by its ID
// @Summary Update a supplier by ID
// @Description Update an existing supplier by its ID
// @Accept json
// @Produce json
// @Param id path int true "Supplier ID"
// @Param input body UpdateSupplierRequest true "Updated supplier data"
// @Success 200 {object} models.Supplier
// @Router /suppliers/{id] [put]
func (sc *SupplierController) UpdateSupplier(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	supplier, err := repositories.NewSupplierRepository(sc.DB).GetSupplierByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Supplier not found")
	}

	if err := c.Bind(supplier); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	if err := repositories.NewSupplierRepository(sc.DB).UpdateSupplier(supplier); err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to update supplier")
	}
	return c.JSON(http.StatusOK, supplier)
}

// DeleteSupplier deletes a supplier by its ID
// @Summary Delete a supplier by ID
// @Description Delete a supplier by its ID
// @Accept json
// @Produce json
// @Param id path int true "Supplier ID"
// @Success 204 "No Content"
// @Router /suppliers/{id} [delete]
func (sc *SupplierController) DeleteSupplier(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	supplier, err := repositories.NewSupplierRepository(sc.DB).GetSupplierByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Supplier not found")
	}

	if err := repositories.NewSupplierRepository(sc.DB).DeleteSupplier(supplier); err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to delete supplier")
	}

	return c.NoContent(http.StatusNoContent)
}
