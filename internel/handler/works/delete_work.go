package works

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
	"timertracker/internel/logic/works"
	"timertracker/internel/repository"
	"timertracker/internel/service"
	"timertracker/internel/types"
)

// DeleteWorkHandler godoc
// @Summary Удаляет работу из базы данных
// @Description Удаляет работу из системы по её ID
// @Tags works
// @Accept json
// @Produce json
// @Param workID path int true "Work ID"
// @Success 204 {object} types.Result
// @Failure 400 {object} types.Errors
// @Failure 404 {object} types.Errors
// @Failure 500 {object} types.Errors
// @Router /v1/work/delete/{taskID} [delete]
func DeleteWorkHandler(ctx *fiber.Ctx, serviceContext service.ServiceContext) error {
	workIdParam := ctx.Params("workID")

	workId, err := strconv.Atoi(workIdParam)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&types.Errors{
			Status:  fiber.StatusBadRequest,
			Message: "Неверный формат ID",
			Error:   err.Error(),
		})
	}

	workRepository := repository.NewWork(serviceContext)
	workDel := works.NewDeleteWorkLogic(serviceContext, *workRepository)

	result, errMsg, status := workDel.DeleteWork(workId)
	if errMsg != nil {
		return ctx.Status(status).JSON(errMsg)
	}

	return ctx.Status(status).JSON(result)
}
