package organizationapi

import (
	"github.com/google/uuid"
	organizationdomain "github.com/ineoo/go-planigramme/pkg/organization"
)

type CreateOrganizationRequest struct {
	Name string `json:"name" validate:"required,lte=255"`
}

type UpdateOrganizationUserRoleRequest struct {
	UserID uuid.UUID                 `json:"user_id" validate:"required"`
	RoleID organizationdomain.RoleID `json:"role_id" validate:"required,oneof=CREATOR ADMINISTRATOR TEAM_MEMBER VALIDATED NOT_VALIDATED"`
}
