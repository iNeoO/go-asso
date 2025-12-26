package authapi

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/ineoo/go-planigramme/config"
	sessiondomain "github.com/ineoo/go-planigramme/pkg/session"
	"github.com/ineoo/go-planigramme/pkg/utils"
)

func Protected(service *sessiondomain.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authorization := c.Get("Authorization")
		if strings.HasPrefix(authorization, "Bearer ") {
			authToken := strings.TrimPrefix(authorization, "Bearer ")
			if claims, err := utils.VerifToken(config.Config("AUTH_SECRET"), authToken); err == nil {
				c.Locals("user_id", claims.UserID)
				c.Locals("email", claims.Email)
				return c.Next()
			}
		}

		refreshToken := c.Cookies("refresh_token")
		if refreshToken == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": true,
				"msg":   "unauthorized",
			})
		}

		session, err := service.GetByToken(refreshToken)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": true,
				"msg":   "unauthorized",
			})
		}

		if time.Now().After(session.ExpiresAt) || session.RevokedAt != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": true,
				"msg":   "session expired",
			})
		}

		claims, err := utils.VerifToken(config.Config("REFRESH_SECRET"), refreshToken)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": true,
				"msg":   "unauthorized",
			})
		}

		if claims.UserID != session.UserID {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": true,
				"msg":   "unauthorized",
			})
		}

		newAuthToken, err := utils.GenerateAuthToken(claims.Email, session.UserID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"msg":   "failed to generate auth token",
			})
		}

		c.Set("new_authToken", newAuthToken.Token)
		c.Set("new_authToken_expires_at", newAuthToken.ExpiresAt.Format(time.RFC3339))
		c.Locals("user_id", claims.UserID)
		c.Locals("email", claims.Email)
		return c.Next()
	}
}
