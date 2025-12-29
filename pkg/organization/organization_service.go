package organization

import (
	"github.com/google/uuid"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Create(organization *Organization) (*Organization, error) {
	return s.repo.Create(organization)
}

func (s *Service) List() ([]Organization, error) {
	return s.repo.List()
}

func (s *Service) ListByUserId(userId uuid.UUID) ([]Organization, error) {
	return s.repo.ListByUserId(userId)
}

func (s *Service) GetByID(id uuid.UUID) (*Organization, error) {
	return s.repo.GetByID(id)
}

func (s *Service) Update(organization *Organization) (*Organization, error) {
	return s.repo.Update(organization)
}

func (s *Service) Delete(id uuid.UUID) error {
	return s.repo.Delete(id)
}
