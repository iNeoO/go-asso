package userapi

import (
	"github.com/gofiber/fiber/v2"

	userdomain "github.com/ineoo/go-planigramme/pkg/user"
)

func registerRoutes(app fiber.Router, service *userdomain.Service) {
	userGroup := app.Group("/users")
	userGroup.Get("/", GetUsers(service))
	userGroup.Get("/:id", GetUser(service))
	userGroup.Post("/", CreateUser(service))
	// userGroup.Patch("/:id", UpdateUser(service))
	// userGroup.Delete("/:id", RemoveUser(service))
}
