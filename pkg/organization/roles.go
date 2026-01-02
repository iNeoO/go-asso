package organization

type RoleID string

const (
	RoleCreator       RoleID = "CREATOR"
	RoleAdministrator RoleID = "ADMINISTRATOR"
	RoleTeamMember    RoleID = "TEAM_MEMBER"
	RoleValidated     RoleID = "VALIDATED"
	RoleNotValidated  RoleID = "NOT_VALIDATED"
)

func HasReadAccess(roleID RoleID) bool {
	switch roleID {
	case RoleCreator, RoleAdministrator, RoleTeamMember, RoleValidated:
		return true
	default:
		return false
	}
}

func HasWriteAccess(roleID RoleID) bool {
	switch roleID {
	case RoleCreator, RoleAdministrator, RoleTeamMember:
		return true
	default:
		return false
	}
}

func HasAdminAccess(roleID RoleID) bool {
	switch roleID {
	case RoleCreator, RoleAdministrator:
		return true
	default:
		return false
	}
}
