package controllers

import (
	"net/http"

	"github.com/joker0/renvalmart/internal/app/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// ItemController handles CRUD operations for items
type ItemController struct {
	DB *gorm.DB
}

// CreateItem creates a new item
// @Summary Create a new item
// @Description Create a new item with the given data
// @Accept json
// @Produce json
// @Param input body CreateItemRequest true "Item data"
// @Success 201 {object} Item
// @Router /items [post]
func (ic *ItemController) CreateItem(c echo.Context) error {
	item := new(models.Item)
	if err := c.Bind(item); err != nil {
		return err
	}
	if err := ic.DB.Create(item).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, item)
}

// GetItem retrieves an item by ID
func (ic *ItemController) GetItem(c echo.Context) error {
	id := c.Param("id")
	item := new(models.Item)
	if err := ic.DB.First(item, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	return c.JSON(http.StatusOK, item)
}

// UpdateItem updates an item by ID
func (ic *ItemController) UpdateItem(c echo.Context) error {
	id := c.Param("id")
	item := new(models.Item)
	if err := ic.DB.First(item, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	if err := c.Bind(item); err != nil {
		return err
	}
	if err := ic.DB.Save(item).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, item)
}

// DeleteItem soft-deletes an item by ID
func (ic *ItemController) DeleteItem(c echo.Context) error {
	id := c.Param("id")
	item := new(models.Item)
	if err := ic.DB.First(item, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	if err := ic.DB.Delete(item).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.NoContent(http.StatusNoContent)
}
