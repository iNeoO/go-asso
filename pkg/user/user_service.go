package user

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

func (s *Service) Create(user *User) (*User, error) {
	return s.repo.Create(user)
}

func (s *Service) List() (*[]User, error) {
	return s.repo.List()
}

func (s *Service) GetById(id uuid.UUID) (*User, error) {
	return s.repo.GetById(id)
}

func (s *Service) GetByEmail(email string) (*User, error) {
	return s.repo.GetByEmail(email)
}

func (s *Service) Update(user *User) (*User, error) {
	return s.repo.Update(user)
}

func (s *Service) Delete(id uuid.UUID) error {
	return s.repo.Delete(id)
}
