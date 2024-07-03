package tasks

import (
	"github.com/gofiber/fiber/v2"
	"timertracker/internel/logic/tasks"
	"timertracker/internel/repository"
	"timertracker/internel/service"
	"timertracker/internel/types"
)

// AddTaskHandler godoc
// @Summary Добавляет в базу данных таску
// @Description Добавляет новую таску в систему
// @Tags tasks
// @Accept json
// @Produce json
// @Param task body types.TaskSwagger true "Объект таски, который нужно добавить"
// @Success 201 {object} types.TaskSwagger
// @Failure 500 {object} types.Errors
// @Router /v1/task [post]
func AddTaskHandler(ctx *fiber.Ctx, serviceContext service.ServiceContext) error {
	var task types.Task
	if err := ctx.BodyParser(&task); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&types.Errors{
			Status:  fiber.StatusBadRequest,
			Message: "Ваш запрос пустой",
			Error:   err.Error(),
		})
	}

	taskRepository := repository.NewTask(serviceContext)
	taskLogic := tasks.NewAddTaskLogic(serviceContext, *taskRepository)

	result, errMsg, status := taskLogic.AddTask(&task)
	if errMsg != nil {
		return ctx.Status(status).JSON(errMsg)
	}

	return ctx.Status(status).JSON(result)
}
