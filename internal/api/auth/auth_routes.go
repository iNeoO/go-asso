package authapi

import (
	"github.com/gofiber/fiber/v2"

	sessiondomain "github.com/ineoo/go-planigramme/pkg/session"
	userdomain "github.com/ineoo/go-planigramme/pkg/user"
)

func registerRoutes(app fiber.Router, userService *userdomain.Service, sessionService *sessiondomain.Service) {
	userGroup := app.Group("/auth")
	userGroup.Post("/login", Login(userService, sessionService))
	userGroup.Post("/logout", Logout(sessionService))
}
