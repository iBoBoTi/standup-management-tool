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

type SprintHandler interface {
	CreateSprint(ctx *gin.Context)
	GetAllSprints(ctx *gin.Context)
}

type sprintHandler struct {
	srv           *server.Server
	sprintService service.SprintService
}

func NewSprintHandler(srv *server.Server, sprintService service.SprintService) *sprintHandler {
	return &sprintHandler{
		srv:           srv,
		sprintService: sprintService,
	}
}

func (sh *sprintHandler) CreateSprint(ctx *gin.Context) {

	admin := sh.srv.ContextGetUser(ctx)

	var sprintRequest dtos.Sprint

	if err := ctx.ShouldBindJSON(&sprintRequest); err != nil {
		server.ErrorJSONResponse(ctx, http.StatusBadRequest, err)
		return
	}

	sprintRequest.CreatorID = admin.ID

	createdSprint, err := sh.sprintService.CreateSprint(&sprintRequest)
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

	server.SuccessJSONResponse(ctx, http.StatusCreated, "sprint created successfully", createdSprint)
}

func (sh *sprintHandler) GetAllSprints(ctx *gin.Context) {
	var req dtos.PaginatedRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		server.ErrorJSONResponse(ctx, http.StatusBadRequest, err)
		return
	}

	sprints, err := sh.sprintService.GetAllSprints(&req)
	if err != nil {
		server.ErrorJSONResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	server.SuccessJSONResponse(ctx, http.StatusOK, "job posts retrieved successfully", sprints)
}
