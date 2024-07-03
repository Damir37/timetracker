package works

import (
	"github.com/gofiber/fiber/v2"
	"timertracker/internel/repository"
	"timertracker/internel/service"
	"timertracker/internel/types"
)

type EditWorkLogic struct {
	ServiceContext service.ServiceContext
	workRepository repository.WorkRep
}

func NewEditWorkLogic(serviceContext service.ServiceContext, workRepository repository.WorkRep) *EditWorkLogic {
	return &EditWorkLogic{
		ServiceContext: serviceContext,
		workRepository: workRepository,
	}
}

func (logic *EditWorkLogic) EditWork(work *types.Work) (result *types.Work, errMsg *types.Errors, status int) {
	editWork := repository.Work{
		ID:      work.Id,
		TaskID:  work.TaskID,
		UserID:  work.UserID,
		Hours:   work.Hours,
		Minutes: work.Minutes,
		Date:    work.Date,
	}

	resultDb, err := logic.workRepository.Update(editWork)

	if err != nil {
		if err == repository.ErrorMsg {
			return nil, &types.Errors{
				Status:  fiber.StatusNotFound,
				Message: "Работа не найдена",
				Error:   err.Error(),
			}, fiber.StatusNotFound
		}

		return nil, &types.Errors{
			Status:  fiber.StatusInternalServerError,
			Message: "База данных не работает",
			Error:   err.Error(),
		}, fiber.StatusInternalServerError
	}

	result = &types.Work{
		Id:      resultDb.ID,
		TaskID:  resultDb.TaskID,
		UserID:  resultDb.UserID,
		Hours:   resultDb.Hours,
		Minutes: resultDb.Minutes,
		Date:    resultDb.Date,
	}

	return result, nil, fiber.StatusOK
}
