package tasks

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"timertracker/internel/repository"
	"timertracker/internel/service"
	"timertracker/internel/types"
	"timertracker/internel/utils"
)

type EditTaskLogic struct {
	ServiceContext service.ServiceContext
	taskRepository repository.TaskRep
}

func NewEditTaskLogic(serviceContext service.ServiceContext, taskRepository repository.TaskRep) *EditTaskLogic {
	return &EditTaskLogic{
		ServiceContext: serviceContext,
		taskRepository: taskRepository,
	}
}

func (logic *EditTaskLogic) EditTask(task *types.Task) (result *types.Task, errMsg *types.Errors, status int) {
	editTask := repository.Task{
		Model: gorm.Model{
			ID: task.Id,
		},
		UserID:      task.UserID,
		Description: task.Description,
		StartTime:   task.StartTime,
		EndTime:     task.EndTime,
		Duration:    task.Duration,
		Works:       []repository.Work{},
	}

	resultDb, err := logic.taskRepository.Update(editTask)

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
			Message: "База данных не работает",
			Error:   err.Error(),
		}, fiber.StatusInternalServerError
	}

	resp := &types.Task{
		Id:          resultDb.ID,
		UserID:      resultDb.UserID,
		Description: resultDb.Description,
		StartTime:   resultDb.StartTime,
		EndTime:     resultDb.EndTime,
		Duration:    resultDb.Duration,
		Works:       utils.ConvertWorks(resultDb.Works),
	}

	return resp, nil, fiber.StatusOK
}
