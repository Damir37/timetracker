package works

import (
	"github.com/gofiber/fiber/v2"
	"timertracker/internel/logic/works"
	"timertracker/internel/repository"
	"timertracker/internel/service"
	"timertracker/internel/types"
)

// AddWorkHandler godoc
// @Summary Добавляет в базу данных работу, которую взял пользователь
// @Description Добавляет новую работу в систему
// @Tags works
// @Accept json
// @Produce json
// @Param task body types.Work true "Объект работы, который нужно добавить"
// @Success 201 {object} types.Work
// @Failure 500 {object} types.Errors
// @Router /v1/work [post]
func AddWorkHandler(ctx *fiber.Ctx, serviceContext service.ServiceContext) error {
	var work types.Work
	if err := ctx.BodyParser(&work); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&types.Errors{
			Status:  fiber.StatusBadRequest,
			Message: "Ваш запрос пустой",
			Error:   err.Error(),
		})
	}

	workRepository := repository.NewWork(serviceContext)
	workLogic := works.NewAddWorkLogic(serviceContext, *workRepository)

	result, errMsg, status := workLogic.AddWork(&work)
	if errMsg != nil {
		return ctx.Status(status).JSON(errMsg)
	}

	return ctx.Status(status).JSON(result)
}
