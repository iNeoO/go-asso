package meapi

import (
	commonapi "github.com/ineoo/go-planigramme/internal/api/common"
	organizationapi "github.com/ineoo/go-planigramme/internal/api/organization"
	userapi "github.com/ineoo/go-planigramme/internal/api/user"
	organizationdomain "github.com/ineoo/go-planigramme/pkg/organization"
	userdomain "github.com/ineoo/go-planigramme/pkg/user"
)

// @name MeErrorEnvelope
type MeErrorEnvelope = commonapi.ErrorEnvelope

// @name MeProfileEnvelope
type MeProfileEnvelope = userapi.UserEnvelope

// @name MeOrganizationEnvelope
type MeOrganizationsEnvelope = organizationapi.OrganizationsEnvelope

func MeErrorResponse(message string) MeErrorEnvelope {
	return MeErrorEnvelope{
		Status: false,
		Data:   nil,
		Error:  message,
	}
}

func MeProfileSuccessResponse(data *userdomain.User) MeProfileEnvelope {
	return userapi.UserSuccessResponse(data)
}

func MeOrganizationSuccessResponse(data []organizationdomain.Organization) MeOrganizationsEnvelope {
	return organizationapi.OrganizationsSuccessResponse(data)
}
