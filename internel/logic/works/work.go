package works

import (
	"github.com/gofiber/fiber/v2"
	"timertracker/internel/repository"
	"timertracker/internel/service"
	"timertracker/internel/types"
)

type GetWorkLogic struct {
	ServiceContext service.ServiceContext
	workRepository repository.WorkRep
}

func NewGetWorkLogic(serviceContext service.ServiceContext, workRepository repository.WorkRep) *GetWorkLogic {
	return &GetWorkLogic{
		ServiceContext: serviceContext,
		workRepository: workRepository,
	}
}

func (logic *GetWorkLogic) GetWork(workID int) (result *types.Work, errMsg *types.Errors, status int) {
	work, err := logic.workRepository.FindByID(workID)
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
			Message: "Проблема с базой данных",
			Error:   err.Error(),
		}, fiber.StatusInternalServerError
	}

	data := &types.Work{
		Id:      work.ID,
		TaskID:  work.TaskID,
		UserID:  work.UserID,
		Hours:   work.Hours,
		Minutes: work.Minutes,
		Date:    work.Date,
	}

	return data, nil, fiber.StatusOK
}

func (logic *GetWorkLogic) GetWorks() (result []*types.Work, errMsg *types.Errors, status int) {
	works, err := logic.workRepository.FindAll()
	if err != nil {
		return nil, &types.Errors{
			Status:  fiber.StatusInternalServerError,
			Message: "Проблема с базой данных",
			Error:   err.Error(),
		}, fiber.StatusInternalServerError
	}

	var data []*types.Work
	for _, work := range works {
		data = append(data, &types.Work{
			Id:      work.ID,
			TaskID:  work.TaskID,
			UserID:  work.UserID,
			Hours:   work.Hours,
			Minutes: work.Minutes,
			Date:    work.Date,
		})
	}

	return data, nil, fiber.StatusOK
}
