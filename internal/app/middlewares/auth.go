package middleware

import (
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

// JWTConfig defines the configuration for JWT middleware.
var JWTConfig = echojwt.Config{
	SigningKey: []byte(os.Getenv("JWT_Secret_Key")), // Replace with your secret key
}

// AuthMiddleware creates and returns JWT middleware.
func AuthMiddleware() echo.MiddlewareFunc {
	return echojwt.WithConfig(JWTConfig)
}
