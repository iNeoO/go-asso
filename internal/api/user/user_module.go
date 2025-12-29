package userapi

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"

	userdomain "github.com/ineoo/go-planigramme/pkg/user"
)

func RegisterRoutes(app fiber.Router, db *sqlx.DB) {
	repo := userdomain.NewRepository(db)
	service := userdomain.NewService(*repo)

	registerRoutes(app, service)
}
