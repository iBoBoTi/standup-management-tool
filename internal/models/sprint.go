package models

import (
	"time"

	"github.com/google/uuid"
)

type Sprint struct {
	Model
	Name                 string          `gorm:"not null"`
	ProjectName          string          `gorm:"unique;not null"`
	StartDateTime        time.Time       `gorm:"not null"`
	EndDateTime          time.Time       `gorm:"not null"`
	Duration             int64           `gorm:"not null"` // in weeks
	CreatorID            uuid.UUID       `gorm:"not null"`
	DailyUpdateStartTime string          `gorm:"not null"`
	StandupUpdates       []StandupUpdate `gorm:"foreignKey:SprintID;constraint:OnDelete:CASCADE"`
}
