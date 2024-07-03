package works

import (
	"github.com/gofiber/fiber/v2"
	"timertracker/internel/repository"
	"timertracker/internel/service"
	"timertracker/internel/types"
)

type AddWorkLogic struct {
	ServiceContext service.ServiceContext
	workRepository repository.WorkRep
}

func NewAddWorkLogic(serviceContext service.ServiceContext, workRepository repository.WorkRep) *AddWorkLogic {
	return &AddWorkLogic{
		ServiceContext: serviceContext,
		workRepository: workRepository,
	}
}

func (logic *AddWorkLogic) AddWork(work *types.Work) (result *types.Work, errMsg *types.Errors, status int) {
	createWork := repository.Work{
		TaskID:  work.TaskID,
		UserID:  work.UserID,
		Hours:   work.Hours,
		Minutes: work.Minutes,
		Date:    work.Date,
	}

	resultDb, err := logic.workRepository.Create(createWork)
	if err != nil {
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

	return result, nil, fiber.StatusCreated
}
