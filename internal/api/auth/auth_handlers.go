package authapi

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	httpUtils "github.com/ineoo/go-planigramme/internal/utils"
	sessiondomain "github.com/ineoo/go-planigramme/pkg/session"
	userdomain "github.com/ineoo/go-planigramme/pkg/user"
	cryptoutils "github.com/ineoo/go-planigramme/pkg/utils"
)

// Login godoc
// @Summary Login
// @Description Authenticates a user, sets refresh cookie, and returns an access token.
// @Tags auth
// @Accept json
// @Produce json
// @Param payload body LoginRequest true "Credentials"
// @Success 200 {object} AuthEnvelope
// @Failure 400 {object} AuthErrorEnvelope
// @Failure 401 {object} AuthErrorEnvelope
// @Failure 500 {object} AuthErrorEnvelope
// @Router /auth/login [post]
func Login(userService *userdomain.Service, sessionService *sessiondomain.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		payload := new(LoginRequest)
		if err := c.BodyParser(payload); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(AuthErrorResponse(errors.New("invalid request payload")))
		}
		if err := cryptoutils.NewValidator().Struct(payload); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(AuthErrorResponse(errors.New("invalid request payload")))
		}
		u, err := userService.GetByEmail(payload.Email)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(AuthErrorResponse(errors.New("invalid email or password")))
		}
		if !cryptoutils.CheckPasswordHash(payload.Password, u.PasswordHash) {
			return c.Status(fiber.StatusUnauthorized).JSON(AuthErrorResponse(errors.New("invalid email or password")))
		}

		authToken, err := cryptoutils.GenerateAuthToken(u.Email, u.ID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(AuthErrorResponse(errors.New("failed to generate auth token")))
		}
		refreshToken, err := cryptoutils.GenerateRefreshToken(u.Email, u.ID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(AuthErrorResponse(errors.New("failed to generate refresh token")))
		}

		session := &sessiondomain.Session{
			UserID:    u.ID,
			ExpiresAt: refreshToken.ExpiresAt,
			TokenHash: refreshToken.Token,
		}

		_, err = sessionService.Create(session)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(AuthErrorResponse(errors.New("failed to create session")))
		}

		cookie := httpUtils.CreateRefreshCookie(refreshToken.Token, refreshToken.ExpiresAt)

		c.Cookie(cookie)

		return c.JSON(AuthSuccessResponse(&AuthData{
			Token:     authToken.Token,
			ExpiresAt: authToken.ExpiresAt.Unix(),
		}))
	}
}

func Logout(sessionService *sessiondomain.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		_ = sessionService
		// Implementation of logout handler
		return nil
	}
}
