package registration

import (
	"github.com/google/uuid"
	"github.com/ineoo/go-planigramme/pkg/activity"
)

type Service struct {
	repo Repository
}

type RegistrationActivity struct {
	UserID     uuid.UUID         `json:"user_id"`
	StatusID   uuid.UUID         `json:"status_id"`
	ActivityID uuid.UUID         `json:"activity_id"`
	Activity   activity.Activity `json:"activity"`
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(reg *Registration) (*Registration, error) {
	return s.repo.Create(reg)
}

func (s *Service) UpdateStatus(userID uuid.UUID, statusID uuid.UUID) (*Registration, error) {
	return s.repo.UpdateStatus(userID, statusID)
}

func (s *Service) GetByUserID(userID uuid.UUID) (*Registration, error) {
	return s.repo.GetByUserID(userID)
}

func (s *Service) GetRegistrationActivities(userID uuid.UUID) ([]RegistrationActivity, error) {
	return s.repo.GetRegistrationActivities(userID)
}

func (s *Service) GetUsersByActivity(activityID uuid.UUID) ([]RegisteredUser, error) {
	return s.repo.GetUsersByActivity(activityID)
}
