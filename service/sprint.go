package service

import (
	"github.com/iBoBoTi/standup-management-tool/internal/dtos"
	"github.com/iBoBoTi/standup-management-tool/internal/mappers"
	"github.com/iBoBoTi/standup-management-tool/internal/validator"
	"github.com/iBoBoTi/standup-management-tool/repository"
)

type SprintService interface {
	CreateSprint(sprintDto *dtos.Sprint) (*dtos.Sprint, error)
	GetAllSprints(paginateReq *dtos.PaginatedRequest) ([]dtos.Sprint, error)
}

type sprintService struct {
	sprintRepository repository.SprintRepository
}

func NewSprintRepository(sprintRepository repository.SprintRepository) *sprintService {
	return &sprintService{
		sprintRepository: sprintRepository,
	}
}

func (ss *sprintService) CreateSprint(sprintDto *dtos.Sprint) (*dtos.Sprint, error) {

	v := validator.NewValidator()

	emailExist, err := ss.sprintRepository.SprintProjectNameExist(sprintDto.ProjectName)
	if err != nil {
		return nil, err
	}
	v.Check(!emailExist, "project_name", "already exist")

	if !sprintDto.Validate(v) {
		return nil, validator.NewValidationError("validation failed", v.Errors)
	}

	sprintModel := mappers.SprintDtoMapToSprintModel(sprintDto)

	if err := ss.sprintRepository.CreateSprint(sprintModel); err != nil {
		return nil, err
	}

	return mappers.SprintModelMapToSprintDto(sprintModel), nil
}

func (ss *sprintService) GetAllSprints(paginateReq *dtos.PaginatedRequest) ([]dtos.Sprint, error) {
	paginateReq.Normalize()

	sprintsModel, err := ss.sprintRepository.GetAllSprints(paginateReq.Limit, paginateReq.Page)
	if err != nil {
		return nil, err
	}

	sprintsDto := make([]dtos.Sprint, 0)
	if len(sprintsModel) > 0 {
		for _, sprint := range sprintsModel {
			sprintDto := mappers.SprintModelMapToSprintDto(&sprint)
			sprintsDto = append(sprintsDto, *sprintDto)
		}
	}

	return sprintsDto, nil
}
