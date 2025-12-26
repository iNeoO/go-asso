package organization

// RoleID represents a role identifier in the roles_enum table.
type RoleID string

const (
	RoleCreator       RoleID = "CREATOR"
	RoleAdministrator RoleID = "ADMINISTRATOR"
	RoleTeamMember    RoleID = "TEAM_MEMBER"
	RoleValidated     RoleID = "VALIDATED"
	RoleNotValidated  RoleID = "NOT_VALIDATED"
)
