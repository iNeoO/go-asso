package userapi

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"

	userdomain "github.com/ineoo/go-planigramme/pkg/user"
)

// RegisterRoutes wires the user module for the given router.
func RegisterRoutes(app fiber.Router, db *sqlx.DB) {
	repo := userdomain.NewRepository(db)
	service := userdomain.NewService(*repo)

	registerRoutes(app, service)
}
