package controllers

import (
	"net/http"

	"github.com/joker0/renvalmart/internal/app/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// @title Supplier API
// @version 1.0
// @description API for managing suppliers

// SupplierController handles CRUD operations for suppliers
type SupplierController struct {
	DB *gorm.DB
}

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
		return err
	}
	if err := sc.DB.Create(supplier).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, supplier)
}

// @Summary Get a supplier by ID
// @Description Get a supplier by its ID
// @Accept json
// @Produce json
// @Param id path int true "Supplier ID"
// @Success 200 {object} models.Supplier
// @Router /suppliers/{id} [get]
func (sc *SupplierController) GetSupplier(c echo.Context) error {
	id := c.Param("id")
	supplier := new(models.Supplier)
	if err := sc.DB.First(supplier, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	return c.JSON(http.StatusOK, supplier)
}

// @Summary Get all suppliers
// @Description Get a list of all suppliers
// @Accept json
// @Produce json
// @Success 200 {array} models.Supplier
// @Router /suppliers [get]
func (sc *SupplierController) GetSuppliers(c echo.Context) error {
	suppliers := []models.Supplier{}
	if err := sc.DB.Find(&suppliers).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, suppliers)
}

// @Summary Update a supplier by ID
// @Description Update an existing supplier by its ID
// @Accept json
// @Produce json
// @Param id path int true "Supplier ID"
// @Param input body UpdateSupplierRequest true "Updated supplier data"
// @Success 200 {object} models.Supplier
// @Router /suppliers/{id] [put]
func (sc *SupplierController) UpdateSupplier(c echo.Context) error {
	id := c.Param("id")
	supplier := new(models.Supplier)
	if err := sc.DB.First(supplier, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	if err := c.Bind(supplier); err != nil {
		return err
	}
	if err := sc.DB.Save(supplier).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, supplier)
}

// @Summary Delete a supplier by ID
// @Description Delete a supplier by its ID
// @Accept json
// @Produce json
// @Param id path int true "Supplier ID"
// @Success 204 "No Content"
// @Router /suppliers/{id} [delete]
func (sc *SupplierController) DeleteSupplier(c echo.Context) error {
	id := c.Param("id")
	supplier := new(models.Supplier)
	if err := sc.DB.First(supplier, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	if err := sc.DB.Delete(supplier).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.NoContent(http.StatusNoContent)
}
