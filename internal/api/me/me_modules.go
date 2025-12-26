package meapi

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"

	organizationdomain "github.com/ineoo/go-planigramme/pkg/organization"
	sessiondomain "github.com/ineoo/go-planigramme/pkg/session"
	userdomain "github.com/ineoo/go-planigramme/pkg/user"
)

func RegisterRoutes(app fiber.Router, db *sqlx.DB) {
	userService := userdomain.NewService(*userdomain.NewRepository(db))
	organizationService := organizationdomain.NewService(*organizationdomain.NewRepository(db))
	sessionService := sessiondomain.NewService(*sessiondomain.NewRepository(db))

	registerRoutes(app, userService, organizationService, sessionService)
}
