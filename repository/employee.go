package repository

import (
	"github.com/google/uuid"
	"github.com/iBoBoTi/standup-management-tool/internal/models"
	"gorm.io/gorm"
)

type EmployeeRepository interface {
	CreateEmployee(user *models.Employee) error
	FindEmployeeByID(id uuid.UUID) (*models.Employee, error)
	FindEmployeeByEmail(email string) (*models.Employee, error)
	EmailExist(email string) (bool, error)
}

type employeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) *employeeRepository {
	return &employeeRepository{
		db: db,
	}
}

func (u *employeeRepository) CreateEmployee(employee *models.Employee) error {
	return u.db.Create(employee).Error
}

func (u *employeeRepository) FindEmployeeByID(id uuid.UUID) (*models.Employee, error) {
	employee := &models.Employee{}
	if err := u.db.Model(&models.Employee{}).Where("id", id).First(employee).Error; err != nil {
		return nil, err
	}
	return employee, nil
}

func (u *employeeRepository) FindEmployeeByEmail(email string) (*models.Employee, error) {
	employee := &models.Employee{}
	if err := u.db.Model(&models.Employee{}).Where("email", email).First(employee).Error; err != nil {
		return nil, err
	}
	return employee, nil
}

func (u *employeeRepository) EmailExist(email string) (bool, error) {
	var num int
	tx := u.db.Raw("SELECT CASE WHEN EXISTS (SELECT * FROM employees WHERE email = ?) THEN CAST(1 AS BIT)ELSE CAST(0 AS BIT) END", email).Scan(&num)
	if num == 1 {
		return true, tx.Error
	}
	return false, tx.Error
}
