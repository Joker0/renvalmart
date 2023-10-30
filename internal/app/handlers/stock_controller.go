package controllers

import (
	"net/http"
	"strconv"

	"github.com/joker0/renvalmart/internal/app/models"
	"github.com/joker0/renvalmart/internal/app/repositories"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// StockController handles CRUD operations for stocks
type StockController struct {
	DB *gorm.DB
}

// CreateStock creates a new stock
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
	if err := repositories.NewStockRepository(sc.DB).CreateStock(stock); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, stock)
}

// GetStock retrieves a stock by ID
// @Summary Get a stock by ID
// @Description Get a stock by its ID
// @Accept json
// @Produce json
// @Param id path int true "Stock ID"
// @Success 200 {object} models.Stock
// @Router /stocks/{id} [get]
func (sc *StockController) GetStock(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	stock, err := repositories.NewStockRepository(sc.DB).GetStockByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	return c.JSON(http.StatusOK, stock)
}

// UpdateStock updates a stock by ID
// @Summary Update a stock by ID
// @Description Update an existing stock by its ID
// @Accept json
// @Produce json
// @Param id path int true "Stock ID"
// @Param input body UpdateStockRequest true "Updated stock data"
// @Success 200 {object} models.Stock
// @Router /stocks/{id] [put]
func (sc *StockController) UpdateStock(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	itemID, _ := strconv.Atoi(c.QueryParam("item_id"))
	supplierID, _ := strconv.Atoi(c.QueryParam("supplier_id"))

	stock, err := repositories.NewStockRepository(sc.DB).GetStockByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	if err := c.Bind(stock); err != nil {
		return err
	}
	if err := repositories.NewStockRepository(sc.DB).UpdateStock(stock, itemID, supplierID); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, stock)
}

// GetStocks retrieves a list of all stocks
// @Summary Get all stocks
// @Description Get a list of all stocks
// @Accept json
// @Produce json
// @Success 200 {array} models.Stock
// @Router /stocks [get]
func (sc *StockController) GetStocks(c echo.Context) error {
	stocks, err := repositories.NewStockRepository(sc.DB).GetStocks()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, stocks)
}

// DeleteStock deletes a stock by ID
// @Summary Delete a stock by ID
// @Description Delete a stock by its ID
// @Accept json
// @Produce json
// @Param id path int true "Stock ID"
// @Success 204 "No Content"
// @Router /stocks/{id} [delete]
func (sc *StockController) DeleteStock(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	stock, err := repositories.NewStockRepository(sc.DB).GetStockByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	if err := repositories.NewStockRepository(sc.DB).DeleteStock(stock); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.NoContent(http.StatusNoContent)
}
