package jwt_middleware

import (
	"os"
	"strings"
	response_dto "tasteplorer-internal-api/app/dto/response"
	employee_service "tasteplorer-internal-api/app/service/employee"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func JWTMiddleware(c *fiber.Ctx) error {
	var responseDto response_dto.ResponseDto

	// Get the token from the Authorization header
	authHeader := c.Get("authorization")
	if authHeader == "" {
		responseDto = response_dto.ResponseDto{
			Message: "Missing authorization token",
			Code:    fiber.StatusUnauthorized,
			Error:   "Missing authorization token",
			Data:    nil,
		}
		return c.Status(fiber.StatusUnauthorized).JSON(responseDto)
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
		responseDto = response_dto.ResponseDto{
			Message: "Invalid or expired token",
			Code:    fiber.StatusUnauthorized,
			Error:   err,
			Data:    nil,
		}
		return c.Status(fiber.StatusUnauthorized).JSON(responseDto)
		// return fiber.NewError(fiber.StatusUnauthorized, "Invalid or expired token")
	}

	// Token is valid, so store the claims in the context for further use
	c.Locals("user", claims)

	return c.Next() // Continue to the next handler
}
