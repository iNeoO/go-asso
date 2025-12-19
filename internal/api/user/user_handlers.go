package userapi

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	userdomain "github.com/ineoo/go-planigramme/pkg/user"
	"github.com/ineoo/go-planigramme/pkg/utils"
)

func GetUsers(service *userdomain.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		users, err := service.List()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		}

		return c.JSON(UsersSuccessResponse(users))
	}
}

func GetUser(service *userdomain.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(UserErrorResponse(errors.New("invalid user id")))
		}

		user, err := service.GetById(id)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(UserErrorResponse(errors.New("user was not found")))
		}

		return c.JSON(UserSuccessResponse(user))
	}
}

func CreateUser(service *userdomain.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		payload := new(CreateUserRequest)
		if err := c.BodyParser(payload); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(UserErrorResponse(errors.New("invalid request payload")))
		}

		passwordHash, err := utils.HashPassword(payload.Password)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(UserErrorResponse(errors.New("failed to hash password")))
		}

		newUser := &userdomain.User{
			FirstName:    payload.FirstName,
			LastName:     payload.LastName,
			Email:        payload.Email,
			PasswordHash: passwordHash,
		}

		u, err := service.Create(newUser)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(UserErrorResponse(err))
		}

		return c.Status(fiber.StatusCreated).JSON(UserSuccessResponse(u))
	}
}
