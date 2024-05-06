package service

import (
	"fmt"
	"time"

	"github.com/iBoBoTi/standup-management-tool/internal/dtos"
	appError "github.com/iBoBoTi/standup-management-tool/internal/errors"
	"github.com/iBoBoTi/standup-management-tool/internal/mappers"
	"github.com/iBoBoTi/standup-management-tool/internal/security"
	"github.com/iBoBoTi/standup-management-tool/internal/validator"
	"github.com/iBoBoTi/standup-management-tool/repository"
)

var AccessTokenDuration = 30 * time.Minute

type AuthService interface {
	CreateEmployee(employeeDto *dtos.Employee) (*dtos.Employee, error)
	Login(loginRequest *dtos.LoginRequest) (*security.AuthPayload, error)
}

type authService struct {
	tokenMaker         security.Maker
	employeeRepository repository.EmployeeRepository
}

func NewAuthService(tokenMaker security.Maker, employeeRepository repository.EmployeeRepository) *authService {
	return &authService{
		tokenMaker:         tokenMaker,
		employeeRepository: employeeRepository,
	}
}

func (as *authService) CreateEmployee(employeeDto *dtos.Employee) (*dtos.Employee, error) {

	v := validator.NewValidator()

	emailExist, err := as.employeeRepository.EmailExist(employeeDto.Email)
	if err != nil {
		return nil, err
	}
	v.Check(!emailExist, "email", "email already exist")

	if !employeeDto.Validate(v) {
		return nil, validator.NewValidationError("validation failed", v.Errors)
	}

	if err := employeeDto.HashPassword(); err != nil {
		return nil, err
	}

	employee := mappers.EmployeeDtoMapToEmployeeModel(employeeDto)

	if err := as.employeeRepository.CreateEmployee(employee); err != nil {
		return nil, err
	}

	return mappers.EmployeeModelMapToEmployeeDto(employee), nil
}

func (as *authService) Login(loginRequest *dtos.LoginRequest) (*security.AuthPayload, error) {
	v := validator.NewValidator()
	if !loginRequest.Validate(v) {
		return nil, validator.NewValidationError("validation failed", v.Errors)
	}

	emailExist, err := as.employeeRepository.EmailExist(loginRequest.Email)
	if err != nil {
		return nil, err
	}

	if !emailExist {
		return nil, fmt.Errorf("%v: email not found", appError.ErrNotFound)
	}

	foundUser, err := as.employeeRepository.FindEmployeeByEmail(loginRequest.Email)
	if err != nil {
		return nil, appError.ErrInternalServer
	}

	if err := loginRequest.CheckPassword(foundUser.PasswordHash); err != nil {
		return nil, appError.ErrInvalidCredential
	}

	var res security.AuthPayload
	res.Data = make(map[string]any)

	if err := as.tokenMaker.GenerateAuthAccessToken(foundUser, &res, AccessTokenDuration); err != nil {
		return nil, err
	}

	return &res, nil
}
