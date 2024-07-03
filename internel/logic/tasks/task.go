package tasks

import (
	"github.com/gofiber/fiber/v2"
	"timertracker/internel/repository"
	"timertracker/internel/service"
	"timertracker/internel/types"
	"timertracker/internel/utils"
)

type GetTaskLogic struct {
	ServiceContext service.ServiceContext
	taskRepository repository.TaskRep
}

func NewGetTaskLogic(serviceContext service.ServiceContext, taskRepository repository.TaskRep) *GetTaskLogic {
	return &GetTaskLogic{
		ServiceContext: serviceContext,
		taskRepository: taskRepository,
	}
}

func (logic *GetTaskLogic) GetTasks() (result []*types.Task, errMsg *types.Errors, status int) {
	all, err := logic.taskRepository.FindAll()
	if err != nil {
		return nil, &types.Errors{
			Status:  fiber.StatusInternalServerError,
			Message: "Проблема с базой данных",
			Error:   err.Error(),
		}, fiber.StatusInternalServerError
	}

	for _, task := range all {
		data := &types.Task{
			Id:          task.ID,
			UserID:      task.UserID,
			Description: task.Description,
			StartTime:   task.StartTime,
			EndTime:     task.EndTime,
			Duration:    task.Duration,
			Works:       utils.ConvertWorks(task.Works),
		}
		result = append(result, data)
	}

	return result, nil, fiber.StatusOK
}

func (logic *GetTaskLogic) GetTask(taskID int) (result *types.Task, errMsg *types.Errors, status int) {
	task, err := logic.taskRepository.FindByID(taskID)
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

	data := &types.Task{
		Id:          task.ID,
		UserID:      task.UserID,
		Description: task.Description,
		StartTime:   task.StartTime,
		EndTime:     task.EndTime,
		Duration:    task.Duration,
		Works:       utils.ConvertWorks(task.Works),
	}

	return data, nil, fiber.StatusOK
}
