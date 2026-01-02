package membership

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"github.com/ineoo/go-planigramme/pkg/organization"
)

type OrganizationUser struct {
	UserID         uuid.UUID           `db:"user_id"`
	OrganizationID uuid.UUID           `db:"organization_id"`
	RoleID         organization.RoleID `db:"role_id"`
	CreatedAt      time.Time           `db:"created_at"`
	UpdatedAt      time.Time           `db:"updated_at"`
}

type OrganizationMemberRecord struct {
	UserID    uuid.UUID           `db:"user_id"`
	RoleID    organization.RoleID `db:"role_id"`
	CreatedAt time.Time           `db:"created_at"`
	UpdatedAt time.Time           `db:"updated_at"`
	FirstName string              `db:"first_name"`
	LastName  string              `db:"last_name"`
	Email     string              `db:"email"`
}

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Join(id uuid.UUID, userID uuid.UUID, roleID organization.RoleID) (*OrganizationUser, error) {
	const query = `INSERT INTO user_organizations (user_id, organization_id, role_id) VALUES ($1, $2, $3) RETURNING *`
	organizationUser := OrganizationUser{}
	if err := r.db.GetContext(context.Background(), &organizationUser, query, userID, id, roleID); err != nil {
		return nil, err
	}
	return &organizationUser, nil
}

func (r *Repository) AssignRole(id uuid.UUID, userID uuid.UUID, roleID organization.RoleID) (*OrganizationUser, error) {
	const query = `UPDATE user_organizations SET role_id=$1, updated_at=$2 WHERE user_id=$3 AND organization_id=$4 RETURNING *`
	updatedAt := time.Now()
	organizationUser := OrganizationUser{}
	if err := r.db.GetContext(context.Background(), &organizationUser, query, roleID, updatedAt, userID, id); err != nil {
		return nil, err
	}
	return &organizationUser, nil
}

func (r *Repository) GetOrganizationUser(id uuid.UUID, userID uuid.UUID) (*OrganizationUser, error) {
	const query = `SELECT * FROM user_organizations WHERE user_id=$1 AND organization_id=$2`
	organizationUser := OrganizationUser{}
	if err := r.db.GetContext(context.Background(), &organizationUser, query, userID, id); err != nil {
		return nil, err
	}
	return &organizationUser, nil
}

func (r *Repository) GetOrganizationMember(userID uuid.UUID, organizationID uuid.UUID) (*OrganizationMember, error) {
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
	WHERE uo.organization_id = $1 AND uo.user_id = $2
	`

	organizationMember := OrganizationMemberRecord{}
	if err := r.db.GetContext(context.Background(), &organizationMember, query, organizationID, userID); err != nil {
		return nil, err
	}
	return &OrganizationMember{
		UserID:    organizationMember.UserID,
		RoleID:    organizationMember.RoleID,
		CreatedAt: organizationMember.CreatedAt,
		UpdatedAt: organizationMember.UpdatedAt,
		User: OrganizationMemberUser{
			FirstName: organizationMember.FirstName,
			LastName:  organizationMember.LastName,
			Email:     organizationMember.Email,
		},
	}, nil
}

func (r *Repository) ListOrganizationMembers(id uuid.UUID) ([]OrganizationMember, error) {
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

	organizationUsers := make([]OrganizationMemberRecord, 0)
	if err := r.db.SelectContext(context.Background(), &organizationUsers, query, id); err != nil {
		return nil, err
	}

	members := make([]OrganizationMember, 0, len(organizationUsers))
	for _, ou := range organizationUsers {
		members = append(members, OrganizationMember{
			UserID:    ou.UserID,
			RoleID:    ou.RoleID,
			CreatedAt: ou.CreatedAt,
			UpdatedAt: ou.UpdatedAt,
			User: OrganizationMemberUser{
				FirstName: ou.FirstName,
				LastName:  ou.LastName,
				Email:     ou.Email,
			},
		})
	}

	return members, nil
}

func (r *Repository) Leave(id uuid.UUID, userID uuid.UUID) error {
	const query = `DELETE FROM user_organizations WHERE user_id=$1 AND organization_id=$2`
	_, err := r.db.ExecContext(context.Background(), query, userID, id)
	return err
}

func (r *Repository) ListOrganizationsForUser(userID uuid.UUID) ([]OrganizationUser, error) {
	const query = `SELECT user_id, organization_id, role_id, created_at, updated_at FROM user_organizations WHERE user_id = $1`
	organizationUsers := make([]OrganizationUser, 0)
	if err := r.db.SelectContext(context.Background(), &organizationUsers, query, userID); err != nil {
		return nil, err
	}
	return organizationUsers, nil
}
