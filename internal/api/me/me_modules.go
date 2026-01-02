package meapi

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"

	activitydomain "github.com/ineoo/go-planigramme/pkg/activity"
	membershipdomain "github.com/ineoo/go-planigramme/pkg/membership"
	organizationdomain "github.com/ineoo/go-planigramme/pkg/organization"
	sessiondomain "github.com/ineoo/go-planigramme/pkg/session"
	userdomain "github.com/ineoo/go-planigramme/pkg/user"
)

func RegisterRoutes(app fiber.Router, db *sqlx.DB) {
	userService := userdomain.NewService(*userdomain.NewRepository(db))
	organizationService := organizationdomain.NewService(*organizationdomain.NewRepository(db))
	sessionService := sessiondomain.NewService(*sessiondomain.NewRepository(db))
	membershipService := membershipdomain.NewService(*membershipdomain.NewRepository(db))
	activityService := activitydomain.NewService(*activitydomain.NewRepository(db), *membershipService)

	registerRoutes(app, userService, organizationService, sessionService, activityService)
}
