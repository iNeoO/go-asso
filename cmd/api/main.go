package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	swaggerdocs "github.com/ineoo/go-planigramme/swagger"
	_ "github.com/joho/godotenv/autoload"

	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/ineoo/go-planigramme/config"
	authapi "github.com/ineoo/go-planigramme/internal/api/auth"
	meapi "github.com/ineoo/go-planigramme/internal/api/me"
	organizationapi "github.com/ineoo/go-planigramme/internal/api/organization"
	swaggerapi "github.com/ineoo/go-planigramme/internal/api/swagger"
	userapi "github.com/ineoo/go-planigramme/internal/api/user"
	"github.com/ineoo/go-planigramme/internal/database"
	"github.com/ineoo/go-planigramme/pkg/utils"
)

// @title API
// @version 1.0
// @description This is an auto-generated API Docs.
// @termsOfService http://swagger.io/terms/
// @contact.name go-planigramme
// @contact.email go-planigramme@tuturu.io
// @license.name MIT
// @license.url https://mit-license.org/
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @BasePath /api/v1
func main() {
	swaggerdocs.SwaggerInfo.BasePath = "/api/v1"

	app := fiber.New(config.Fiber())

	db, err := database.Connect()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))
	app.Use(requestid.New())

	apiGroup := app.Group("/api")
	swaggerapi.RegisterRoutes(apiGroup)

	v1Group := apiGroup.Group("/v1")

	authapi.RegisterRoutes(v1Group, db)
	userapi.RegisterRoutes(v1Group, db)
	organizationapi.RegisterRoutes(v1Group, db)
	meapi.RegisterRoutes(v1Group, db)
	utils.StartServerWithGracefulShutdown(app)
}
