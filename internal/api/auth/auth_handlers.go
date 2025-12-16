package authapi

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/ineoo/go-planigramme/internal/utils"
	authdomain "github.com/ineoo/go-planigramme/pkg/auth"
	sessiondomain "github.com/ineoo/go-planigramme/pkg/session"
	userdomain "github.com/ineoo/go-planigramme/pkg/user"
)

func Login(authService *authdomain.Service, userService *userdomain.Service, sessionService *sessiondomain.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		payload := new(LoginRequest)
		if err := c.BodyParser(payload); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(AuthErrorResponse(errors.New("invalid request payload")))
		}
		u, err := userService.GetByEmail(payload.Email)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(AuthErrorResponse(errors.New("invalid email or password")))
		}
		if !authService.CheckPasswordHash(payload.Password, u.PasswordHash) {
			return c.Status(fiber.StatusUnauthorized).JSON(AuthErrorResponse(errors.New("invalid email or password")))
		}

		authToken, err := authService.GenerateAuthToken(u.Email, u.ID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(AuthErrorResponse(errors.New("failed to generate auth token")))
		}
		refreshToken, err := authService.GenerateRefreshToken(u.Email, u.ID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(AuthErrorResponse(errors.New("failed to generate refresh token")))
		}

		sesion := &sessiondomain.Session{
			UserID:       u.ID,
			ExpiresAt:    refreshToken.ExpiresAt,
			Token: refreshToken.Token,
		}

		_, err = sessionService.Create(sesion)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(AuthErrorResponse(errors.New("failed to create session")))
		}

		cookie :=utils.CreateRefreshCookie(refreshToken.Token, refreshToken.ExpiresAt)

		c.Cookie(cookie)
		
		return c.JSON(AuthSuccessResponse(&AuthData{
			Token:     authToken.Token,
			ExpiresAt: authToken.ExpiresAt.UTC().Unix(),
		}))
	}
}

func Logout(authService *authdomain.Service, userService *userdomain.Service, sessionService *sessiondomain.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Implementation of logout handler
		return nil
	}
}