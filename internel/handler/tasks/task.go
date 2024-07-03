package tasks

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
	"timertracker/internel/logic/tasks"
	"timertracker/internel/repository"
	"timertracker/internel/service"
	"timertracker/internel/types"
)

// GetTasksHandler godoc
// @Summary Получает все задачи из базы данных
// @Description Получает список всех задач в системе
// @Tags tasks
// @Accept json
// @Produce json
// @Success 200 {array} types.TaskSwagger
// @Failure 500 {object} types.Errors
// @Router /v1/tasks [get]
func GetTasksHandler(ctx *fiber.Ctx, serviceContext service.ServiceContext) error {
	taskRepository := repository.NewTask(serviceContext)
	taskGet := tasks.NewGetTaskLogic(serviceContext, *taskRepository)

	result, errMsg, status := taskGet.GetTasks()
	if errMsg != nil {
		return ctx.Status(status).JSON(errMsg)
	}

	return ctx.Status(status).JSON(result)
}

// GetTaskHandler godoc
// @Summary Получает задачу из базы данных по ID
// @Description Получает задачу из системы по её ID
// @Tags tasks
// @Accept json
// @Produce json
// @Param taskID path int true "Task ID"
// @Success 200 {object} types.TaskSwagger
// @Failure 400 {object} types.Errors
// @Failure 404 {object} types.Errors
// @Failure 500 {object} types.Errors
// @Router /v1/tasks/{taskID} [get]
func GetTaskHandler(ctx *fiber.Ctx, serviceContext service.ServiceContext) error {
	taskIdParam := ctx.Params("taskID")

	taskId, err := strconv.Atoi(taskIdParam)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&types.Errors{
			Status:  fiber.StatusBadRequest,
			Message: "Неверный формат ID",
			Error:   err.Error(),
		})
	}

	taskRepository := repository.NewTask(serviceContext)
	taskGet := tasks.NewGetTaskLogic(serviceContext, *taskRepository)

	result, errMsg, status := taskGet.GetTask(taskId)
	if errMsg != nil {
		return ctx.Status(status).JSON(errMsg)
	}

	return ctx.Status(status).JSON(result)
}
