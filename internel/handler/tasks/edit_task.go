package tasks

import (
	"github.com/gofiber/fiber/v2"
	"timertracker/internel/logic/tasks"
	"timertracker/internel/repository"
	"timertracker/internel/service"
	"timertracker/internel/types"
)

// EditTaskHandler godoc
// @Summary Редактирует задачу в базе данных
// @Description Редактирует существующую задачу в системе
// @Tags tasks
// @Accept json
// @Produce json
// @Param task body types.TaskSwagger true "Объект задачи, которую необходимо обновить"
// @Success 200 {object} types.TaskSwagger
// @Failure 400 {object} types.Errors
// @Failure 404 {object} types.Errors
// @Failure 500 {object} types.Errors
// @Router /v1/tasks/edit [put]
func EditTaskHandler(ctx *fiber.Ctx, serviceContext service.ServiceContext) error {
	var task types.Task
	if err := ctx.BodyParser(&task); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&types.Errors{
			Status:  fiber.StatusBadRequest,
			Message: "Ваш запрос пустой",
			Error:   err.Error(),
		})
	}

	taskRepository := repository.NewTask(serviceContext)
	taskLogic := tasks.NewEditTaskLogic(serviceContext, *taskRepository)

	userEdit, err, status := taskLogic.EditTask(&task)
	if err != nil {
		return ctx.Status(status).JSON(err)
	}

	return ctx.Status(status).JSON(userEdit)
}
