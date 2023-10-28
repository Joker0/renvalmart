package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/joker0/renvalmart/config"
	"github.com/joker0/renvalmart/internal/app/routes"
	"github.com/joker0/renvalmart/internal/model"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Initialize Echo framework
	e := echo.New()

	// Load application configurations
	config, err := LoadConfig() // Implement your configuration loading function
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	// Initialize the database connection
	db, err := InitializeDatabase(config)
	if err != nil {
		log.Fatalf("Error initializing the database: %v", err)
	}

	// Set up middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Define your API routes
	routes.RegisterItemRoutes(e, db) // Implement your route initialization function

	// Start the Echo server
	port := os.Getenv("PORT") // You can set the port using an environment variable
	if port == "" {
		port = "8080" // Default port
	}

	serverAddr := fmt.Sprintf(":%s", port)
	e.Logger.Fatal(e.Start(serverAddr))
}

func LoadConfig() (config.DatabaseConfig, error) {
	var config config.DatabaseConfig

	// Load environment variables from the .env file
	if err := godotenv.Load(); err != nil {
		return config, fmt.Errorf("Error loading .env file: %v", err)
	}

	// Read configuration from environment variables
	config.DBPort = os.Getenv("DB_Port")
	config.DBHost = os.Getenv("DB_Host")
	config.DBUsername = os.Getenv("DB_Username")
	config.DBPassword = os.Getenv("DB_Password")
	config.DBName = os.Getenv("DB_Name")

	// If configuration values are missing, you can set defaults or return an error
	if config.DBHost == "" {
		return config, fmt.Errorf("DB_Host is not set in environment variables")
	}

	if config.DBPort == "" {
		config.DBPort = "5432" // Default PostgreSQL port
	}

	if config.DBUsername == "" {
		return config, fmt.Errorf("DB_Username is not set in environment variables")
	}

	if config.DBPassword == "" {
		return config, fmt.Errorf("DB_Password is not set in environment variables")
	}

	if config.DBName == "" {
		return config, fmt.Errorf("DB_Name is not set in environment variables")
	}

	return config, nil
}

func InitializeDatabase(config config.DatabaseConfig) (*gorm.DB, error) {
	// Initialize the database connection
	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	// Migrate the database schemas (run database migrations here)
	db.AutoMigrate(&model.Item{}) // Uncomment and replace with your model

	return db, nil
}
