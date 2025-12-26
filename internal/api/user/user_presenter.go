package userapi

import (
	"time"

	"github.com/google/uuid"

	commonapi "github.com/ineoo/go-planigramme/internal/api/common"
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

// @name UserEnvelope
type UserEnvelope struct {
	Status bool    `json:"status"`
	Data   User    `json:"data"`
	Error  *string `json:"error"`
}

// @name UsersEnvelope
type UsersEnvelope struct {
	Status bool    `json:"status"`
	Data   []User  `json:"data"`
	Count  int     `json:"count"`
	Error  *string `json:"error"`
}

type UserErrorEnvelope = commonapi.ErrorEnvelope

func UserSuccessResponse(data *userdomain.User) UserEnvelope {
	user := User{
		ID:        data.ID,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Email:     data.Email,
	}
	return UserEnvelope{
		Status: true,
		Data:   user,
		Error:  nil,
	}
}

func UsersSuccessResponse(data *[]userdomain.User) UsersEnvelope {
	users := make([]User, 0, len(*data))
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

	return UsersEnvelope{
		Status: true,
		Data:   users,
		Count:  len(users),
		Error:  nil,
	}
}

func UserErrorResponse(err error) UserErrorEnvelope {
	return UserErrorEnvelope{
		Status: false,
		Data:   nil,
		Error:  err.Error(),
	}
}
