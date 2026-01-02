package activity

import (
	"github.com/google/uuid"
	membershipdomain "github.com/ineoo/go-planigramme/pkg/membership"
	"github.com/ineoo/go-planigramme/pkg/organization"
)

type Service struct {
	repo       Repository
	membership membershipdomain.Service
}

func NewService(repo Repository, membershipService membershipdomain.Service) *Service {
	return &Service{
		repo:       repo,
		membership: membershipService,
	}
}

func (s *Service) GetActivityByID(id uuid.UUID) (*ActivityWithCreator, error) {
	return s.repo.GetActivityByID(id)
}

func (s *Service) CreateActivity(act *ActivityCreation) (*Activity, error) {
	return s.repo.CreateActivity(act)
}

func (s *Service) DeleteActivity(id uuid.UUID) error {
	return s.repo.DeleteActivity(id)
}

func (s *Service) UpdateActivity(act *Activity) (*Activity, error) {
	return s.repo.UpdateActivity(act)
}

func (s *Service) ListActivitiesByOrganization(organizationId uuid.UUID) ([]ActivityWithCreator, error) {
	return s.repo.ListActivitiesByOrganization(organizationId)
}

func (s *Service) ListActivitiesByOrganizations(userID []uuid.UUID) ([]ActivityWithCreator, error) {
	return s.repo.ListActivitiesByOrganizations(userID)
}

func (s *Service) ListActivitiesForUser(userId uuid.UUID) ([]ActivityWithCreator, error) {
	organizationUsers, err := s.membership.ListOrganizationsForUser(userId)
	if err != nil {
		return nil, err
	}

	orgIDs := make([]uuid.UUID, 0, len(organizationUsers))
	for _, ou := range organizationUsers {
		if organization.HasReadAccess(ou.RoleID) {
			orgIDs = append(orgIDs, ou.OrganizationID)
		}
	}

	if len(orgIDs) == 0 {
		return []ActivityWithCreator{}, nil
	}

	return s.repo.ListActivitiesByOrganizations(orgIDs)
}
