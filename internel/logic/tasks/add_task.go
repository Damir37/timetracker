package tasks

import (
	"github.com/gofiber/fiber/v2"
	"timertracker/internel/repository"
	"timertracker/internel/service"
	"timertracker/internel/types"
	"timertracker/internel/utils"
)

type AddTaskLogic struct {
	ServiceContext service.ServiceContext
	taskRepository repository.TaskRep
}

func NewAddTaskLogic(serviceContext service.ServiceContext, taskRepository repository.TaskRep) *AddTaskLogic {
	return &AddTaskLogic{
		ServiceContext: serviceContext,
		taskRepository: taskRepository,
	}
}

func (logic *AddTaskLogic) AddTask(task *types.Task) (result *types.Task, errMsg *types.Errors, status int) {
	createTask := repository.Task{
		UserID:      task.UserID,
		Description: task.Description,
		StartTime:   task.StartTime,
		EndTime:     task.EndTime,
		Duration:    task.Duration,
		Works:       []repository.Work{},
	}

	resultDb, err := logic.taskRepository.Create(createTask)
	if err != nil {
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

	return resp, nil, fiber.StatusCreated
}
