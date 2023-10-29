package controllers

import (
	"net/http"

	"github.com/joker0/renvalmart/internal/app/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// @title Item API
// @version 1.0
// @description API for managing items

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

// @Summary Get all items
// @Description Get a list of all items
// @Accept json
// @Produce json
// @Success 200 {array} models.Item
// @Router /items [get]
func (ic *ItemController) GetItems(c echo.Context) error {
	items := []models.Item{}
	if err := ic.DB.Find(&items).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, items)
}

// @Summary Get an item by ID
// @Description Get an item by its ID
// @Accept json
// @Produce json
// @Param id path int true "Item ID"
// @Success 200 {object} models.Item
// @Router /items/{id} [get]
func (ic *ItemController) GetItem(c echo.Context) error {
	id := c.Param("id")
	item := new(models.Item)
	if err := ic.DB.First(item, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	return c.JSON(http.StatusOK, item)
}

// @Summary Update an item by ID
// @Description Update an existing item by its ID
// @Accept json
// @Produce json
// @Param id path int true "Item ID"
// @Param input body UpdateItemRequest true "Updated item data"
// @Success 200 {object} models.Item
// @Router /items/{id} [put]
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

// @Summary Delete an item by ID
// @Description Delete an item by its ID
// @Accept json
// @Produce json
// @Param id path int true "Item ID"
// @Success 204 "No Content"
// @Router /items/{id} [delete]
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
