package organizationapi

import (
	"github.com/gofiber/fiber/v2"
	authapi "github.com/ineoo/go-planigramme/internal/api/auth"
	organizationdomain "github.com/ineoo/go-planigramme/pkg/organization"
	sessiondomain "github.com/ineoo/go-planigramme/pkg/session"
)

func registerRoutes(app fiber.Router, organizationService *organizationdomain.Service, sessionService *sessiondomain.Service) {
	organizationGroup := app.Group("/organizations")
	organizationGroup.Get("/", authapi.Protected(sessionService), GetOrganizations(organizationService))
	organizationGroup.Get("/:id", authapi.Protected(sessionService), GetOrganization(organizationService))
	organizationGroup.Get("/:id/users", authapi.Protected(sessionService), ListOrganizationUsers(organizationService))
	organizationGroup.Post("/:id/join", authapi.Protected(sessionService), JoinOrganization(organizationService))
	organizationGroup.Delete("/:id/leave", authapi.Protected(sessionService), LeaveOrganization(organizationService))
	organizationGroup.Patch("/:id/users/role", authapi.Protected(sessionService), UpdateOrganizationUserRole(organizationService))
	organizationGroup.Post("/", authapi.Protected(sessionService), CreateOrganization(organizationService))
}
