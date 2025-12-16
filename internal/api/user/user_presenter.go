package userapi

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	userdomain "github.com/ineoo/go-planigramme/pkg/user"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
}

func UserSuccessResponse(data *userdomain.User) *fiber.Map {
	user := User{
		ID:        data.ID,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Email:     data.Email,
	}
	return &fiber.Map{
		"status": true,
		"data":   user,
		"error":  nil,
	}
}

func UsersSuccessResponse(data *[]userdomain.User) *fiber.Map {
	var users []User
	for _, u := range *data {
		user := User{
			ID:        u.ID,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Email:     u.Email,
		}
		users = append(users, user)
	}

	return &fiber.Map{
		"status": true,
		"data":   users,
		"error":  nil,
		"count":  len(*data),
	}
}

func UserErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   nil,
		"error":  err.Error(),
	}
}
