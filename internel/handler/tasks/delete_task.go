package tasks

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
	"timertracker/internel/logic/tasks"
	"timertracker/internel/repository"
	"timertracker/internel/service"
	"timertracker/internel/types"
)

// DeleteTaskHandler godoc
// @Summary Удаляет задачу из базы данных
// @Description Удаляет задачу из системы по её ID
// @Tags tasks
// @Accept json
// @Produce json
// @Param taskID path int true "Task ID"
// @Success 204 {object} types.Result
// @Failure 400 {object} types.Errors
// @Failure 404 {object} types.Errors
// @Failure 500 {object} types.Errors
// @Router /v1/tasks/delete/{taskID} [delete]
func DeleteTaskHandler(ctx *fiber.Ctx, serviceContext service.ServiceContext) error {
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
	taskDel := tasks.NewDeleteTaskLogic(serviceContext, *taskRepository)

	result, errMsg, status := taskDel.DeleteTask(taskId)
	if errMsg != nil {
		return ctx.Status(status).JSON(errMsg)
	}

	return ctx.Status(status).JSON(result)
}
