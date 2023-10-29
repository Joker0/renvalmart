// controllers/stock_controller.go

package controllers

import (
	"net/http"

	"github.com/joker0/renvalmart/internal/app/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// StockController handles CRUD operations for stocks
type StockController struct {
	DB *gorm.DB
}

// @Summary Create a new stock
// @Description Create a new stock with the given data
// @Accept json
// @Produce json
// @Param input body CreateStockRequest true "Stock data"
// @Success 201 {object} models.Stock
// @Router /stocks [post]
func (sc *StockController) CreateStock(c echo.Context) error {
	stock := new(models.Stock)
	if err := c.Bind(stock); err != nil {
		return err
	}
	if err := sc.DB.Create(stock).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, stock)
}

// @Summary Get a stock by ID
// @Description Get a stock by its ID
// @Accept json
// @Produce json
// @Param id path int true "Stock ID"
// @Success 200 {object} models.Stock
// @Router /stocks/{id} [get]
func (sc *StockController) GetStock(c echo.Context) error {
	stockID := c.Param("id")

	var stock models.Stock

	if err := sc.DB.Preload("Item").Preload("Supplier").Where("id = ?", stockID).First(&stock).Error; err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	return c.JSON(http.StatusOK, stock)
}

// @Summary Update a stock by ID
// @Description Update an existing stock by its ID
// @Accept json
// @Produce json
// @Param id path int true "Stock ID"
// @Param input body UpdateStockRequest true "Updated stock data"
// @Success 200 {object} models.Stock
// @Router /stocks/{id} [put]
func (sc *StockController) UpdateStock(c echo.Context) error {
	id := c.Param("id")
	stock := new(models.Stock)
	if err := sc.DB.First(stock, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	if err := c.Bind(stock); err != nil {
		return err
	}
	if err := sc.DB.Save(stock).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, stock)
}

// @Summary Get all stocks
// @Description Get a list of all stocks
// @Accept json
// @Produce json
// @Success 200 {array} models.Stock
// @Router /stocks [get]
func (sc *StockController) GetStocks(c echo.Context) error {
	stocks := []models.Stock{}
	if err := sc.DB.Preload("Item").Preload("Supplier").Find(&stocks).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, stocks)
}

// @Summary Delete a stock by ID
// @Description Delete a stock by its ID
// @Accept json
// @Produce json
// @Param id path int true "Stock ID"
// @Success 204 "No Content"
// @Router /stocks/{id} [delete]
func (sc *StockController) DeleteStock(c echo.Context) error {
	id := c.Param("id")
	stock := new(models.Stock)
	if err := sc.DB.First(stock, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	if err := sc.DB.Delete(stock).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.NoContent(http.StatusNoContent)
}
