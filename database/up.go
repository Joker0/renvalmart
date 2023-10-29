package database

import (
	"github.com/joker0/renvalmart/internal/app/models"
	"gorm.io/gorm"
)

func Up(db *gorm.DB) {
	Migrateerr := db.AutoMigrate(&models.Item{}, &models.Stock{}, &models.Supplier{}, &models.User{})
	if Migrateerr != nil {
		panic("Migration Failed")
	}
}
