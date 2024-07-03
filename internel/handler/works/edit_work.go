package works

import (
	"github.com/gofiber/fiber/v2"
	"timertracker/internel/logic/works"
	"timertracker/internel/repository"
	"timertracker/internel/service"
	"timertracker/internel/types"
)

// EditWorkHandler godoc
// @Summary Редактирует работу в базе данных
// @Description Редактирует существующею работу в системе
// @Tags works
// @Accept json
// @Produce json
// @Param task body types.Work true "Объект работы, которую необходимо обновить"
// @Success 200 {object} types.Work
// @Failure 400 {object} types.Errors
// @Failure 404 {object} types.Errors
// @Failure 500 {object} types.Errors
// @Router /v1/work/edit [put]
func EditWorkHandler(ctx *fiber.Ctx, serviceContext service.ServiceContext) error {
	var work types.Work

	if err := ctx.BodyParser(&work); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&types.Errors{
			Status:  fiber.StatusBadRequest,
			Message: "Ваш запрос пустой",
			Error:   err.Error(),
		})
	}

	workRepository := repository.NewWork(serviceContext)
	workLogic := works.NewEditWorkLogic(serviceContext, *workRepository)

	userEdit, err, status := workLogic.EditWork(&work)
	if err != nil {
		return ctx.Status(status).JSON(err)
	}

	return ctx.Status(status).JSON(userEdit)
}
