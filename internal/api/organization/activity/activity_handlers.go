package organizationapi

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/ineoo/go-planigramme/pkg/activity"
	"github.com/ineoo/go-planigramme/pkg/organization"
)

func CreateActivity(organizationService organization.Service, activityService activity.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse("invalid organization id"))
		}

		payload := new(CreateActivityRequest)
		if err := c.BodyParser(payload); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(OrganizationErrorResponse("invalid request payload"))
		}

		org, err := organizationService.GetByID(id)
		if err != nil || org == nil {
			return c.Status(fiber.StatusNotFound).JSON(OrganizationErrorResponse("organization not found"))
		}

	}
}
