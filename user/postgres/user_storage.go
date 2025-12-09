package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/ineoo/go-planigramme/user"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) List(ctx context.Context) ([]user.User, error) {
	const query = `SELECT * FROM users`
	var users []user.User
	if err := r.db.SelectContext(ctx, &users, query); err != nil {
		return nil, err
	}
	return users, nil
}

func (r *Repository) Get(ctx context.Context, id uuid.UUID) (user.User, error) {
	const query = `SELECT * FROM users WHERE id=$1`
	var u user.User
	if err := r.db.GetContext(ctx, &u, query, id); err != nil {
		return user.User{}, err
	}
	return u, nil
}

func (r *Repository) Create(ctx context.Context, u *user.User) error {
	const query = `INSERT INTO users (first_name, last_name, email, password_hash) VALUES ($1, $2, $3, $4) RETURNING *`
	return r.db.GetContext(ctx, u, query, u.FirstName, u.LastName, u.Email, u.PasswordHash)
}

var _ user.Repository = (*Repository)(nil)
