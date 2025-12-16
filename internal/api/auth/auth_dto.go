package authapi

type LoginRequest struct {
	Email string `json:"username" validate:"required,lte=255"`
	Password string `json:"password" validate:"required,min=8,max=72"`
}