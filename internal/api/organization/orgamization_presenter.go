package organizationapi

import (
	"time"

	"github.com/google/uuid"
	commonapi "github.com/ineoo/go-planigramme/internal/api/common"
	"github.com/ineoo/go-planigramme/pkg/membership"
	"github.com/ineoo/go-planigramme/pkg/organization"
)

type Organization struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
}

type OrganizationUser struct {
	UserID         uuid.UUID           `json:"user_id"`
	OrganizationID uuid.UUID           `json:"organization_id"`
	RoleID         organization.RoleID `json:"role_id"`
	CreatedAt      time.Time           `json:"created_at"`
	UpdatedAt      time.Time           `json:"updated_at"`
}

// @name OrganizationEnvelope
type OrganizationEnvelope struct {
	Status bool         `json:"status"`
	Data   Organization `json:"data"`
	Error  *string      `json:"error"`
}

// @name OrganizationsEnvelope
type OrganizationsEnvelope struct {
	Status bool           `json:"status"`
	Data   []Organization `json:"data"`
	Count  int            `json:"count"`
	Error  *string        `json:"error"`
}

// @name OrganizationUserEnvelope
type OrganizationUserEnvelope struct {
	Status bool             `json:"status"`
	Data   OrganizationUser `json:"data"`
	Error  *string          `json:"error"`
}

type OrganizationUsersEnvelope struct {
	Status bool               `json:"status"`
	Data   []OrganizationUser `json:"data"`
	Count  int                `json:"count"`
	Error  *string            `json:"error"`
}

type OrganizationMembersEnvelope struct {
	Status bool                            `json:"status"`
	Data   []membership.OrganizationMember `json:"data"`
	Count  int                             `json:"count"`
	Error  *string                         `json:"error"`
}

// @name ErrorEnvelope
type ErrorEnvelope = commonapi.ErrorEnvelope

func OrganizationSuccessResponse(organization *organization.Organization) OrganizationEnvelope {
	org := Organization{
		ID:        organization.ID,
		CreatedAt: organization.CreatedAt,
		UpdatedAt: organization.UpdatedAt,
		Name:      organization.Name,
	}
	return OrganizationEnvelope{
		Status: true,
		Data:   org,
		Error:  nil,
	}
}

func OrganizationsSuccessResponse(organizations []organization.Organization) OrganizationsEnvelope {
	orgs := make([]Organization, 0, len(organizations))
	for _, o := range organizations {
		org := Organization{
			ID:        o.ID,
			CreatedAt: o.CreatedAt,
			UpdatedAt: o.UpdatedAt,
			Name:      o.Name,
		}
		orgs = append(orgs, org)
	}

	return OrganizationsEnvelope{
		Status: true,
		Data:   orgs,
		Count:  len(orgs),
		Error:  nil,
	}
}

func OrganizationUserSuccessResponse(organizationUser *membership.OrganizationUser) OrganizationUserEnvelope {
	orgUser := OrganizationUser{
		UserID:         organizationUser.UserID,
		OrganizationID: organizationUser.OrganizationID,
		RoleID:         organizationUser.RoleID,
		CreatedAt:      organizationUser.CreatedAt,
		UpdatedAt:      organizationUser.UpdatedAt,
	}
	return OrganizationUserEnvelope{
		Status: true,
		Data:   orgUser,
		Error:  nil,
	}
}

func OrganizationErrorResponse(message string) ErrorEnvelope {
	return ErrorEnvelope{
		Status: false,
		Data:   nil,
		Error:  message,
	}
}

func OrganizationMembersSuccessResponse(organizationUsers []membership.OrganizationMember) OrganizationMembersEnvelope {

	return OrganizationMembersEnvelope{
		Status: true,
		Data:   organizationUsers,
		Count:  len(organizationUsers),
		Error:  nil,
	}
}
