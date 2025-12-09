package fiberhandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/ineoo/go-planigramme/internal/http/dto"
	"github.com/ineoo/go-planigramme/pkg/utils"
	"github.com/ineoo/go-planigramme/user"
)

type UserHandler struct {
	service *user.Service
}

func NewUserHandler(service *user.Service) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) Register(router fiber.Router) {
	router.Get("/users", h.GetUsers)
	router.Get("/user/:id", h.GetUser)
	router.Post("/user", h.CreateUser)
}

func (h *UserHandler) GetUsers(c *fiber.Ctx) error {
	users, err := h.service.ListUsers(c.UserContext())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error": false,
		"msg":   "",
		"count": len(users),
		"users": mapUsersToResponse(users),
	})
}

// GetUser handles GET /user/:id.
func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "invalid user id",
		})
	}

	user, err := h.service.GetUser(c.UserContext(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "user was not found",
			"user":  nil,
		})
	}

	return c.JSON(fiber.Map{
		"error": false,
		"msg":   "user was found",
		"user":  mapUserToResponse(user),
	})
}

// CreateUser handles POST /user.
func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	payload := new(dto.CreateUserRequest)
	if err := c.BodyParser(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "invalid request body",
		})
	}

	if err := utils.NewValidator().Struct(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  true,
			"msg":    "payload validation failed",
			"fields": utils.ValidatorErrors(err),
		})
	}

	u, err := h.service.CreateUser(c.UserContext(), user.CreateUserInput{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  payload.Password,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error": false,
		"msg":   "user created successfully",
		"user":  mapUserToResponse(u),
	})
}

func mapUsersToResponse(users []user.User) []dto.UserResponse {
	resp := make([]dto.UserResponse, 0, len(users))
	for _, u := range users {
		resp = append(resp, mapUserToResponse(u))
	}
	return resp
}

func mapUserToResponse(u user.User) dto.UserResponse {
	return dto.UserResponse{
		ID:        u.ID,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
	}
}
