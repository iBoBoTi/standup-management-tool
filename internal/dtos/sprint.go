package dtos

import (
	"time"

	"github.com/google/uuid"
	"github.com/iBoBoTi/standup-management-tool/internal/validator"
)

type Sprint struct {
	ID                   uuid.UUID `json:"id"`
	Name                 string    `json:"name" binding:"required"`
	ProjectName          string    `json:"project_name" binding:"required"`
	StartDateTime        time.Time `json:"start_date_time" binding:"required"`
	EndDateTime          time.Time `json:"end_date_time"`
	Duration             int64     `json:"duration" binding:"required"` // in weeks
	CreatorID            uuid.UUID `json:"creator_id" binding:"required"`
	DailyUpdateStartTime string    `json:"daily_update_start_time" binding:"required"` //time only
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
}

func (s *Sprint) Validate(v *validator.Validator) bool {
	v.Check(s.Name != "", "name", "must not be blank")

	v.Check(s.ProjectName != "", "project_name", "must not be blank")

	v.Check(!s.StartDateTime.Before(time.Now()), "start_date_time", "must be later time")

	v.Check(s.Duration > 0, "duration", "should be greater than 0")
	v.Check(s.DailyUpdateStartTime != "", "daily_update_start_time", "must not be blank")

	return v.Valid()
}
