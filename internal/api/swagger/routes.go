package swaggerapi

import (
	"github.com/gofiber/fiber/v2"

	swagger "github.com/arsmn/fiber-swagger/v2"
)

func registerRoutes(a fiber.Router) {
	route := a.Group("/swagger")
	route.Get("*", swagger.HandlerDefault)
}
