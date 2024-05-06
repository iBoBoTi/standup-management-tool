package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iBoBoTi/standup-management-tool/internal/dtos"
	appError "github.com/iBoBoTi/standup-management-tool/internal/errors"
	"github.com/iBoBoTi/standup-management-tool/internal/validator"
	"github.com/iBoBoTi/standup-management-tool/server"
	"github.com/iBoBoTi/standup-management-tool/service"
)

type AuthHandler interface {
	SignUp(ctx *gin.Context)
	Login(ctx *gin.Context)
}

type authHandler struct {
	srv         *server.Server
	authService service.AuthService
}

func NewAuthHandler(srv *server.Server, authService service.AuthService) *authHandler {
	return &authHandler{
		srv:         srv,
		authService: authService,
	}
}

func (ah *authHandler) SignUp(ctx *gin.Context) {
	var employeeRequest dtos.Employee

	if err := ctx.ShouldBindJSON(&employeeRequest); err != nil {
		server.ErrorJSONResponse(ctx, http.StatusBadRequest, err)
		return
	}

	employeeRequest.Role = dtos.AdminRole

	createdEmployee, err := ah.authService.CreateEmployee(&employeeRequest)
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

	server.SuccessJSONResponse(ctx, http.StatusCreated, "employee signup successfully", createdEmployee)
}

func (ah *authHandler) Login(ctx *gin.Context) {
	var loginRequest dtos.LoginRequest

	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		server.ErrorJSONResponse(ctx, http.StatusBadRequest, err)
		return
	}

	res, err := ah.authService.Login(&loginRequest)
	if err != nil {
		var e *validator.ValidationError
		switch {
		case errors.As(err, &e):
			server.SendValidationError(ctx, e)
		default:
			ah.srv.Logger.Error(err, nil)
			server.ErrorJSONResponse(ctx, appError.ErrStatusCode(err), appError.ErrInternalServer)
		}
		return
	}

	server.SuccessJSONResponse(ctx, http.StatusOK, "login successful", res.Data)
}
