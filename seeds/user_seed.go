package seeds

import (
	"github.com/joker0/renvalmart/internal/app/models"
	"gorm.io/gorm"
)

func SeedUsers(db *gorm.DB) {
	users := []models.User{
		{Name: "Admin User", Password: "adminpassword", Email: "admin@example.com", Role: "admin"},
		{Name: "Owner User", Password: "ownerpassword", Email: "owner@example.com", Role: "owner"},
	}

	for _, user := range users {
		db.Create(&user)
	}
}
