package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/ineoo/go-planigramme/app/models"
	"github.com/ineoo/go-planigramme/plateform/database"
)

// GetUsers func gets all exists users.
// @Description Get all exists users.
// @Summary get all exists users
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {array} models.User
// @Router /v1/books [get]
func GetUsers(c *fiber.Ctx) error {
	db, err := database.OpenDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
		})
	}

	users, err := db.UserQueries.GetUsers()
	if (err != nil) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "users were not found",
			"count": 0,
			"users": nil,
		})
	}
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   "",
		"count": len(users),
		"users": users,
	})
}

// GetUser func gets user by given ID or 404 error.
// @Description Get user by given ID.
// @Summary get user by given ID
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} models.User
// @Router /v1/user/{id} [get]
func GetUser(c *fiber.Ctx) error {

	id, err := uuid.Parse(c.Params("id"))
    if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
		})
	}

	db, err := database.OpenDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	user, err := db.UserQueries.GetUserByID(id)
	if (err != nil) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "user was not found",
			"user":  nil,
		})
	}
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   "user was found",
		"user":  user,
	})
}

// CreateUser func for creates a new user.
// @Description Create a new user.
// @Summary create a new user
// @Tags User
// @Accept json
// @Produce json
// @Param title body string true "Title"
// @Param author body string true "Author"
// @Param user_attrs body models.UserAttrs true "User attributes"
// @Success 200 {object} models.User
// @Security ApiKeyAuth
// @Router /v1/user [post]
func CreateUser(c *fiber.Ctx) error {
	db, err := database.OpenDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Parse the request body into a User struct
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "invalid request body",
		})
	}

	err = db.CreateUser(*user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "could not create user",
		})
	}

	return c.JSON(fiber.Map{
		"error": false,
		"msg":   "user created successfully",
		"user":  user,
	})
}

// UpdateUser func for updates an existing user.
// @Description Update an existing user.
// @Summary update an existing user
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user_attrs body models.UserAttrs true "User attributes"
// @Success 200 {object} models.User
// @Security ApiKeyAuth
// @Router /v1/user/{id} [put]
func UpdateUser(c *fiber.Ctx) error {
	return c.SendString("UpdateUser")
}

// DeleteUser func for deletes an existing user.
// @Description Delete an existing user.
// @Summary delete an existing user
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {string} string "User deleted successfully"
// @Security ApiKeyAuth
// @Router /v1/user/{id} [delete]
func DeleteUser(c *fiber.Ctx) error {
	return c.SendString("DeleteUser")
}