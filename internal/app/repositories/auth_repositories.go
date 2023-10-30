package repositories

import (
	"github.com/joker0/renvalmart/internal/app/models"
	"gorm.io/gorm"
)

type AuthRepository struct {
	DB *gorm.DB // Initialize this with your database connection
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{DB: db}
}

func (ur *AuthRepository) FindUserByEmail(email string) (*models.User, error) {
	user := new(models.User)
	result := ur.DB.Where("email = ?", email).First(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (ur *AuthRepository) FindUserByUsername(username string) (*models.User, error) {
	user := new(models.User)
	result := ur.DB.Where("name = ?", username).First(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (ur *AuthRepository) CreateUser(user *models.User) error {
	result := ur.DB.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
