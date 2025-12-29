package organizationapi

type CreateActivityRequest struct {
	Name            string `json:"name"`
	Description     string `json:"description"`
	StartAt         string `json:"start_at"`
	EndAt           string `json:"end_at"`
	DurationMinutes int    `json:"duration_minutes"`
	Capacity        int    `json:"capacity"`
}
