package activityapi

import (
	"github.com/gofiber/fiber/v2"
	authapi "github.com/ineoo/go-planigramme/internal/api/auth"
	"github.com/ineoo/go-planigramme/pkg/activity"
	"github.com/ineoo/go-planigramme/pkg/membership"
	"github.com/ineoo/go-planigramme/pkg/session"
)

func registerRoutes(a fiber.Router, activityService *activity.Service, membershipService *membership.Service, sessionService *session.Service) {
	activityGroup := a.Group("/activities")
	activityGroup.Get("/:id", authapi.Protected(sessionService), GetActivity(activityService, membershipService))
	activityGroup.Post("/:id/join", authapi.Protected(sessionService), joinActivity(activityService))
}
