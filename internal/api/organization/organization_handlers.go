package organizationapi

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/ineoo/go-planigramme/pkg/membership"
	"github.com/ineoo/go-planigramme/pkg/organization"
	"github.com/ineoo/go-planigramme/pkg/utils"
)

// GetOrganizations godoc
// @Summary List organizations
// @Description Returns all organizations. Requires bearer auth or refresh cookie.
// @Tags organizations
// @Security BearerAuth
// @Produce json
// @Success 200 {object} OrganizationsEnvelope
// @Failure 401 {object} ErrorEnvelope
// @Failure 500 {object} ErrorEnvelope
// @Router /organizations [get]
func GetOrganizations(service *organization.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		organizations, err := service.List()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(OrganizationErrorResponse("failed to retrieve organizations"))
		}

		return c.JSON(OrganizationsSuccessResponse(organizations))
	}
}

// GetOrganization godoc
// @Summary Get organization by ID
// @Description Returns a single organization. Requires bearer auth or refresh cookie.
// @Tags organizations
// @Security BearerAuth
// @Produce json
// @Param id path string true "Organization ID"
// @Success 200 {object} OrganizationEnvelope
// @Failure 400 {object} ErrorEnvelope
// @Failure 401 {object} ErrorEnvelope
// @Failure 404 {object} ErrorEnvelope
// @Router /organizations/{id} [get]
func GetOrganization(service *organization.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(OrganizationErrorResponse("invalid organization id"))
		}

		organization, err := service.GetByID(id)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(OrganizationErrorResponse("organization was not found"))
		}

		return c.JSON(OrganizationSuccessResponse(organization))
	}
}

// CreateOrganization godoc
// @Summary Create organization
// @Description Creates an organization. Requires bearer auth or refresh cookie.
// @Tags organizations
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param payload body CreateOrganizationRequest true "Organization payload"
// @Success 201 {object} OrganizationEnvelope
// @Failure 400 {object} ErrorEnvelope
// @Failure 401 {object} ErrorEnvelope
// @Failure 500 {object} ErrorEnvelope
// @Router /organizations [post]
func CreateOrganization(organizationService *organization.Service, membershipService *membership.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		payload := new(CreateOrganizationRequest)
		if err := c.BodyParser(payload); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(OrganizationErrorResponse("invalid request payload"))
		}
		if err := utils.NewValidator().Struct(payload); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(OrganizationErrorResponse("invalid request payload"))
		}
		userId, ok := c.Locals("user_id").(uuid.UUID)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(OrganizationErrorResponse("unauthorized"))
		}

		newOrganization := &organization.Organization{
			Name: payload.Name,
		}

		createdOrganization, err := organizationService.Create(newOrganization)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(OrganizationErrorResponse("failed to create organization"))
		}

		if _, err := membershipService.Join(createdOrganization.ID, userId, organization.RoleCreator); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(OrganizationErrorResponse("failed to join organization"))
		}

		return c.Status(fiber.StatusCreated).JSON(OrganizationSuccessResponse(createdOrganization))
	}
}

// JoinOrganization godoc
// @Summary Join organization
// @Description Joins an organization. Requires bearer auth or refresh cookie.
// @Tags organizations
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "Organization ID"
// @Param payload body JoinOrganizationRequest true "Join organization payload"
// @Success 200 {object} OrganizationUserEnvelope
// @Failure 400 {object} ErrorEnvelope
// @Failure 401 {object} ErrorEnvelope
// @Failure 500 {object} ErrorEnvelope
// @Router /organizations/{id}/join [post]
func JoinOrganization(service *membership.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(OrganizationErrorResponse("invalid organization id"))
		}
		userId, ok := c.Locals("user_id").(uuid.UUID)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(OrganizationErrorResponse("unauthorized"))
		}

		organizationUser, err := service.Join(id, userId, organization.RoleNotValidated)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(OrganizationErrorResponse("failed to join organization"))
		}

		return c.JSON(OrganizationUserSuccessResponse(organizationUser))
	}
}

