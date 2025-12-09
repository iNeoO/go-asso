package main

import (
	"github.com/gofiber/fiber/v2"
	// _ "github.com/ineoo/go-planigramme/docs"
	"github.com/ineoo/go-planigramme/pkg/configs"
	"github.com/ineoo/go-planigramme/pkg/routes"
	"github.com/ineoo/go-planigramme/pkg/utils"
	_ "github.com/joho/godotenv/autoload"
)

// @title API
// @version 1.0
// @description This is an auto-generated API Docs.
// @termsOfService http://swagger.io/terms/
// @contact.name go-planigramme
// @contact.email go-planigramme@mail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /api
func main() {
	
	config := configs.FiberConfig()
	app := fiber.New(config)

	routes.SwaggerRoute(app)
	routes.PublicRoutes(app)

	routes.NotFoundRoute(app)

	utils.StartServerWithGracefulShutdown(app)
}