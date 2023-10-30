package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/joker0/renvalmart/config"
	"github.com/joker0/renvalmart/database"
	_ "github.com/joker0/renvalmart/docs"
	middlewares "github.com/joker0/renvalmart/internal/app/middlewares"
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

	// Set up middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Define your API routes
	routes.RegisterItemRoutes(e, db)
	routes.RegisterSupplierRoutes(e, db)
	routes.RegisterStockRoutes(e, db)
	routes.RegisterUserRoutes(e, db)
	routes.RegisterAuthRoutes(e, db)

	e.Use(middlewares.AuthMiddleware())

	e.GET("/", func(c echo.Context) error {
		token, ok := c.Get("user").(*jwt.Token) // by default token is stored under `user` key
		if !ok {
			return errors.New("JWT token missing or invalid")
		}
		claims, ok := token.Claims.(jwt.MapClaims) // by default claims is of type `jwt.MapClaims`
		if !ok {
			return errors.New("failed to cast claims as jwt.MapClaims")
		}
		return c.JSON(http.StatusOK, claims)
	})

	// Start the Echo server
	port := os.Getenv("PORT") // You can set the port using an environment variable
	if port == "" {
		port = "8080" // Default port
	}

	serverAddr := fmt.Sprintf(":%s", port)
	e.Logger.Fatal(e.Start(serverAddr))
}
