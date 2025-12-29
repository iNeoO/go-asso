package userapi

type CreateUserRequest struct {
	FirstName string `json:"first_name" validate:"required,lte=255"`
	LastName  string `json:"last_name" validate:"required,lte=255"`
	Email     string `json:"email" validate:"required,email,lte=255"`
	Password  string `json:"password" validate:"required,min=8,max=72"`
}
