package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iBoBoTi/standup-management-tool/internal/dtos"
	"github.com/iBoBoTi/standup-management-tool/internal/validator"
	"github.com/iBoBoTi/standup-management-tool/server"
	"github.com/iBoBoTi/standup-management-tool/service"
)

type EmployeeHandler interface {
	AdminCreateEmployee(ctx *gin.Context)
}

type employeeHandler struct {
	srv             *server.Server
	employeeService service.EmployeeService
}

func NewEmployeeHandler(srv *server.Server, employeeService service.EmployeeService) *employeeHandler {
	return &employeeHandler{
		srv:             srv,
		employeeService: employeeService,
	}
}

func (eh *employeeHandler) AdminCreateEmployee(ctx *gin.Context) {

	admin := eh.srv.ContextGetUser(ctx)

	var employeeRequest dtos.AdminCreateEmployeeDto

	if err := ctx.ShouldBindJSON(&employeeRequest); err != nil {
		server.ErrorJSONResponse(ctx, http.StatusBadRequest, err)
		return
	}

	employeeRequest.Company = admin.Company
	employeeRequest.Role = dtos.EmployeeRole

	createdEmployee, err := eh.employeeService.AdminCreateEmployee(&employeeRequest)
	if err != nil {
		var e *validator.ValidationError
		switch {
		case errors.As(err, &e):
			server.SendValidationError(ctx, e)
		default:
			server.ErrorJSONResponse(ctx, http.StatusInternalServerError, err)
		}
		return
	}

	server.SuccessJSONResponse(ctx, http.StatusCreated, "employee created successfully", createdEmployee)
}
