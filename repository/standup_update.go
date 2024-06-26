package repository

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/iBoBoTi/standup-management-tool/internal/models"
	"gorm.io/gorm"
)

type StandupUpdateRepository interface {
	CreateStandupUpdate(standupUpdate *models.StandupUpdate) error
	GetAllStandupUpdate(limit, page int, query models.StandupUpdateQuery) ([]models.StandupUpdate, error)
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

func (sr *standupUpdateRepository) GetAllStandupUpdate(limit, page int, query models.StandupUpdateQuery) ([]models.StandupUpdate, error) {

	// Construct the query
	queryBuilder := sr.db.Model(&models.StandupUpdate{})
	if query.Owner != "" {
		queryBuilder = queryBuilder.Where("employee_id LIKE ?", fmt.Sprintf("%%%s%%", query.Owner))
	}
	if !query.Day.IsZero() {
		queryBuilder = queryBuilder.Where("check_in_time::date = ?", query.Day)
	}
	if !query.WeekStart.IsZero() && !query.WeekEnd.IsZero() {
		queryBuilder = queryBuilder.Where("check_in_time::date BETWEEN ? AND ?", query.WeekStart, query.WeekEnd)
	}
	if query.Sprint != "" {
		queryBuilder = queryBuilder.Where("sprint_id LIKE ?", fmt.Sprintf("%%%s%%", query.Sprint))
	}

	var standupUpdate []models.StandupUpdate
	if err := queryBuilder.Order("created_at ASC").Scopes(models.NewPaginate(limit, page).PaginatedResult).Find(&standupUpdate).Error; err != nil {
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
