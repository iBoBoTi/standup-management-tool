package service

import (
	"fmt"
	"log"
	"time"

	"github.com/iBoBoTi/standup-management-tool/internal/dtos"
	"github.com/iBoBoTi/standup-management-tool/internal/errors"
	"github.com/iBoBoTi/standup-management-tool/internal/mappers"
	"github.com/iBoBoTi/standup-management-tool/internal/models"
	"github.com/iBoBoTi/standup-management-tool/internal/validator"
	"github.com/iBoBoTi/standup-management-tool/repository"
)

const (
	BeforeStandup = "before standup"
	AfterStandup  = "after standup"
	WithinStandup = "within standup"

	StandUpDuration = 15
)

type StandupUpdateService interface {
	CreateStandupUpdate(sprintDto *dtos.StandupUpdate) (*dtos.StandupUpdate, error)
	GetAllStandupUpdate(paginateReq *dtos.PaginatedRequest) ([]dtos.StandupUpdate, error)
}

type standupUpdateService struct {
	standupUpdateRepository repository.StandupUpdateRepository
	sprintRepository        repository.SprintRepository
}

func NewStandupUpdateService(standupUpdateRepository repository.StandupUpdateRepository, sprintRepository repository.SprintRepository) *standupUpdateService {
	return &standupUpdateService{
		standupUpdateRepository: standupUpdateRepository,
		sprintRepository:        sprintRepository,
	}
}

func (ss *standupUpdateService) CreateStandupUpdate(standupUpdateDto *dtos.StandupUpdate) (*dtos.StandupUpdate, error) {

	v := validator.NewValidator()

	presentDayStandupUpdateExist, err := ss.standupUpdateRepository.DoesTodayStandupForEmployeeExist(standupUpdateDto.EmployeeID)
	if err != nil {
		return nil, err
	}
	v.Check(!presentDayStandupUpdateExist, "employee", "standup update already exist")

	sprintExist, err := ss.sprintRepository.SprintExist(standupUpdateDto.SprintID)
	if err != nil {
		return nil, err
	}
	v.Check(sprintExist, "sprint_id", "sprint id doesn't exist")

	var sprint *models.Sprint

	if sprintExist {
		s, err := ss.sprintRepository.GetSprintByID(standupUpdateDto.SprintID)
		if err != nil {
			return nil, err
		}
		sprint = s

		v.Check(!hasSprintStarted(sprint.StartDateTime), "sprint_id", "sprint has not started")
		v.Check(!hasSprintEnded(sprint.EndDateTime), "sprint_id", "sprint has ended")

		standupUpdateDto.SprintName = sprint.Name
	}

	if !standupUpdateDto.Validate(v) {
		return nil, validator.NewValidationError("validation failed", v.Errors)
	}

	//checktime, compare with sprint daily update start time and set standup status
	standupUpdateDto.Status, err = getStandupUpdateStatus(standupUpdateDto.CheckInTime, sprint.DailyUpdateStartTime)
	if err != nil {
		log.Println(fmt.Errorf("error getting status for standup update: %v", err))
		return nil, errors.ErrInternalServer
	}

	standupUpdateModel := mappers.StandupUpdateDtoMapToStandupUpdateModel(standupUpdateDto)

	if err := ss.standupUpdateRepository.CreateStandupUpdate(standupUpdateModel); err != nil {
		return nil, err
	}

	return mappers.StandupUpdateModelMapToStandupUpdateDto(standupUpdateModel), nil
}

func (ss *standupUpdateService) GetAllStandupUpdate(paginateReq *dtos.PaginatedRequest) ([]dtos.StandupUpdate, error) {
	// paginateReq.Normalize()

	// sprintsModel, err := ss.standupUpdateRepository.GetAllStandupUpdate(paginateReq.Limit, paginateReq.Page)
	// if err != nil {
	// 	return nil, err
	// }

	// sprintsDto := make([]dtos.Sprint, 0)
	// if len(sprintsModel) > 0 {
	// 	for _, sprint := range sprintsModel {
	// 		sprintDto := mappers.SprintModelMapToSprintDto(&sprint)
	// 		sprintsDto = append(sprintsDto, *sprintDto)
	// 	}
	// }

	// return sprintsDto, nil
	return nil, nil
}

func hasSprintStarted(sprintStartDate time.Time) bool {

	today := time.Now().UTC()

	t := time.Date(sprintStartDate.Year(), sprintStartDate.Month(), sprintStartDate.Day(), 0, 0, 0, 0, time.UTC)
	sprintStartYear, sprintStartMonth, sprintStartDay := t.Date()

	todayYear, todayMonth, todayDay := today.Date()

	if todayYear > sprintStartYear ||
		(todayYear == sprintStartYear && todayMonth > sprintStartMonth) ||
		(todayYear == sprintStartYear && todayMonth == sprintStartMonth && todayDay > sprintStartDay) {
		return false
	}
	return true
}

func hasSprintEnded(sprintEndDate time.Time) bool {
	today := time.Now().UTC()

	t := time.Date(sprintEndDate.Year(), sprintEndDate.Month(), sprintEndDate.Day(), 0, 0, 0, 0, time.UTC)
	sprintEndYear, sprintEndMonth, sprintEndDay := t.Date()

	todayYear, todayMonth, todayDay := today.Date()

	if todayYear < sprintEndYear ||
		(todayYear == sprintEndYear && todayMonth < sprintEndMonth) ||
		(todayYear == sprintEndYear && todayMonth == sprintEndMonth && todayDay < sprintEndDay) {
		return false
	}
	return true
}

func getStandupUpdateStatus(checkInTime time.Time, timeForStandUp string) (string, error) {

	// Parse the time string into a time.Time object
	standupStartTime, err := time.Parse(time.Kitchen, timeForStandUp)
	if err != nil {
		return "", fmt.Errorf("error parsing time: %v", err)
	}

	standupEndTime := standupStartTime.Add(time.Duration(StandUpDuration) * time.Minute)

	switch {
	case checkInTime.Hour() < standupStartTime.Hour() || (checkInTime.Hour() == standupStartTime.Hour() && checkInTime.Minute() < standupStartTime.Minute()):
		return BeforeStandup, nil
	case checkInTime.Hour() == standupStartTime.Hour() && checkInTime.Hour() == standupStartTime.Minute() && checkInTime.Minute() <= standupEndTime.Minute():
		return WithinStandup, nil
	default:
		return AfterStandup, nil

	}

}
