package authapi

import (
	"github.com/gofiber/fiber/v2"

	authdomain "github.com/ineoo/go-planigramme/pkg/auth"
	sessiondomain "github.com/ineoo/go-planigramme/pkg/session"
	userdomain "github.com/ineoo/go-planigramme/pkg/user"
)

func registerRoutes(app fiber.Router, authService *authdomain.Service,userService *userdomain.Service, sessionService *sessiondomain.Service) {
	userGroup := app.Group("/auth")
	userGroup.Post("/login", Login(authService,userService, sessionService))
	userGroup.Post("/logout", Logout(authService, userService, sessionService))
}
