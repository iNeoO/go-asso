package user

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type User struct {
	ID              uuid.UUID `db:"id"`
	CreatedAt       time.Time `db:"created_at"`
	UpdatedAt       time.Time `db:"updated_at"`
	FirstName       string    `db:"first_name"`
	LastName        string    `db:"last_name"`
	Email           string    `db:"email"`
	IsEmailVerified bool      `db:"is_email_verified" default:"false"`
	PasswordHash    string    `db:"password_hash"`
}

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetById(id uuid.UUID) (*User, error) {
	const query = `SELECT * FROM users WHERE id=$1`
	var u User
	if err := r.db.GetContext(context.Background(), &u, query, id); err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *Repository) GetByEmail(email string) (*User, error) {
	const query = `SELECT * FROM users WHERE email=$1`
	var u User
	if err := r.db.GetContext(context.Background(), &u, query, email); err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *Repository) List() (*[]User, error) {
	const query = `SELECT * FROM users`
	users := make([]User, 0)
	if err := r.db.SelectContext(context.Background(), &users, query); err != nil {
		return nil, err
	}
	return &users, nil
}

func (r *Repository) Create(u *User) (*User, error) {
	const query = `INSERT INTO users (first_name, last_name, email, password_hash) VALUES ($1, $2, $3, $4) RETURNING *`
	if err := r.db.GetContext(context.Background(), u, query, u.FirstName, u.LastName, u.Email, u.PasswordHash); err != nil {
		return nil, err
	}
	return u, nil
}

func (r *Repository) Update(u *User) (*User, error) {
	return u, nil
}

func (r *Repository) Delete(id uuid.UUID) error {
	return nil
}
