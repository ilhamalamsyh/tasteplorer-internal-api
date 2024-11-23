package jwt_middleware

import (
	"os"
	"strings"
	dto_custom_error "tasteplorer-internal-api/app/dto/errors"
	employee_service "tasteplorer-internal-api/app/service/employee"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func JWTMiddleware(c *fiber.Ctx) error {
	var customeError dto_custom_error.CustomError

	// Get the token from the Authorization header
	authHeader := c.Get("authorization")
	if authHeader == "" {
		customeError = dto_custom_error.CustomError{
			Message: "Missing authorization token",
			Code:    fiber.StatusBadRequest,
		}
		return c.Status(fiber.StatusBadRequest).JSON(customeError)
	}

	// Split "Bearer <token>"
	tokenString := strings.Split(authHeader, " ")[1]

	// Parse the token
	claims := &employee_service.Claims{}
	secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))

	// Parse the JWT token and validate it
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil || !token.Valid {
		customeError = dto_custom_error.CustomError{
			Message: "Invalid or expired token",
			Code:    fiber.StatusBadRequest,
		}
		return c.Status(fiber.StatusBadRequest).JSON(customeError)
		// return fiber.NewError(fiber.StatusUnauthorized, "Invalid or expired token")
	}

	// Token is valid, so store the claims in the context for further use
	c.Locals("user", claims)

	return c.Next() // Continue to the next handler
}
