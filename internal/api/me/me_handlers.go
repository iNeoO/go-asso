package meapi

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/ineoo/go-planigramme/pkg/activity"
	"github.com/ineoo/go-planigramme/pkg/organization"
	"github.com/ineoo/go-planigramme/pkg/user"
)

// GetProfile godoc
// @Summary Get user authenticated
// @Description Returns the user authenticated.
// @Tags me
// @Security BearerAuth
// @Produce json
// @Success 200 {object} MeProfileEnvelope
// @Failure 401 {object} MeErrorEnvelope
// @Failure 500 {object} MeErrorEnvelope
// @Router /me/profile [get]
func GetProfile(userService *user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userId, ok := c.Locals("user_id").(uuid.UUID)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(MeErrorResponse("unauthorized"))
		}

		user, err := userService.GetById(userId)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(MeErrorResponse("failed to retrieve user"))
		}

		return c.JSON(MeProfileSuccessResponse(user))
	}
}

// GetOrganizations godoc
// @Summary List organizations for current user
// @Description Returns organizations linked to the authenticated user.
// @Tags me
// @Security BearerAuth
// @Produce json
// @Success 200 {object} MeOrganizationsEnvelope
// @Failure 401 {object} MeErrorEnvelope
// @Failure 500 {object} MeErrorEnvelope
// @Router /me/organizations [get]
func GetOrganizations(organization *organization.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userId, ok := c.Locals("user_id").(uuid.UUID)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(MeErrorResponse("unauthorized"))
		}

		organizations, err := organization.ListByUserId(userId)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(MeErrorResponse("failed to retrieve organizations"))
		}

		return c.JSON(MeOrganizationSuccessResponse(organizations))
	}
}

func GetActivities(activityService *activity.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userId, ok := c.Locals("user_id").(uuid.UUID)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(MeErrorResponse("unauthorized"))
		}

		activities, err := activityService.ListActivitiesForUser(userId)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(MeErrorResponse("failed to retrieve activities"))
		}

		return c.JSON(MeActivitiesSuccessResponse(activities))
	}
}
