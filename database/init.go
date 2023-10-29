package database

import (
	"fmt"

	"github.com/joker0/renvalmart/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeDatabase(config config.DatabaseConfig) (*gorm.DB, error) {
	// Initialize the database connection
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai",
		config.DBHost,
		config.DBPort,
		config.DBUsername,
		config.DBPassword,
		config.DBName)
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
