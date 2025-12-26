package userapi

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	userdomain "github.com/ineoo/go-planigramme/pkg/user"
	"github.com/ineoo/go-planigramme/pkg/utils"
)

// GetUsers godoc
// @Summary List users
// @Description Returns all users.
// @Tags users
// @Produce json
// @Success 200 {object} UsersEnvelope
// @Failure 500 {object} UserErrorEnvelope
// @Router /users [get]
func GetUsers(service *userdomain.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		users, err := service.List()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(UserErrorResponse(errors.New("failed to retrieve users")))
		}

		return c.JSON(UsersSuccessResponse(users))
	}
}

// GetUser godoc
// @Summary Get user by ID
// @Description Returns a single user.
// @Tags users
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} UserEnvelope
// @Failure 400 {object} UserErrorEnvelope
// @Failure 404 {object} UserErrorEnvelope
// @Router /users/{id} [get]
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

// CreateUser godoc
// @Summary Create user
// @Description Creates a new user.
// @Tags users
// @Accept json
// @Produce json
// @Param payload body CreateUserRequest true "User payload"
// @Success 201 {object} UserEnvelope
// @Failure 400 {object} UserErrorEnvelope
// @Failure 500 {object} UserErrorEnvelope
// @Router /users [post]
func CreateUser(service *userdomain.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		payload := new(CreateUserRequest)
		if err := c.BodyParser(payload); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(UserErrorResponse(errors.New("invalid request payload")))
		}
		if err := utils.NewValidator().Struct(payload); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(UserErrorResponse(err))
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
