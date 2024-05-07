package service

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/iBoBoTi/standup-management-tool/internal/dtos"
	"github.com/iBoBoTi/standup-management-tool/internal/mappers"
	"github.com/iBoBoTi/standup-management-tool/internal/validator"
	"github.com/iBoBoTi/standup-management-tool/repository"
)

const (
	passwordLength = 10
)

type EmployeeService interface {
	AdminCreateEmployee(employeeDto *dtos.AdminCreateEmployeeDto) (*dtos.AdminCreateEmployeeDto, error)
}

type employeeService struct {
	employeeRepository repository.EmployeeRepository
}

func NewEmployeeService(employeeRepository repository.EmployeeRepository) *employeeService {
	return &employeeService{
		employeeRepository: employeeRepository,
	}
}

func (es *employeeService) AdminCreateEmployee(employeeDto *dtos.AdminCreateEmployeeDto) (*dtos.AdminCreateEmployeeDto, error) {

	v := validator.NewValidator()

	emailExist, err := es.employeeRepository.EmailExist(employeeDto.Email)
	if err != nil {
		return nil, err
	}
	v.Check(!emailExist, "email", "email already exist")

	if !employeeDto.Validate(v) {
		return nil, validator.NewValidationError("validation failed", v.Errors)
	}

	employeeDto.Password = generateRandomPassword(passwordLength)

	if err := employeeDto.HashPassword(); err != nil {
		return nil, err
	}

	employee := mappers.AdminCreateEmployeeDtoMapToEmployeeModel(employeeDto)

	if err := es.employeeRepository.CreateEmployee(employee); err != nil {
		return nil, err
	}

	createdEmployeeRes := mappers.EmployeeModelMapToAdminCreateEmployeeDto(employee)
	createdEmployeeRes.Password = employeeDto.Password

	return createdEmployeeRes, nil
}

func generateRandomPassword(length int) string {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(b)
}
