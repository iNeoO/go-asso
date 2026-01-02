package organizationapi

import (
	"time"

	"github.com/google/uuid"
	commonapi "github.com/ineoo/go-planigramme/internal/api/common"
	"github.com/ineoo/go-planigramme/pkg/activity"
)

type Activity struct {
	ID              uuid.UUID `json:"id"`
	Name            string    `json:"name"`
	CreatorID       uuid.UUID `json:"creator_id"`
	Description     string    `json:"description"`
	StartAt         time.Time `json:"start_at"`
	EndAt           time.Time `json:"end_at"`
	DurationMinutes int       `json:"duration_minutes"`
	Capacity        int       `json:"capacity"`
	OrganizationID  uuid.UUID `json:"organization_id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type ActivityWithCreator struct {
	ID               uuid.UUID `json:"id"`
	Name             string    `json:"name"`
	CreatorID        uuid.UUID `json:"creator_id"`
	CreatorFirstName string    `json:"creator_first_name"`
	CreatorLastName  string    `json:"creator_last_name"`
	Description      string    `json:"description"`
	StartAt          time.Time `json:"start_at"`
	EndAt            time.Time `json:"end_at"`
	DurationMinutes  int       `json:"duration_minutes"`
	Capacity         int       `json:"capacity"`
	OrganizationID   uuid.UUID `json:"organization_id"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type ErrorEnvelope = commonapi.ErrorEnvelope

type ActivityEnvelope struct {
	Status bool      `json:"status"`
	Data   *Activity `json:"activity"`
	Error  *string   `json:"error"`
}

type ActivitiesEnvelope struct {
	Status bool                  `json:"status"`
	Data   []ActivityWithCreator `json:"activities"`
	Count  int                   `json:"count"`
	Error  *string               `json:"error"`
}

func ActivitySuccessResponse(act *activity.Activity) ActivityEnvelope {
	apiActivity := Activity{
		ID:              act.ID,
		Name:            act.Name,
		CreatorID:       act.CreatorID,
		Description:     act.Description,
		StartAt:         act.StartAt,
		EndAt:           act.EndAt,
		DurationMinutes: act.DurationMinutes,
		Capacity:        act.Capacity,
		OrganizationID:  act.OrganizationID,
		CreatedAt:       act.CreatedAt,
		UpdatedAt:       act.UpdatedAt,
	}
	return ActivityEnvelope{
		Status: true,
		Data:   &apiActivity,
		Error:  nil,
	}
}

func ActivityErrorResponse(message string) ErrorEnvelope {
	return ErrorEnvelope{
		Status: false,
		Data:   nil,
		Error:  message,
	}
}

func ActivitiesListSuccessResponse(activities []activity.ActivityWithCreator) ActivitiesEnvelope {
	items := make([]ActivityWithCreator, 0, len(activities))
	for _, a := range activities {
		items = append(items, ActivityWithCreator{
			ID:               a.ID,
			Name:             a.Name,
			CreatorID:        a.CreatorID,
			CreatorFirstName: a.CreatorFirstName,
			CreatorLastName:  a.CreatorLastName,
			Description:      a.Description,
			StartAt:          a.StartAt,
			EndAt:            a.EndAt,
			DurationMinutes:  a.DurationMinutes,
			Capacity:         a.Capacity,
			OrganizationID:   a.OrganizationID,
			CreatedAt:        a.CreatedAt,
			UpdatedAt:        a.UpdatedAt,
		})
	}

	return ActivitiesEnvelope{
		Status: true,
		Data:   items,
		Count:  len(items),
		Error:  nil,
	}
}
