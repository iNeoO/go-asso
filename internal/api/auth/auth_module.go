package authapi

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"

	sessiondomain "github.com/ineoo/go-planigramme/pkg/session"
	userdomain "github.com/ineoo/go-planigramme/pkg/user"
)

func RegisterRoutes(app fiber.Router, db *sqlx.DB) {
	userrepo := userdomain.NewRepository(db)
	userservice := userdomain.NewService(*userrepo)
	sessionrepo := sessiondomain.NewRepository(db)
	sessionservice := sessiondomain.NewService(*sessionrepo)

	registerRoutes(app, userservice, sessionservice)
}