// LeaveOrganization godoc
// @Summary Leave organization
// @Description Leaves an organization. Requires bearer auth or refresh cookie.
// @Tags organizations
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "Organization ID"
// @Success 200 {object} fiber.Map
// @Failure 400 {object} ErrorEnvelope
// @Failure 401 {object} ErrorEnvelope
// @Failure 404 {object} ErrorEnvelope
// @Failure 500 {object} ErrorEnvelope
// @Router /organizations/{id}/leave [post]
func LeaveOrganization(service *membership.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(OrganizationErrorResponse("invalid organization id"))
		}
		userId, ok := c.Locals("user_id").(uuid.UUID)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(OrganizationErrorResponse("unauthorized"))
		}

		organizationUser, err := service.GetOrganizationUser(id, userId)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(OrganizationErrorResponse("failed to get organization user"))
		}

		if organization.HasAdminAccess(organizationUser.RoleID) {
			return c.Status(fiber.StatusForbidden).JSON(OrganizationErrorResponse("creator cannot leave the organization"))
		}

		if err := service.Leave(id, userId); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(OrganizationErrorResponse("failed to leave organization"))
		}

		return c.JSON(fiber.Map{
			"status":  true,
			"message": "successfully left organization",
		})
	}
}

// UpdateOrganizationUserRole godoc
// @Summary Update organization user role
// @Description Updates an organization user's role. Requires bearer auth or refresh cookie.
// @Tags organizations
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "Organization ID"
// @Param payload body UpdateOrganizationUserRoleRequest true "Update organization user role payload"
// @Success 200 {object} OrganizationUserEnvelope
// @Failure 400 {object} ErrorEnvelope
// @Failure 401 {object} ErrorEnvelope
// @Failure 404 {object} ErrorEnvelope
// @Failure 500 {object} ErrorEnvelope
func UpdateOrganizationUserRole(service *membership.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(OrganizationErrorResponse("invalid organization id"))
		}
		payload := new(UpdateOrganizationUserRoleRequest)
		if err := c.BodyParser(payload); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(OrganizationErrorResponse("invalid request payload"))
		}
		if err := utils.NewValidator().Struct(payload); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(OrganizationErrorResponse("invalid request payload"))
		}
		userId, ok := c.Locals("user_id").(uuid.UUID)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(OrganizationErrorResponse("unauthorized"))
		}

		organizationUser, err := service.GetOrganizationUser(id, userId)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(OrganizationErrorResponse("organization user not found"))
		}

		if !organization.HasAdminAccess(organizationUser.RoleID) {
			return c.Status(fiber.StatusForbidden).JSON(OrganizationErrorResponse("insufficient permissions to update user role"))
		}

		targetUser, err := service.GetOrganizationUser(id, payload.UserID)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(OrganizationErrorResponse("organization user not found"))
		}

		if organization.HasAdminAccess(targetUser.RoleID) {
			return c.Status(fiber.StatusForbidden).JSON(OrganizationErrorResponse("cannot change creator/admin role"))
		}

		updatedOrganizationUser, err := service.AssignRole(id, payload.UserID, payload.RoleID)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(OrganizationErrorResponse("failed to update organization user role"))
		}

		return c.JSON(OrganizationUserSuccessResponse(updatedOrganizationUser))
	}
}

// ListOrganizationUsers godoc
// @Summary List organization users
// @Description Lists users of an organization with their roles. Requires bearer auth or refresh cookie.
// @Tags organizations
// @Security BearerAuth
// @Produce json
// @Param id path string true "Organization ID"
// @Success 200 {object} OrganizationMembersEnvelope
// @Failure 400 {object} ErrorEnvelope
// @Failure 401 {object} ErrorEnvelope
// @Failure 403 {object} ErrorEnvelope
// @Failure 404 {object} ErrorEnvelope
// @Failure 500 {object} ErrorEnvelope
func ListOrganizationUsers(service *membership.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(OrganizationErrorResponse("invalid organization id"))
		}
		userId, ok := c.Locals("user_id").(uuid.UUID)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(OrganizationErrorResponse("unauthorized"))
		}

		currentUser, err := service.GetOrganizationUser(id, userId)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(OrganizationErrorResponse("organization user not found"))
		}

		if !organization.HasAdminAccess(currentUser.RoleID) {
			return c.Status(fiber.StatusForbidden).JSON(OrganizationErrorResponse("insufficient permissions to list organization users"))
		}

		users, err := service.ListOrganizationMembers(id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(OrganizationErrorResponse("failed to list organization users"))
		}

		return c.JSON(OrganizationMembersSuccessResponse(users))
	}
}
