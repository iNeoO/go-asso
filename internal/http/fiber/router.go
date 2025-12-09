package fiberhandler

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App, userHandler *UserHandler) {
	sw := app.Group("/swagger")
	sw.Get("*", swagger.HandlerDefault)

	api := app.Group("/api/v1")
	userHandler.Register(api)

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "sorry, endpoint is not found",
		})
	})
}
