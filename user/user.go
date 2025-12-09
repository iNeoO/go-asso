package user

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `db:"id"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
	FirstName    string    `db:"first_name"`
	LastName     string    `db:"last_name"`
	Email        string    `db:"email"`
	PasswordHash string    `db:"password_hash"`
}

type CreateUserInput struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

type Repository interface {
	List(ctx context.Context) ([]User, error)
	Get(ctx context.Context, id uuid.UUID) (User, error)
	Create(ctx context.Context, u *User) error
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) ListUsers(ctx context.Context) ([]User, error) {
	return s.repo.List(ctx)
}

func (s *Service) GetUser(ctx context.Context, id uuid.UUID) (User, error) {
	return s.repo.Get(ctx, id)
}

func (s *Service) CreateUser(ctx context.Context, input CreateUserInput) (User, error) {
	if input.FirstName == "" || input.LastName == "" {
		return User{}, errors.New("missing user name")
	}
	if input.Email == "" {
		return User{}, errors.New("missing email")
	}

	user := User{
		FirstName:    input.FirstName,
		LastName:     input.LastName,
		Email:        strings.ToLower(input.Email),
		PasswordHash: input.Password,
	}

	if err := s.repo.Create(ctx, &user); err != nil {
		return User{}, err
	}

	return user, nil
}
