package organizationapi

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/ineoo/go-planigramme/pkg/activity"
	"github.com/ineoo/go-planigramme/pkg/membership"
	"github.com/ineoo/go-planigramme/pkg/organization"
)

// CreateActivity godoc
// @Summary Create activity
// @Description Creates an activity within an organization. Requires bearer auth or refresh cookie.
// @Tags activities
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "Organization ID"
// @Param payload body CreateActivityRequest true "Activity payload"
// @Success 201 {object} ActivityEnvelope
// @Failure 400 {object} ErrorEnvelope
// @Failure 401 {object} ErrorEnvelope
// @Failure 403 {object} ErrorEnvelope
// @Failure 500 {object} ErrorEnvelope
// @Router /organizations/{id}/activities [post]
func CreateActivity(organizationService organization.Service, activityService activity.Service, memberShipService membership.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(ActivityErrorResponse("invalid organization id"))
		}

		payload := new(CreateActivityRequest)
		if err := c.BodyParser(payload); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(ActivityErrorResponse("invalid request payload"))
		}

		userId, ok := c.Locals("user_id").(uuid.UUID)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(ActivityErrorResponse("unauthorized"))
		}

		organizationMember, err := memberShipService.GetOrganizationMember(userId, id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(ActivityErrorResponse("failed to verify membership"))
		}

		if !organization.HasWriteAccess(organizationMember.RoleID) {
			return c.Status(fiber.StatusForbidden).JSON(ActivityErrorResponse("insufficient permissions to create activity"))
		}

		newActivity, err := activityService.CreateActivity(&activity.ActivityCreation{
			Name:            payload.Name,
			CreatorID:       userId,
			Description:     payload.Description,
			StartAt:         payload.StartAt,
			EndAt:           payload.EndAt,
			DurationMinutes: payload.DurationMinutes,
			Capacity:        payload.Capacity,
			OrganizationID:  id,
		})
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(ActivityErrorResponse("failed to create activity"))
		}

		return c.Status(fiber.StatusCreated).JSON(ActivitySuccessResponse(newActivity))
	}
}

// GetActivities godoc
// @Summary List activities
// @Description Lists activities for an organization. Requires bearer auth or refresh cookie.
// @Tags activities
// @Security BearerAuth
// @Produce json
// @Param id path string true "Organization ID"
// @Success 200 {object} ActivitiesEnvelope
// @Failure 400 {object} ErrorEnvelope
// @Failure 401 {object} ErrorEnvelope
// @Failure 403 {object} ErrorEnvelope
// @Failure 500 {object} ErrorEnvelope
// @Router /organizations/{id}/activities [get]
func GetActivities(organizationService organization.Service, activityService activity.Service, memberShipService membership.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(ActivityErrorResponse("invalid organization id"))
		}

		userId, ok := c.Locals("user_id").(uuid.UUID)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(ActivityErrorResponse("unauthorized"))
		}

		organizationMember, err := memberShipService.GetOrganizationMember(userId, id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(ActivityErrorResponse("failed to verify membership"))
		}

		if organizationMember == nil || !organization.HasReadAccess(organizationMember.RoleID) {
			return c.Status(fiber.StatusForbidden).JSON(ActivityErrorResponse("insufficient permissions to list activities"))
		}

		activities, err := activityService.ListActivitiesByOrganization(id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(ActivityErrorResponse("failed to list activities"))
		}

		return c.Status(fiber.StatusOK).JSON(ActivitiesListSuccessResponse(activities))
	}
}
