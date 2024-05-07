package mappers

import (
	"github.com/iBoBoTi/standup-management-tool/internal/dtos"
	"github.com/iBoBoTi/standup-management-tool/internal/models"
)

func SprintDtoMapToSprintModel(sprintDto *dtos.Sprint) *models.Sprint {
	return &models.Sprint{
		Name:                 sprintDto.Name,
		ProjectName:          sprintDto.ProjectName,
		StartDateTime:        sprintDto.StartDateTime,
		EndDateTime:          sprintDto.EndDateTime,
		Duration:             sprintDto.Duration,
		CreatorID:            sprintDto.CreatorID,
		DailyUpdateStartTime: sprintDto.DailyUpdateStartTime,
	}
}

func SprintModelMapToSprintDto(sprintModel *models.Sprint) *dtos.Sprint {
	return &dtos.Sprint{
		ID:                   sprintModel.ID,
		Name:                 sprintModel.Name,
		ProjectName:          sprintModel.ProjectName,
		StartDateTime:        sprintModel.StartDateTime,
		EndDateTime:          sprintModel.EndDateTime,
		Duration:             sprintModel.Duration,
		CreatorID:            sprintModel.CreatorID,
		DailyUpdateStartTime: sprintModel.DailyUpdateStartTime,
		CreatedAt:            sprintModel.CreatedAt,
		UpdatedAt:            sprintModel.UpdatedAt,
	}
}
