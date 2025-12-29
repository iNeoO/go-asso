package utils

import (
	"time"

	"github.com/gofiber/fiber/v2"
)


func CreateRefreshCookie(token string, expiresAt time.Time) *fiber.Cookie {
	cookie := new(fiber.Cookie)
	cookie.Name = "refresh_token"
	cookie.Value = token
	cookie.Path = "/"
	cookie.Expires = expiresAt
	cookie.HTTPOnly = true
	return cookie
}
