package organizationapi

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"

	organizationdomain "github.com/ineoo/go-planigramme/pkg/organization"
	sessiondomain "github.com/ineoo/go-planigramme/pkg/session"
)

func RegisterRoutes(app fiber.Router, db *sqlx.DB) {
	organizationRepo := organizationdomain.NewRepository(db)
	sessionRepo := sessiondomain.NewRepository(db)
	organizationService := organizationdomain.NewService(*organizationRepo)
	sessionService := sessiondomain.NewService(*sessionRepo)

	registerRoutes(app, organizationService, sessionService)
}
