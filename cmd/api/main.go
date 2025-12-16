package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"

	"github.com/ineoo/go-planigramme/config"
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
// @contact.email go-planigramme@mail.com
// @license.name MIT
// @license.url https://mit-license.org/
// @in header
// @name Authorization
// @BasePath /api
func main() {
	app := fiber.New(config.Fiber())

	db, err := database.Connect()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	swaggerapi.RegisterRoutes(app)

	apiGroup := app.Group("/api/v1")
	userapi.RegisterRoutes(apiGroup, db)

	utils.StartServerWithGracefulShutdown(app)
}
