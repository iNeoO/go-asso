package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"

	"github.com/ineoo/go-planigramme/config"
	"github.com/ineoo/go-planigramme/internal/database"
	fiberhandler "github.com/ineoo/go-planigramme/internal/http/fiber"
	"github.com/ineoo/go-planigramme/pkg/utils"
	"github.com/ineoo/go-planigramme/user"
	userpostgres "github.com/ineoo/go-planigramme/user/postgres"
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
	app := fiber.New(config.Fiber())

	db, err := database.Connect()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	userRepo := userpostgres.NewRepository(db)
	userService := user.NewService(userRepo)
	userHandler := fiberhandler.NewUserHandler(userService)

	fiberhandler.Register(app, userHandler)

	utils.StartServerWithGracefulShutdown(app)
}
