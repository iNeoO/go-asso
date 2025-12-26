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

type OrganizationUser struct {
	UserID         uuid.UUID `db:"user_id"`
	OrganizationID uuid.UUID `db:"organization_id"`
	RoleID         RoleID    `db:"role_id"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}

type OrganizationMember struct {
	UserID    uuid.UUID `db:"user_id"`
	RoleID    RoleID    `db:"role_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	FirstName string    `db:"first_name"`
	LastName  string    `db:"last_name"`
	Email     string    `db:"email"`
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

func (r *Repository) Join(id uuid.UUID, user_id uuid.UUID, role_id RoleID) (*OrganizationUser, error) {
	const query = `INSERT INTO user_organizations (user_id, organization_id, role_id) VALUES ($1, $2, $3) RETURNING *`
	organizationUser := OrganizationUser{}
	if err := r.db.GetContext(context.Background(), &organizationUser, query, user_id, id, role_id); err != nil {
		return nil, err
	}
	return &organizationUser, nil
}

func (r *Repository) AssignRole(id uuid.UUID, user_id uuid.UUID, role_id RoleID) (*OrganizationUser, error) {
	const query = `UPDATE user_organizations SET role_id=$1, updated_at=$2 WHERE user_id=$3 AND organization_id=$4 RETURNING *`
	updatedAt := time.Now()
	organizationUser := OrganizationUser{}
	if err := r.db.GetContext(context.Background(), &organizationUser, query, role_id, updatedAt, user_id, id); err != nil {
		return nil, err
	}
	return &organizationUser, nil
}

func (r *Repository) GetOrganizationUser(id uuid.UUID, user_id uuid.UUID) (*OrganizationUser, error) {
	const query = `SELECT * FROM user_organizations WHERE user_id=$1 AND organization_id=$2`
	organizationsUser := OrganizationUser{}
	if err := r.db.GetContext(context.Background(), &organizationsUser, query, user_id, id); err != nil {
		return nil, err
	}
	return &organizationsUser, nil
}

func (r *Repository) ListOrganizationUsers(id uuid.UUID) ([]OrganizationMember, error) {
	const query = `
	SELECT 
		uo.user_id,
		uo.role_id,
		uo.created_at,
		uo.updated_at,
		u.first_name,
		u.last_name,
		u.email
	FROM user_organizations uo
	JOIN users u ON u.id = uo.user_id
	WHERE uo.organization_id = $1
	`
	organizationUsers := make([]OrganizationMember, 0)
	if err := r.db.SelectContext(context.Background(), &organizationUsers, query, id); err != nil {
		return nil, err
	}
	return organizationUsers, nil
}

func (r *Repository) Leave(id uuid.UUID, user_id uuid.UUID) error {
	const query = `DELETE FROM user_organizations WHERE user_id=$1 AND organization_id=$2`
	_, err := r.db.ExecContext(context.Background(), query, user_id, id)
	return err
}

func (r *Repository) Delete(id uuid.UUID) error {
	const query = `DELETE FROM organizations WHERE id=$1`
	_, err := r.db.ExecContext(context.Background(), query, id)
	return err
}
