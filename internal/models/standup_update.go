package models

import (
	"time"

	"github.com/google/uuid"
)

type StandupUpdate struct {
	Model
	TaskID             string    `gorm:"not null"`
	EmployeeID         uuid.UUID `gorm:"not null"`
	EmployeeName       string    `gorm:"not null"`
	SprintID           uuid.UUID `gorm:"not null"`
	SprintName         string    `gorm:"not null"`
	NextUpdateToDo     string    `gorm:"not null"`
	PreviousUpdateDone string    `gorm:"not null"`
	BlockedByID        uuid.UUID
	BreakAway          string    `gorm:"not null"`
	CheckInTime        time.Time `gorm:"not null"`
	Status             string    `gorm:"not null"`
}

type StandupUpdateQuery struct {
	Day       time.Time
	Sprint    string
	Owner     string
	WeekStart time.Time
	WeekEnd   time.Time
}
