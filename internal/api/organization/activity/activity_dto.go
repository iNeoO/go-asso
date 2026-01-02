package organizationapi

import "time"

type CreateActivityRequest struct {
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	StartAt         time.Time `json:"start_at"`
	EndAt           time.Time `json:"end_at"`
	DurationMinutes int       `json:"duration_minutes"`
	Capacity        int       `json:"capacity"`
}
