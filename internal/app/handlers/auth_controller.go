package controllers

import (
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joker0/renvalmart/internal/app/models"
	"github.com/joker0/renvalmart/internal/app/repositories"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/labstack/echo/v4"

	"time"
)

// TokenClaims struct for JWT claims
type TokenClaims struct {
	ID   uint   `json:"id"`
	Role string `json:"role"`
	Name string `json:"name"`
	jwt.RegisteredClaims
}

// AuthController handles authentication
type AuthController struct {
	DB *gorm.DB
}

// Register handles user registration
func (ac *AuthController) Register(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	// Hash the user's password before storing it
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to hash the password")
	}
	user.Password = string(hashedPassword)

	// Save the user in your database
	if err := repositories.NewAuthRepository(ac.DB).CreateUser(user); err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to register user")
	}

	return c.JSON(http.StatusCreated, user)
}

// Login handles user login and generates a JWT token
func (ac *AuthController) Login(c echo.Context) error {
	loginRequest := new(models.LoginRequest)
	if err := c.Bind(loginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	user, err := repositories.NewAuthRepository(ac.DB).FindUserByUsername(loginRequest.Username)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "User not found")
	}

	// Compare the provided password with the stored hash
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "Invalid password")
	}

	// Generate a JWT token
	token := generateToken(user)

	// Return the token and user information
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Successfully login",
		"token":   token,
		"payload": user,
	})
}

// generateToken generates a JWT token for a user
func generateToken(user *models.User) string {
	claims := TokenClaims{
		ID:   user.ID,
		Role: user.Role,
		Name: user.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "Admin",
			Subject:   user.Name,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_Secret_Key")))
	if err != nil {
		panic(err)
	}

	return tokenString
}
