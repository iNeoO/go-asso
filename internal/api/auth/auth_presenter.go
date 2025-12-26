package authapi

import commonapi "github.com/ineoo/go-planigramme/internal/api/common"

type AuthData struct {
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expires_at"`
}

// @name AuthEnvelope
type AuthEnvelope struct {
	Status bool      `json:"status"`
	Data   *AuthData `json:"data"`
	Error  *string   `json:"error"`
}

// @name AuthErrorEnvelope
type AuthErrorEnvelope = commonapi.ErrorEnvelope

func AuthErrorResponse(err error) AuthErrorEnvelope {
	return AuthErrorEnvelope{
		Status: false,
		Data:   nil,
		Error:  err.Error(),
	}
}

func AuthSuccessResponse(data *AuthData) AuthEnvelope {
	return AuthEnvelope{
		Status: true,
		Data:   data,
		Error:  nil,
	}
}
