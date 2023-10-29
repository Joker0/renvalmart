package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type DatabaseConfig struct {
	DBHost     string
	DBPort     string
	DBUsername string
	DBPassword string
	DBName     string
	// Add other database configuration options here
}

func LoadConfig() (DatabaseConfig, error) {
	var config DatabaseConfig

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
