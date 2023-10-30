package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joker0/renvalmart/config"
	"github.com/joker0/renvalmart/database"
	_ "github.com/joker0/renvalmart/docs"
	"github.com/joker0/renvalmart/internal/app/models"
	"github.com/joker0/renvalmart/internal/app/routes"
	"github.com/joker0/renvalmart/seeds"
	"github.com/labstack/echo/v4"
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

	// Define your API routes
	routes.RegisterItemRoutes(e, db)
	routes.RegisterSupplierRoutes(e, db)
	routes.RegisterStockRoutes(e, db)
	routes.RegisterUserRoutes(e, db)
	routes.RegisterAuthRoutes(e, db)

	// Start the Echo server
	port := os.Getenv("PORT") // You can set the port using an environment variable
	if port == "" {
		port = "8080" // Default port
	}

	serverAddr := fmt.Sprintf(":%s", port)
	e.Logger.Fatal(e.Start(serverAddr))
}
