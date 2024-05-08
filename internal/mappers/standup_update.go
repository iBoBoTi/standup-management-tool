package mappers

import (
	"github.com/iBoBoTi/standup-management-tool/internal/dtos"
	"github.com/iBoBoTi/standup-management-tool/internal/models"
)

func StandupUpdateDtoMapToStandupUpdateModel(standupUpdateDto *dtos.StandupUpdate) *models.StandupUpdate {
	return &models.StandupUpdate{
		TaskID:             standupUpdateDto.TaskID,
		EmployeeID:         standupUpdateDto.EmployeeID,
		EmployeeName:       standupUpdateDto.EmployeeName,
		SprintID:           standupUpdateDto.SprintID,
		SprintName:         standupUpdateDto.SprintName,
		NextUpdateToDo:     standupUpdateDto.NextUpdateToDo,
		PreviousUpdateDone: standupUpdateDto.PreviousUpdateDone,
		BlockedByID:        standupUpdateDto.BlockedByID,
		BreakAway:          standupUpdateDto.BreakAway,
		CheckInTime:        standupUpdateDto.CheckInTime,
		Status:             standupUpdateDto.Status,
	}
}

func StandupUpdateModelMapToStandupUpdateDto(standupUpdateModel *models.StandupUpdate) *dtos.StandupUpdate {
	return &dtos.StandupUpdate{
		ID:                 standupUpdateModel.ID,
		TaskID:             standupUpdateModel.TaskID,
		EmployeeID:         standupUpdateModel.EmployeeID,
		EmployeeName:       standupUpdateModel.EmployeeName,
		SprintID:           standupUpdateModel.SprintID,
		SprintName:         standupUpdateModel.SprintName,
		NextUpdateToDo:     standupUpdateModel.NextUpdateToDo,
		PreviousUpdateDone: standupUpdateModel.PreviousUpdateDone,
		BlockedByID:        standupUpdateModel.BlockedByID,
		BreakAway:          standupUpdateModel.BreakAway,
		CheckInTime:        standupUpdateModel.CheckInTime,
		Status:             standupUpdateModel.Status,
		CreatedAt:          standupUpdateModel.CreatedAt,
		UpdatedAt:          standupUpdateModel.UpdatedAt,
	}
}
