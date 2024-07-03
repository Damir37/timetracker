package works

import (
	"github.com/gofiber/fiber/v2"
	"timertracker/internel/repository"
	"timertracker/internel/service"
	"timertracker/internel/types"
)

type DeleteWorkLogic struct {
	ServiceContext service.ServiceContext
	workRepository repository.WorkRep
}

func NewDeleteWorkLogic(serviceContext service.ServiceContext, workRepository repository.WorkRep) *DeleteWorkLogic {
	return &DeleteWorkLogic{
		ServiceContext: serviceContext,
		workRepository: workRepository,
	}
}

func (logic *DeleteWorkLogic) DeleteWork(workID int) (result *types.Result, errMsg *types.Errors, status int) {
	err := logic.workRepository.Delete(workID)

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

	return &types.Result{
		Status:  fiber.StatusNoContent,
		Message: "Работа удалена",
	}, nil, fiber.StatusNoContent
}
