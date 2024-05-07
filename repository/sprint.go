package repository

import (
	"github.com/iBoBoTi/standup-management-tool/internal/models"
	"gorm.io/gorm"
)

type SprintRepository interface {
	CreateSprint(sprint *models.Sprint) error
	GetAllSprints(limit, page int) ([]models.Sprint, error)
	SprintProjectNameExist(name string) (bool, error)
}

type sprintRepository struct {
	db *gorm.DB
}

func NewSprintRepository(db *gorm.DB) *sprintRepository {
	return &sprintRepository{
		db: db,
	}
}

func (sr *sprintRepository) CreateSprint(sprint *models.Sprint) error {
	return sr.db.Create(sprint).Error
}

func (sr *sprintRepository) GetAllSprints(limit, page int) ([]models.Sprint, error) {
	var sprints []models.Sprint
	if err := sr.db.Model(&models.Sprint{}).Order("created_at ASC").Scopes(models.NewPaginate(limit, page).PaginatedResult).Find(&sprints).Error; err != nil {
		return nil, err
	}
	return sprints, nil
}

func (sr *sprintRepository) SprintProjectNameExist(name string) (bool, error) {
	var num int
	tx := sr.db.Raw("SELECT CASE WHEN EXISTS (SELECT * FROM sprints WHERE project_name = ?) THEN CAST(1 AS BIT)ELSE CAST(0 AS BIT) END", name).Scan(&num)
	if num == 1 {
		return true, tx.Error
	}
	return false, tx.Error
}
