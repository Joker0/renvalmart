package routes

import (
	"your-package-name/internal/controllers"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// RegisterItemRoutes registers item-related routes
func RegisterItemRoutes(e *echo.Echo, db *gorm.DB) {
	itemController := controllers.ItemController{DB: db}
	itemRoutes := e.Group("/items")

	itemRoutes.POST("", itemController.CreateItem)
	itemRoutes.GET("/:id", itemController.GetItem)
	itemRoutes.PUT("/:id", itemController.UpdateItem)
	itemRoutes.DELETE("/:id", itemController.DeleteItem)
}
