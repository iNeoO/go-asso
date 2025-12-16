package session

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Session struct {
	ID        uuid.UUID `db:"id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	ExpiresAt	time.Time `db:"expires_at"`
	Token     string    `db:"token"`
	IsRevoked bool      `db:"is_revoked" default:"false"`
	UserID    uuid.UUID `db:"user_id"`
}

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetByToken(token string) (*Session, error) {
	const query = `SELECT * FROM sessions WHERE token=$1`
	var s Session
	if err := r.db.GetContext(context.Background(), &s, query, token); err != nil {
		return nil, err
	}
	return &s, nil
}

func (r *Repository) Create(s *Session) (*Session, error) {
	const query = `INSERT INTO sessions (token, user_id, expires_at) VALUES ($1, $2, $3) RETURNING *`
	if err := r.db.GetContext(context.Background(), s, query, s.Token, s.UserID, s.ExpiresAt); err != nil {
		return nil, err
	}
	return s, nil
}

func (r *Repository) Update(s *Session) (*Session, error) {
	const query = `UPDATE sessions SET is_revoked=$1, updated_at=$2 WHERE id=$3 RETURNING *`
	s.UpdatedAt = time.Now()
	if err := r.db.GetContext(context.Background(), s, query, s.IsRevoked, s.UpdatedAt, s.ID); err != nil {
		return nil, err
	}
	return s, nil
}

func (r *Repository) DeleteByToken(token string) error {
	const query = `DELETE FROM sessions WHERE token=$1`
	_, err := r.db.ExecContext(context.Background(), query, token)
	return err
}