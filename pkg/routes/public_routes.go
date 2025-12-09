package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ineoo/go-planigramme/app/controllers"
)

func PublicRoutes(a *fiber.App) {
	route := a.Group("/api/v1")

	route.Get("/users", controllers.GetUsers)
	route.Get("/user/:id", controllers.GetUser)
	route.Post("/user", controllers.CreateUser)
	route.Put("/user/:id", controllers.UpdateUser)
	route.Delete("/user/:id", controllers.DeleteUser)
}