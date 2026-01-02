package activityapi

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/ineoo/go-planigramme/pkg/activity"
	"github.com/ineoo/go-planigramme/pkg/membership"
	oranizationdomain "github.com/ineoo/go-planigramme/pkg/organization"
)

func GetActivity(activityService *activity.Service, memberShipService *membership.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(ActivityErrorResponse("invalid activity id"))
		}

		userId, ok := c.Locals("user_id").(uuid.UUID)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(ActivityErrorResponse("unauthorized"))
		}

		organizations, err := memberShipService.ListOrganizationsForUser(userId)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(ActivityErrorResponse("failed to retrieve organizations"))
		}

		foundActivity, err := activityService.GetActivityByID(id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(ActivityErrorResponse("failed to retrieve activity"))
		}

		hasAccess := false
		for _, org := range organizations {
			if org.OrganizationID == foundActivity.OrganizationID && oranizationdomain.HasReadAccess(org.RoleID) {
				hasAccess = true
				break
			}
		}
		if !hasAccess {
			return c.Status(fiber.StatusForbidden).JSON(ActivityErrorResponse("forbidden"))
		}

		return c.JSON(foundActivity)
	}
}

func joinActivity(activity *activity.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNotImplemented)
	}
}
