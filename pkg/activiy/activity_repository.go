package activiy

import (
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type activity struct {
	ID              string    `db:"id"`
	Name            string    `db:"name"`
	CreatorId       uuid.UUID `db:"creator_id"`
	Description     string    `db:"description"`
	StartAt         time.Time `db:"start_at"`
	EndAt           time.Time `db:"end_at"`
	DurationMinutes int       `db:"duration_minutes"`
	Capacity        int       `db:"capacity"`
	createdAt       time.Time `db:"created_at"`
	updatedAt       time.Time `db:"updated_at"`
}

type activityWithCreator struct {
	ID               string    `db:"id"`
	Name             string    `db:"name"`
	CreatorId        uuid.UUID `db:"creator_id"`
	CreatorFirstName string    `db:"creator_first_name"`
	CreatorLastName  string    `db:"creator_last_name"`
	Description      string    `db:"description"`
	StartAt          time.Time `db:"start_at"`
	EndAt            time.Time `db:"end_at"`
	DurationMinutes  int       `db:"duration_minutes"`
	Capacity         int       `db:"capacity"`
	createdAt        time.Time `db:"created_at"`
	updatedAt        time.Time `db:"updated_at"`
}

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetActivityByID(id string) (*activityWithCreator, error) {
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
	act := activityWithCreator{}
	if err := r.db.Get(&act, query, id); err != nil {
		return nil, err
	}
	return &act, nil
}

func (r *Repository) CreateActivity(act *activity) (*activity, error) {
	const query = `
	INSERT INTO activities
		(id, name, creator_id, description, start_at, end_at, duration_minutes, capacity, created_at, updated_at)
	VALUES
		($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	RETURNING *
	`
	act.createdAt = time.Now()
	act.updatedAt = time.Now()
	if err := r.db.Get(act, query,
		act.ID,
		act.Name,
		act.CreatorId,
		act.Description,
		act.StartAt,
		act.EndAt,
		act.DurationMinutes,
		act.Capacity,
		act.createdAt,
		act.updatedAt,
	); err != nil {
		return nil, err
	}
	return act, nil
}

func (r *Repository) UpdateActivity(act *activity) (*activity, error) {
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
	act.updatedAt = time.Now()
	if err := r.db.Get(act, query,
		act.Name,
		act.Description,
		act.StartAt,
		act.EndAt,
		act.DurationMinutes,
		act.Capacity,
		act.updatedAt,
		act.ID,
	); err != nil {
		return nil, err
	}
	return act, nil
}

func (r *Repository) DeleteActivity(id string) error {
	const query = `DELETE FROM activities WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *Repository) ListActivitiesByOrganization(organizationId uuid.UUID) ([]activityWithCreator, error) {
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
	activities := make([]activityWithCreator, 0)
	if err := r.db.Select(&activities, query, organizationId); err != nil {
		return nil, err
	}
	return activities, nil
}

func (r *Repository) ListActivitiesByOrganizations(organizationIds []uuid.UUID) ([]activityWithCreator, error) {
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
	activities := make([]activityWithCreator, 0)
	if err := r.db.Select(&activities, query, organizationIds); err != nil {
		return nil, err
	}
	return activities, nil
}
