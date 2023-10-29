package database

import (
	"github.com/joker0/renvalmart/internal/app/models"
	"gorm.io/gorm"
)

func Down(db *gorm.DB) {
	db.Migrator().DropTable(&models.Item{}, &models.Stock{}, &models.Supplier{}, &models.User{})
}
