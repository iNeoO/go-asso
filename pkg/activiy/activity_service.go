package activiy

import "github.com/google/uuid"

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetActivityByID(id string) (*activityWithCreator, error) {
	return s.repo.GetActivityByID(id)
}

func (s *Service) CreateActivity(act *activity) (*activity, error) {
	return s.repo.CreateActivity(act)
}

func (s *Service) DeleteActivity(id string) error {
	return s.repo.DeleteActivity(id)
}

func (s *Service) UpdateActivity(act *activity) (*activity, error) {
	return s.repo.UpdateActivity(act)
}

func (s *Service) ListActivitiesByOrganization(organizationId uuid.UUID) ([]activityWithCreator, error) {
	return s.repo.ListActivitiesByOrganization(organizationId)
}

func (s *Service) ListActivitiesByOrganizations(userID []uuid.UUID) ([]activityWithCreator, error) {
	return s.repo.ListActivitiesByOrganizations(userID)
}
