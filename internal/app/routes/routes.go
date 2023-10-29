package routes

import (
	controllers "github.com/joker0/renvalmart/internal/app/handlers"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// RegisterItemRoutes registers item-related routes
func RegisterItemRoutes(e *echo.Echo, db *gorm.DB) {
	itemController := controllers.ItemController{DB: db}
	itemRoutes := e.Group("/items")

	itemRoutes.POST("", itemController.CreateItem)
	itemRoutes.GET("", itemController.GetItems)
	itemRoutes.GET("/:id", itemController.GetItem)
	itemRoutes.PUT("/:id", itemController.UpdateItem)
	itemRoutes.DELETE("/:id", itemController.DeleteItem)
}

// RegisterSupplierRoutes registers supplier-related routes
func RegisterSupplierRoutes(e *echo.Echo, db *gorm.DB) {
	supplierController := controllers.SupplierController{DB: db}
	supplierRoutes := e.Group("/suppliers")

	supplierRoutes.POST("", supplierController.CreateSupplier)
	supplierRoutes.GET("", supplierController.GetSuppliers)
	supplierRoutes.GET("/:id", supplierController.GetSupplier)
	supplierRoutes.PUT("/:id", supplierController.UpdateSupplier)
	supplierRoutes.DELETE("/:id", supplierController.DeleteSupplier)
}

// RegisterStockRoutes registers stock-related routes
func RegisterStockRoutes(e *echo.Echo, db *gorm.DB) {
	stockController := controllers.StockController{DB: db}
	stockRoutes := e.Group("/stocks")

	stockRoutes.POST("", stockController.CreateStock)
	stockRoutes.GET("", stockController.GetStocks)
	stockRoutes.GET("/:id", stockController.GetStock)
	stockRoutes.PUT("/:id", stockController.UpdateStock)
	stockRoutes.DELETE("/:id", stockController.DeleteStock)
}

func RegisterUserRoutes(e *echo.Echo, db *gorm.DB) {
	userController := controllers.UserController{DB: db}
	userRoutes := e.Group("/users")

	//userRoutes.Use(middleware.AuthMiddleware())

	userRoutes.POST("", userController.CreateUser)
	userRoutes.GET("", userController.GetUsers)
	userRoutes.GET("/:id", userController.GetUser)
	userRoutes.PUT("/:id", userController.UpdateUser)
	userRoutes.DELETE("/:id", userController.DeleteUser)
}

// func RegisterAuthRoutes(e *echo.Echo, db *gorm.DB) {
// 	authRoutes := e.Group("/auth")
// 	authRoutes.POST("/login", middlewares.AuthMiddleware, authController.Login)
// 	authRoutes.POST("/register", authController.Register)
// 	// Other authentication-related routes
// }
