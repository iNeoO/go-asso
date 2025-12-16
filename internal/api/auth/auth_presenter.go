package authapi

import "github.com/gofiber/fiber/v2"

func AuthErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   nil,
		"error":  err.Error(),
	}
}

type AuthData struct {
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expires_at"`
}

func AuthSuccessResponse(data *AuthData) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}