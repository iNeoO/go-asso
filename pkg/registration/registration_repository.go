package registration

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"github.com/ineoo/go-planigramme/pkg/activity"
)

type Repository struct {
	db *sqlx.DB
}

type Registration struct {
	UserID     uuid.UUID `db:"user_id"`
	StatusID   uuid.UUID `db:"status_id"`
	ActivityID uuid.UUID `db:"activity_id"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}

type RegisteredUser struct {
	ID        uuid.UUID `db:"id"`
	FirstName string    `db:"first_name"`
	LastName  string    `db:"last_name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(registration *Registration) (*Registration, error) {
	const query = `INSERT INTO registrations (user_id, status_id, activity_id) VALUES ($1, $2, $3) RETURNING *`
	if err := r.db.GetContext(context.Background(), registration, query, registration.UserID, registration.StatusID, registration.ActivityID); err != nil {
		return nil, err
	}
	return registration, nil
}
func (r *Repository) UpdateStatus(userID uuid.UUID, statusID uuid.UUID) (*Registration, error) {
	const query = `UPDATE registrations SET status_id=$1 WHERE user_id=$2 RETURNING *`
	registration := Registration{}
	if err := r.db.GetContext(context.Background(), &registration, query, statusID, userID); err != nil {
		return nil, err
	}
	return &registration, nil
}
func (r *Repository) GetByUserID(userID uuid.UUID) (*Registration, error) {
	const query = `SELECT * FROM registrations WHERE user_id=$1`
	var registration Registration
	if err := r.db.GetContext(context.Background(), &registration, query, userID); err != nil {
		return nil, err
	}
	return &registration, nil
}

func (r *Repository) GetRegistrationActivities(userID uuid.UUID) ([]RegistrationActivity, error) {
	const query = `
	SELECT
		r.user_id,
		r.status_id,
		r.activity_id,
		r.created_at,
		r.updated_at,
		a.id,
		a.name,
		a.creator_id,
		a.description,
		a.start_at,
		a.end_at,
		a.duration_minutes,
		a.capacity,
		a.organization_id,
		a.created_at AS activity_created_at,
		a.updated_at AS activity_updated_at
	FROM registrations r
	JOIN activities a ON r.activity_id = a.id
	WHERE r.user_id = $1
	`

	var rows = make([]struct {
		ID                uuid.UUID `db:"id"`
		UserID            uuid.UUID `db:"user_id"`
		StatusID          uuid.UUID `db:"status_id"`
		ActivityID        uuid.UUID `db:"activity_id"`
		CreatedAt         time.Time `db:"created_at"`
		UpdatedAt         time.Time `db:"updated_at"`
		Name              string    `db:"name"`
		CreatorID         uuid.UUID `db:"creator_id"`
		Description       string    `db:"description"`
		StartAt           time.Time `db:"start_at"`
		EndAt             time.Time `db:"end_at"`
		DurationMinutes   int       `db:"duration_minutes"`
		Capacity          int       `db:"capacity"`
		OrganizationID    uuid.UUID `db:"organization_id"`
		ActivityCreatedAt time.Time `db:"activity_created_at"`
		ActivityUpdatedAt time.Time `db:"activity_updated_at"`
	}, 0)

	if err := r.db.GetContext(context.Background(), &rows, query, userID); err != nil {
		return nil, err
	}

	RegistrationActivitys := make([]RegistrationActivity, 0, len(rows))
	for _, row := range rows {
		RegistrationActivitys = append(RegistrationActivitys, RegistrationActivity{
			UserID:     row.UserID,
			StatusID:   row.StatusID,
			ActivityID: row.ActivityID,
			Activity: activity.Activity{
				ID:              row.ID,
				Name:            row.Name,
				CreatorID:       row.CreatorID,
				Description:     row.Description,
				StartAt:         row.StartAt,
				EndAt:           row.EndAt,
				DurationMinutes: row.DurationMinutes,
				Capacity:        row.Capacity,
				OrganizationID:  row.OrganizationID,
				CreatedAt:       row.ActivityCreatedAt,
				UpdatedAt:       row.ActivityUpdatedAt,
			},
		})
	}

	return RegistrationActivitys, nil
}

func (r *Repository) GetUsersByActivity(activityID uuid.UUID) ([]RegisteredUser, error) {
	const query = `
	SELECT 
		u.id,
		u.first_name,
		u.last_name,
		r.created_at,
		r.updated_at
	FROM users u
	JOIN registrations r ON u.id = r.user_id
	WHERE r.activity_id = $1
	`
	users := make([]RegisteredUser, 0)

	if err := r.db.SelectContext(context.Background(), &users, query, activityID); err != nil {
		return nil, err
	}
	return users, nil
}
