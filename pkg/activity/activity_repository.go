package activity

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type activity struct {
	ID              uuid.UUID `db:"id"`
	Name            string    `db:"name"`
	CreatorID       uuid.UUID `db:"creator_id"`
	Description     string    `db:"description"`
	StartAt         time.Time `db:"start_at"`
	EndAt           time.Time `db:"end_at"`
	DurationMinutes int       `db:"duration_minutes"`
	Capacity        int       `db:"capacity"`
	CreatedAt       time.Time `db:"created_at"`
	UpdatedAt       time.Time `db:"updated_at"`
}

type Activity struct {
	ID              uuid.UUID `db:"id"`
	Name            string    `db:"name"`
	CreatorID       uuid.UUID `db:"creator_id"`
	Description     string    `db:"description"`
	StartAt         time.Time `db:"start_at"`
	EndAt           time.Time `db:"end_at"`
	DurationMinutes int       `db:"duration_minutes"`
	Capacity        int       `db:"capacity"`
	OrganizationID  uuid.UUID `db:"organization_id"`
	CreatedAt       time.Time `db:"created_at"`
	UpdatedAt       time.Time `db:"updated_at"`
}

type ActivityWithCreator struct {
	ID               uuid.UUID `db:"id"`
	OrganizationID   uuid.UUID `db:"organization_id"`
	Name             string    `db:"name"`
	CreatorID        uuid.UUID `db:"creator_id"`
	CreatorFirstName string    `db:"creator_first_name"`
	CreatorLastName  string    `db:"creator_last_name"`
	Description      string    `db:"description"`
	StartAt          time.Time `db:"start_at"`
	EndAt            time.Time `db:"end_at"`
	DurationMinutes  int       `db:"duration_minutes"`
	Capacity         int       `db:"capacity"`
	CreatedAt        time.Time `db:"created_at"`
	UpdatedAt        time.Time `db:"updated_at"`
}

type ActivityCreation struct {
	Name            string    `db:"name"`
	CreatorID       uuid.UUID `db:"creator_id"`
	Description     string    `db:"description"`
	StartAt         time.Time `db:"start_at"`
	EndAt           time.Time `db:"end_at"`
	DurationMinutes int       `db:"duration_minutes"`
	Capacity        int       `db:"capacity"`
	OrganizationID  uuid.UUID `db:"organization_id"`
}

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetActivityByID(id uuid.UUID) (*ActivityWithCreator, error) {
	const query = `
	SELECT
		a.id,
		a.name,
		a.creator_id,
		u.first_name AS creator_first_name,
		u.last_name AS creator_last_name,
		a.description,
		a.start_at,
		a.end_at,
		a.duration_minutes,
		a.capacity,
		a.created_at,
		a.updated_at
	FROM activities a
	JOIN users u on a.creator_id = u.id
	WHERE a.id = $1
	`
	act := ActivityWithCreator{}
	if err := r.db.GetContext(context.Background(), &act, query, id); err != nil {
		return nil, err
	}
	return &act, nil
}

func (r *Repository) CreateActivity(act *ActivityCreation) (*Activity, error) {
	const query = `
	INSERT INTO activities
		(name, creator_id, description, start_at, end_at, duration_minutes, capacity, organization_id, created_at, updated_at)
	VALUES
		($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	RETURNING *
	`
	created := Activity{}
	if err := r.db.GetContext(context.Background(), &created, query,
		act.Name,
		act.CreatorID,
		act.Description,
		act.StartAt,
		act.EndAt,
		act.DurationMinutes,
		act.Capacity,
		act.OrganizationID,
		time.Now(),
		time.Now(),
	); err != nil {
		return nil, err
	}

	return &created, nil
}

func (r *Repository) UpdateActivity(act *Activity) (*Activity, error) {
	const query = `
	UPDATE activities
	SET
		name = $1,
		description = $2,
		start_at = $3,
		end_at = $4,
		duration_minutes = $5,
		capacity = $6,
		updated_at = $7
	WHERE id = $8
	RETURNING *
	`
	act.UpdatedAt = time.Now()
	if err := r.db.GetContext(context.Background(), act, query,
		act.Name,
		act.Description,
		act.StartAt,
		act.EndAt,
		act.DurationMinutes,
		act.Capacity,
		act.UpdatedAt,
		act.ID,
	); err != nil {
		return nil, err
	}
	return act, nil
}

func (r *Repository) DeleteActivity(id uuid.UUID) error {
	const query = `DELETE FROM activities WHERE id = $1`
	_, err := r.db.ExecContext(context.Background(), query, id)
	return err
}

func (r *Repository) ListActivitiesByOrganization(organizationId uuid.UUID) ([]ActivityWithCreator, error) {
	const query = `
	SELECT
		a.id,
		a.name,
		a.creator_id,
		u.first_name AS creator_first_name,
		u.last_name AS creator_last_name,
		a.description,
		a.start_at,
		a.end_at,
		a.duration_minutes,
		a.capacity,
		a.created_at,
		a.updated_at
	FROM activities a
	JOIN users u on a.creator_id = u.id
	WHERE a.organization_id = $1
	ORDER BY a.created_at DESC
	`
	activities := make([]ActivityWithCreator, 0)
	if err := r.db.SelectContext(context.Background(), &activities, query, organizationId); err != nil {
		return nil, err
	}
	return activities, nil
}

func (r *Repository) ListActivitiesByOrganizations(organizationIds []uuid.UUID) ([]ActivityWithCreator, error) {
	const query = `
	SELECT
		a.id,
		a.name,
		a.creator_id,
		u.first_name AS creator_first_name,
		u.last_name AS creator_last_name,
		a.description,
		a.start_at,
		a.end_at,
		a.duration_minutes,
		a.capacity,
		a.created_at,
		a.updated_at
	FROM activities a
	JOIN users u on a.creator_id = u.id
	WHERE a.organization_id = ANY($1)
	ORDER BY a.created_at DESC
	`
	activities := make([]ActivityWithCreator, 0)
	if err := r.db.SelectContext(context.Background(), &activities, query, organizationIds); err != nil {
		return nil, err
	}
	return activities, nil
}
