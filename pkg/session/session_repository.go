package session

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Session struct {
	ID        uuid.UUID  `db:"id"`
	CreatedAt time.Time  `db:"created_at"`
	ExpiresAt time.Time  `db:"expires_at"`
	RevokedAt *time.Time `db:"revoked_at"`
	TokenHash string     `db:"token_hash"`
	UserID    uuid.UUID  `db:"user_id"`
}

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetByToken(token string) (*Session, error) {
	const query = `SELECT * FROM sessions WHERE token_hash=$1`
	var s Session
	if err := r.db.GetContext(context.Background(), &s, query, hashToken(token)); err != nil {
		return nil, err
	}
	return &s, nil
}

func (r *Repository) Create(s *Session) (*Session, error) {
	const query = `INSERT INTO sessions (token_hash, user_id, expires_at) VALUES ($1, $2, $3) RETURNING *`
	s.TokenHash = hashToken(s.TokenHash)
	if err := r.db.GetContext(context.Background(), s, query, s.TokenHash, s.UserID, s.ExpiresAt); err != nil {
		return nil, err
	}
	return s, nil
}

func (r *Repository) Update(s *Session) (*Session, error) {
	const query = `UPDATE sessions SET revoked_at=$1 WHERE id=$2 RETURNING *`
	if s.RevokedAt == nil {
		now := time.Now()
		s.RevokedAt = &now
	}
	if err := r.db.GetContext(context.Background(), s, query, s.RevokedAt, s.ID); err != nil {
		return nil, err
	}
	return s, nil
}

func (r *Repository) DeleteByToken(token string) error {
	const query = `DELETE FROM sessions WHERE token_hash=$1`
	_, err := r.db.ExecContext(context.Background(), query, hashToken(token))
	return err
}

func hashToken(token string) string {
	sum := sha256.Sum256([]byte(token))
	return hex.EncodeToString(sum[:])
}
