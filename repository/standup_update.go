package repository

import (
	"github.com/google/uuid"
	"github.com/iBoBoTi/standup-management-tool/internal/models"
	"gorm.io/gorm"
)

type StandupUpdateRepository interface {
	CreateStandupUpdate(standupUpdate *models.StandupUpdate) error
	GetAllStandupUpdate(limit, page int) ([]models.StandupUpdate, error)
	DoesTodayStandupForEmployeeExist(employeeID uuid.UUID) (bool, error)
}

type standupUpdateRepository struct {
	db *gorm.DB
}

func NewStandupUpdateRepository(db *gorm.DB) *standupUpdateRepository {
	return &standupUpdateRepository{
		db: db,
	}
}

func (sr *standupUpdateRepository) CreateStandupUpdate(standupUpdate *models.StandupUpdate) error {
	return sr.db.Create(standupUpdate).Error
}

func (sr *standupUpdateRepository) GetAllStandupUpdate(limit, page int) ([]models.StandupUpdate, error) {
	var standupUpdate []models.StandupUpdate
	if err := sr.db.Model(&models.Sprint{}).Order("created_at ASC").Scopes(models.NewPaginate(limit, page).PaginatedResult).Find(&standupUpdate).Error; err != nil {
		return nil, err
	}
	return standupUpdate, nil
}

func (sr *standupUpdateRepository) DoesTodayStandupForEmployeeExist(employeeID uuid.UUID) (bool, error) {
	var num int
	tx := sr.db.Raw("SELECT CASE WHEN EXISTS (SELECT * FROM standup_updates WHERE employee_id = ? AND DATE(check_in_time) = DATE(now())) THEN CAST(1 AS BIT)ELSE CAST(0 AS BIT) END", employeeID).Scan(&num)
	if num == 1 {
		return true, tx.Error
	}
	return false, tx.Error
}
