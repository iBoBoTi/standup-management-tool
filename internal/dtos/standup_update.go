package dtos

import (
	"time"

	"github.com/google/uuid"
	"github.com/iBoBoTi/standup-management-tool/internal/validator"
)

type StandupUpdate struct {
	ID                    uuid.UUID `json:"id"`
	TaskID                string    `json:"task_id" binding:"required"`
	EmployeeID            uuid.UUID `json:"employee_id"`
	EmployeeName          string    `json:"employee_name"`
	SprintID              uuid.UUID `json:"sprint_id"`
	SprintName            string    `json:"sprint_name"`
	NextUpdateToDo        string    `json:"next_update_todo" binding:"required"`
	PreviousUpdateDone    string    `json:"previous_update_done" binding:"required"`
	BlockedByID           uuid.UUID `json:"blocked_by_id"`
	BreakAway             string    `json:"break_away"`
	CheckInTime           time.Time `json:"check_in_time"`
	Status                string    `json:"status"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
}

func (s *StandupUpdate) Validate(v *validator.Validator) bool {
	v.Check(s.TaskID != "", "task", "must not be blank")

	v.Check(s.NextUpdateToDo != "", "next_update_todo", "must not be blank")

	v.Check(s.PreviousUpdateDone != "", "previous_update_done", "must be later time")

	return v.Valid()
}
