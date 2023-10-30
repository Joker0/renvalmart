package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joker0/renvalmart/config"
	"github.com/joker0/renvalmart/database"
	_ "github.com/joker0/renvalmart/docs"
	jwtmiddleware "github.com/joker0/renvalmart/internal/app/middlewares"
	"github.com/joker0/renvalmart/internal/app/models"
	"github.com/joker0/renvalmart/internal/app/routes"
	"github.com/joker0/renvalmart/seeds"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Initialize Echo framework
	e := echo.New()

	// Load application configurations
	config, err := config.LoadConfig() // Implement your configuration loading function
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	// Initialize the database connection
	db, err := database.InitializeDatabase(config)
	if err != nil {
		log.Fatalf("Error initializing the database: %v", err)
	}

	//Migrate the database
	database.Up(db)

	var itemCount int64
	db.Model(&models.Item{}).Count(&itemCount)

	if itemCount != 0 {
		database.Down(db)
	} else {
		//Seed data for the Item, Supplier, and User tables
		seeds.SeedItems(db)
		seeds.SeedSuppliers(db)
		seeds.SeedUsers(db)
		seeds.SeedStocks(db)
	}

	//Set up middleware
	e.Use(middleware.Logger())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: 2 * time.Second,
	}))
	e.Use(middleware.Recover())

	routes.RegisterAuthRoutes(e, db)

	e.Use(jwtmiddleware.AuthMiddleware())

	// Define your API routes
	routes.RegisterItemRoutes(e, db)
	routes.RegisterSupplierRoutes(e, db)
	routes.RegisterStockRoutes(e, db)
	routes.RegisterUserRoutes(e, db)

	// Start the Echo server
	port := os.Getenv("PORT") // You can set the port using an environment variable
	if port == "" {
		port = "8080" // Default port
	}

	serverAddr := fmt.Sprintf(":%s", port)
	e.Logger.Fatal(e.Start(serverAddr))
}
