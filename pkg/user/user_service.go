package user

import (
	"time"

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

func (r *Service) Create(user *User) (*User, error) {
	return r.repo.Create(user)
}

func (r *Service) List() (*[]User, error) {
	return r.repo.List()
}

func (r *Service) GetById(id uuid.UUID) (*User, error) {
	return r.repo.GetById(id)
}

func (r *Service) GetByEmail(email string) (*User, error) {
	return r.repo.GetByEmail(email)
}

func (r *Service) Update(user *User) (*User, error) {
	user.UpdatedAt = time.Now()
	return r.repo.Update(user)
}

func (r *Service) Delete(id uuid.UUID) error {
	return r.repo.Delete(id)
}