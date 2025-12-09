package dto

import (
	"time"

	"github.com/google/uuid"
)

// CreateUserRequest defines the payload required to create a user.
type CreateUserRequest struct {
	FirstName string `json:"first_name" validate:"required,lte=255"`
	LastName  string `json:"last_name" validate:"required,lte=255"`
	Email     string `json:"email" validate:"required,email,lte=255"`
	Password  string `json:"password" validate:"required,min=8,max=72"`
}

// UpdateUserRequest defines the payload allowed when updating a user profile.
type UpdateUserRequest struct {
	FirstName *string `json:"first_name,omitempty" validate:"omitempty,lte=255"`
	LastName  *string `json:"last_name,omitempty" validate:"omitempty,lte=255"`
	Email     *string `json:"email,omitempty" validate:"omitempty,email,lte=255"`
	Password  *string `json:"password,omitempty" validate:"omitempty,min=8,max=72"`
}

// UserResponse is what is serialized back to API clients.
type UserResponse struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
}
