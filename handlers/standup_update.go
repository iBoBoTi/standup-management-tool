package handlers

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/iBoBoTi/standup-management-tool/internal/dtos"
	"github.com/iBoBoTi/standup-management-tool/internal/validator"
	"github.com/iBoBoTi/standup-management-tool/server"
	"github.com/iBoBoTi/standup-management-tool/service"
)

type StandupUpdateHandler interface {
	CreateStandupUpdate(ctx *gin.Context)
	GetAllStandupUpdate(ctx *gin.Context)
}

type standupUpdateHandler struct {
	srv                  *server.Server
	standupUpdateService service.StandupUpdateService
}

func NewStandupUpdateHandler(srv *server.Server, standupUpdateService service.StandupUpdateService) *standupUpdateHandler {
	return &standupUpdateHandler{
		srv:                  srv,
		standupUpdateService: standupUpdateService,
	}
}

func (sh *standupUpdateHandler) CreateStandupUpdate(ctx *gin.Context) {

	employee := sh.srv.ContextGetUser(ctx)

	var standupUpdateRequest dtos.StandupUpdate

	if err := ctx.ShouldBindJSON(&standupUpdateRequest); err != nil {
		server.ErrorJSONResponse(ctx, http.StatusBadRequest, err)
		return
	}

	standupUpdateRequest.EmployeeID = employee.ID
	standupUpdateRequest.EmployeeName = employee.FirstName + " " + employee.LastName

	sprintID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		server.ErrorJSONResponse(ctx, http.StatusUnprocessableEntity, errors.New("invalid sprint_id param"))
		return
	}

	standupUpdateRequest.SprintID = sprintID
	standupUpdateRequest.CheckInTime = time.Now().Local()

	createdStandupUpdate, err := sh.standupUpdateService.CreateStandupUpdate(&standupUpdateRequest)
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

	server.SuccessJSONResponse(ctx, http.StatusCreated, "standup update created successfully", createdStandupUpdate)
}
