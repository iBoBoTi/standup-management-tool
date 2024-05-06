package mappers

import (
	"github.com/iBoBoTi/standup-management-tool/internal/dtos"
	"github.com/iBoBoTi/standup-management-tool/internal/models"
)

func EmployeeDtoMapToEmployeeModel(userDto *dtos.Employee) *models.Employee {
	return &models.Employee{
		FirstName:    userDto.FirstName,
		LastName:     userDto.LastName,
		Email:        userDto.Email,
		PasswordHash: userDto.Password,
		Role:         userDto.Role,
	}
}

func EmployeeModelMapToEmployeeDto(user *models.Employee) *dtos.Employee {
	return &dtos.Employee{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Role:      user.Role,
		Company:   user.Company,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
