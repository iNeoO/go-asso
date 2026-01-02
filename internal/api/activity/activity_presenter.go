package activityapi

import commonapi "github.com/ineoo/go-planigramme/internal/api/common"

type ErrorEnvelope = commonapi.ErrorEnvelope

func ActivityErrorResponse(message string) ErrorEnvelope {
	return ErrorEnvelope{
		Status: false,
		Data:   nil,
		Error:  message,
	}
}
