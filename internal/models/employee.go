package models

type Employee struct {
	Model
	FirstName      string `gorm:"not null"`
	LastName       string `gorm:"not null"`
	Email          string `gorm:"unique;not null"`
	PasswordHash   string `gorm:"not null"`
	Role           string `gorm:"not null"`
	Company        string
	StandupUpdates []StandupUpdate `gorm:"foreignKey:EmployeeID;constraint:OnDelete:CASCADE"`
}
