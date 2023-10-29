package controllers

// import (
//     "net/http"

// 	"github.com/joker0/renvalmart/internal/app/models"

// 	"github.com/labstack/echo-jwt/v4"

//     "github.com/labstack/echo/v4"

//     "time"
// )

// // AuthController handles user authentication and token generation
// type AuthController struct {
//     SecretKey []byte // Secret key for JWT token signing
// }

// // @Summary Login and generate JWT token
// // @Description Authenticate the user and return a JWT token
// // @Accept json
// // @Produce json
// // @Param input body LoginRequest true "Login data"
// // @Success 200 {object} TokenResponse
// // @Router /auth/login [post]
// func (ac *AuthController) Login(c echo.Context) error {
//     // Implement user authentication logic here
//     // Check user credentials, generate a JWT token, and return it

//     // For demonstration, we'll use a dummy user ID (123) and generate a token
//     userID := 123

//     // Create a JWT token with the user's ID
//     token := echojwt.
//     claims := token.Claims.(jwt.MapClaims)
//     claims["user_id"] = userID
//     claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expiration time

//     // Sign the token with the secret key
//     tokenString, err := token.SignedString(ac.SecretKey)

//     if err != nil {
//         return c.JSON(http.StatusInternalServerError, err)
//     }

//     return c.JSON(http.StatusOK, TokenResponse{Token: tokenString})
// }
