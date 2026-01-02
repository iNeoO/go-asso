package organization

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Organization struct {
	ID        uuid.UUID `db:"id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Name      string    `db:"name"`
}

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(o *Organization) (*Organization, error) {
	const query = `INSERT INTO organizations (name) VALUES ($1) RETURNING *`
	if err := r.db.GetContext(context.Background(), o, query, o.Name); err != nil {
		return nil, err
	}
	return o, nil
}

func (r *Repository) GetByID(id uuid.UUID) (*Organization, error) {
	const query = `SELECT * FROM organizations WHERE id=$1`
	var o Organization
	if err := r.db.GetContext(context.Background(), &o, query, id); err != nil {
		return nil, err
	}
	return &o, nil
}

func (r *Repository) List() ([]Organization, error) {
	const query = `SELECT * FROM organizations`
	organizations := make([]Organization, 0)
	if err := r.db.SelectContext(context.Background(), &organizations, query); err != nil {
		return nil, err
	}
	return organizations, nil
}

func (r *Repository) ListByUserIdWithRoles(id uuid.UUID, roles []RoleID) ([]uuid.UUID, error) {
	const query = `
	SELECT o.id FROM organizations o
	JOIN user_organizations uo ON o.id = uo.organization_id
	WHERE uo.user_id = $1 AND uo.role_id IN $2
	`
	organizationIds := make([]uuid.UUID, 0)
	if err := r.db.SelectContext(context.Background(), &organizationIds, query, id, roles); err != nil {
		return nil, err
	}
	return organizationIds, nil
}

func (r *Repository) ListByUserId(userId uuid.UUID) ([]Organization, error) {
	const query = `
	SELECT o.* FROM organizations o
	JOIN user_organizations uo ON o.id = uo.organization_id
	WHERE uo.user_id = $1
	`
	organizations := make([]Organization, 0)
	if err := r.db.SelectContext(context.Background(), &organizations, query, userId); err != nil {
		return nil, err
	}
	return organizations, nil
}

func (r *Repository) Update(o *Organization) (*Organization, error) {
	const query = `UPDATE organizations SET name=$1, updated_at=$2 WHERE id=$3 RETURNING *`
	o.UpdatedAt = time.Now()
	if err := r.db.GetContext(context.Background(), o, query, o.Name, o.UpdatedAt, o.ID); err != nil {
		return nil, err
	}
	return o, nil
}

func (r *Repository) Delete(id uuid.UUID) error {
	const query = `DELETE FROM organizations WHERE id=$1`
	_, err := r.db.ExecContext(context.Background(), query, id)
	return err
}
