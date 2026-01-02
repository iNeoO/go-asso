package activityapi

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"

	activitydomain "github.com/ineoo/go-planigramme/pkg/activity"
	membershipdomain "github.com/ineoo/go-planigramme/pkg/membership"
	sessiondomain "github.com/ineoo/go-planigramme/pkg/session"
)

func RegisterRoutes(app fiber.Router, db *sqlx.DB) {
	activityRepo := activitydomain.NewRepository(db)
	membershipRepo := membershipdomain.NewRepository(db)
	sessionRepo := sessiondomain.NewRepository(db)
	membershipService := membershipdomain.NewService(*membershipRepo)
	sessionService := sessiondomain.NewService(*sessionRepo)
	activityService := activitydomain.NewService(*activityRepo, *membershipService)

	registerRoutes(app, activityService, membershipService, sessionService)
}
