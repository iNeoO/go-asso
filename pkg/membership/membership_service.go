package membership

import (
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

func (s *Service) Join(id uuid.UUID, userID uuid.UUID, roleID organization.RoleID) (*OrganizationUser, error) {
	return s.repo.Join(id, userID, roleID)
}

func (s *Service) AssignRole(id uuid.UUID, userID uuid.UUID, roleID organization.RoleID) (*OrganizationUser, error) {
	return s.repo.AssignRole(id, userID, roleID)
}

func (s *Service) GetOrganizationUser(id uuid.UUID, userID uuid.UUID) (*OrganizationUser, error) {
	return s.repo.GetOrganizationUser(id, userID)
}

func (s *Service) ListOrganizationUsers(id uuid.UUID) ([]OrganizationMember, error) {
	return s.repo.ListOrganizationUsers(id)
}

func (s *Service) Leave(id uuid.UUID, userID uuid.UUID) error {
	return s.repo.Leave(id, userID)
}
