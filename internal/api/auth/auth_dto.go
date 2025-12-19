package authapi

type LoginRequest struct {
	Email string `json:"email" validate:"required,lte=255"`
	Password string `json:"password" validate:"required,min=8,max=72"`
}
