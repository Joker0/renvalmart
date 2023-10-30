package controllers

import (
	"net/http"
	"strconv"

	"github.com/joker0/renvalmart/internal/app/models"
	"github.com/joker0/renvalmart/internal/app/repositories"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// UserController handles CRUD operations for users
type UserController struct {
	DB *gorm.DB
}

// @Summary Create a new user
// @Description Create a new user with the given data
// @Accept json
// @Produce json
// @Param input body CreateUserRequest true "User data"
// @Success 201 {object} models.User
// @Router /users [post]
func (uc *UserController) CreateUser(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return err
	}
	if err := repositories.NewUserRepository(uc.DB).CreateUser(user); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, models.UserResponse{
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	})
}

// @Summary Get all users
// @Description Get a list of all users
// @Accept json
// @Produce json
// @Success 200 {array} models.User
// @Router /users [get]
func (uc *UserController) GetUsers(c echo.Context) error {
	users, err := repositories.NewUserRepository(uc.DB).GetUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, users)
}

// @Summary Get a user by ID
// @Description Get a user by its ID
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.User
// @Router /users/{id} [get]
func (uc *UserController) GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := repositories.NewUserRepository(uc.DB).FindUserByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	return c.JSON(http.StatusOK, user)
}

// @Summary Update a user by ID
// @Description Update an existing user by its ID
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param input body UpdateUserRequest true "Updated user data"
// @Success 200 {object} models.User
// @Router /users/{id} [put]
func (uc *UserController) UpdateUser(c echo.Context) error {
	id := c.Param("id")
	user := new(models.User)
	if err := uc.DB.First(user, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	if err := c.Bind(user); err != nil {
		return err
	}
	if err := repositories.NewUserRepository(uc.DB).UpdateUser(user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, user)
}

// @Summary Delete a user by ID
// @Description Delete a user by its ID
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 204 "No Content"
// @Router /users/{id} [delete]
func (uc *UserController) DeleteUser(c echo.Context) error {
	id := c.Param("id")
	user := new(models.User)
	if err := uc.DB.First(user, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	if err := repositories.NewUserRepository(uc.DB).DeleteUser(user); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.NoContent(http.StatusNoContent)
}
