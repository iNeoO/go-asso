package organizationapi

import (
	"github.com/gofiber/fiber/v2"
	authapi "github.com/ineoo/go-planigramme/internal/api/auth"
	membershipdomain "github.com/ineoo/go-planigramme/pkg/membership"
	organizationdomain "github.com/ineoo/go-planigramme/pkg/organization"
	sessiondomain "github.com/ineoo/go-planigramme/pkg/session"
)

func registerRoutes(app fiber.Router, organizationService *organizationdomain.Service, membershipService *membershipdomain.Service, sessionService *sessiondomain.Service) {
	organizationGroup := app.Group("/organizations")
	organizationGroup.Get("/", authapi.Protected(sessionService), GetOrganizations(organizationService))
	organizationGroup.Get("/:id", authapi.Protected(sessionService), GetOrganization(organizationService))
	organizationGroup.Get("/:id/users", authapi.Protected(sessionService), ListOrganizationUsers(membershipService))
	organizationGroup.Post("/:id/join", authapi.Protected(sessionService), JoinOrganization(membershipService))
	organizationGroup.Delete("/:id/leave", authapi.Protected(sessionService), LeaveOrganization(membershipService))
	organizationGroup.Patch("/:id/users/role", authapi.Protected(sessionService), UpdateOrganizationUserRole(membershipService))
	organizationGroup.Post("/", authapi.Protected(sessionService), CreateOrganization(organizationService, membershipService))
}
