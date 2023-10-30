package repositories

import (
	"github.com/joker0/renvalmart/internal/app/models"
	"gorm.io/gorm"
)

// UserRepository represents a repository for user data access
type UserRepository struct {
	DB *gorm.DB // Initialize this with your database connection
}

// NewUserRepository creates a new UserRepository instance
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// CreateUser creates a new user in the database
func (ur *UserRepository) CreateUser(user *models.User) error {
	result := ur.DB.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetUsers retrieves all users from the database
func (ur *UserRepository) GetUsers() ([]models.User, error) {
	users := []models.User{}
	result := ur.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

// FindUserByID finds a user by their ID
func (ur *UserRepository) FindUserByID(id int) (*models.User, error) {
	user := new(models.User)
	result := ur.DB.Where("id = ?", id).First(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

// UpdateUser updates a user in the database
func (ur *UserRepository) UpdateUser(user *models.User) error {
	result := ur.DB.Save(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// DeleteUser deletes a user from the database
func (ur *UserRepository) DeleteUser(user *models.User) error {
	result := ur.DB.Delete(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
