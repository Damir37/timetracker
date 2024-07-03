package tasks

import (
	"github.com/gofiber/fiber/v2"
	"timertracker/internel/repository"
	"timertracker/internel/service"
	"timertracker/internel/types"
)

type DeleteTaskLogic struct {
	ServiceContext service.ServiceContext
	taskRepository repository.TaskRep
}

func NewDeleteTaskLogic(serviceContext service.ServiceContext, taskRepository repository.TaskRep) *DeleteTaskLogic {
	return &DeleteTaskLogic{
		ServiceContext: serviceContext,
		taskRepository: taskRepository,
	}
}

func (logic *DeleteTaskLogic) DeleteTask(taskID int) (result *types.Result, errMsg *types.Errors, status int) {
	err := logic.taskRepository.Delete(taskID)

	if err != nil {
		if err == repository.ErrorMsg {
			return nil, &types.Errors{
				Status:  fiber.StatusNotFound,
				Message: "Задача не найдена",
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
		Message: "Задача удалена",
	}, nil, fiber.StatusNoContent
}
