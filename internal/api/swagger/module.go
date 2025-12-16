package swaggerapi

import (
	"github.com/gofiber/fiber/v2"
)

// RegisterRoutes wires the swagger endpoints.
func RegisterRoutes(app fiber.Router) {
	registerRoutes(app)
}
