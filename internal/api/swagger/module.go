package swaggerapi

import (
	"github.com/gofiber/fiber/v2"

	_ "github.com/ineoo/go-planigramme/swagger"
)

func RegisterRoutes(app fiber.Router) {
	registerRoutes(app)
}
