package membership

import (
	"time"

	"github.com/google/uuid"

	"github.com/ineoo/go-planigramme/pkg/organization"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

type OrganizationMemberUser struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type OrganizationMember struct {
	UserID    uuid.UUID              `json:"user_id"`
	RoleID    organization.RoleID    `json:"role_id"`
	CreatedAt time.Time              `json:"created_at"`
	UpdatedAt time.Time              `json:"updated_at"`
	User      OrganizationMemberUser `json:"user,omitempty"`
}

func (s *Service) Join(id uuid.UUID, userID uuid.UUID, roleID organization.RoleID) (*OrganizationUser, error) {
	return s.repo.Join(id, userID, roleID)
}

func (s *Service) AssignRole(id uuid.UUID, userID uuid.UUID, roleID organization.RoleID) (*OrganizationUser, error) {
	return s.repo.AssignRole(id, userID, roleID)
}

func (s *Service) GetOrganizationUser(id uuid.UUID, userID uuid.UUID) (*OrganizationUser, error) {
	return s.repo.GetOrganizationUser(id, userID)
}

func (s *Service) ListOrganizationMembers(id uuid.UUID) ([]OrganizationMember, error) {
	return s.repo.ListOrganizationMembers(id)
}

func (s *Service) GetOrganizationMember(userID uuid.UUID, organizationID uuid.UUID) (*OrganizationMember, error) {
	return s.repo.GetOrganizationMember(userID, organizationID)
}

func (s *Service) Leave(id uuid.UUID, userID uuid.UUID) error {
	return s.repo.Leave(id, userID)
}

func (s *Service) ListOrganizationsForUser(userID uuid.UUID) ([]OrganizationUser, error) {
	return s.repo.ListOrganizationsForUser(userID)
}
