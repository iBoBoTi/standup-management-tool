package mappers

import (
	"github.com/iBoBoTi/standup-management-tool/internal/dtos"
	"github.com/iBoBoTi/standup-management-tool/internal/models"
)

func EmployeeDtoMapToEmployeeModel(employeeDto *dtos.Employee) *models.Employee {
	return &models.Employee{
		FirstName:    employeeDto.FirstName,
		LastName:     employeeDto.LastName,
		Email:        employeeDto.Email,
		PasswordHash: employeeDto.Password,
		Company:      employeeDto.Company,
		Role:         employeeDto.Role,
	}
}

func EmployeeModelMapToEmployeeDto(employee *models.Employee) *dtos.Employee {
	return &dtos.Employee{
		ID:        employee.ID,
		FirstName: employee.FirstName,
		LastName:  employee.LastName,
		Email:     employee.Email,
		Role:      employee.Role,
		Company:   employee.Company,
		CreatedAt: employee.CreatedAt,
		UpdatedAt: employee.UpdatedAt,
	}
}

func AdminCreateEmployeeDtoMapToEmployeeModel(employeeDto *dtos.AdminCreateEmployeeDto) *models.Employee {
	return &models.Employee{
		FirstName:    employeeDto.FirstName,
		LastName:     employeeDto.LastName,
		Email:        employeeDto.Email,
		Company:      employeeDto.Company,
		PasswordHash: employeeDto.PasswordHash,
		Role:         employeeDto.Role,
	}
}

func EmployeeModelMapToAdminCreateEmployeeDto(employee *models.Employee) *dtos.AdminCreateEmployeeDto {
	return &dtos.AdminCreateEmployeeDto{
		ID:        employee.ID,
		FirstName: employee.FirstName,
		LastName:  employee.LastName,
		Email:     employee.Email,
		Role:      employee.Role,
		Company:   employee.Company,
		CreatedAt: employee.CreatedAt,
		UpdatedAt: employee.UpdatedAt,
	}
}
