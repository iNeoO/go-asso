package meapi

import (
	"github.com/gofiber/fiber/v2"
	authapi "github.com/ineoo/go-planigramme/internal/api/auth"
	"github.com/ineoo/go-planigramme/pkg/activity"
	"github.com/ineoo/go-planigramme/pkg/organization"
	"github.com/ineoo/go-planigramme/pkg/session"
	"github.com/ineoo/go-planigramme/pkg/user"
)

func registerRoutes(a fiber.Router, userService *user.Service, organizationService *organization.Service, sessionService *session.Service, activityService *activity.Service) {
	meGroup := a.Group("/me")
	meGroup.Get("/profile", authapi.Protected(sessionService), GetProfile(userService))
	meGroup.Get("/organizations", authapi.Protected(sessionService), GetOrganizations(organizationService))
	meGroup.Get("/activities", authapi.Protected(sessionService), GetActivities(activityService))
}
