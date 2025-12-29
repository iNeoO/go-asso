package organizationapi

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"

	membershipdomain "github.com/ineoo/go-planigramme/pkg/membership"
	organizationdomain "github.com/ineoo/go-planigramme/pkg/organization"
	sessiondomain "github.com/ineoo/go-planigramme/pkg/session"
)

func RegisterRoutes(app fiber.Router, db *sqlx.DB) {
	organizationRepo := organizationdomain.NewRepository(db)
	membershipRepo := membershipdomain.NewRepository(db)
	sessionRepo := sessiondomain.NewRepository(db)
	organizationService := organizationdomain.NewService(*organizationRepo)
	membershipService := membershipdomain.NewService(*membershipRepo)
	sessionService := sessiondomain.NewService(*sessionRepo)

	registerRoutes(app, organizationService, membershipService, sessionService)
}
