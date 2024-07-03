package works

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
	"timertracker/internel/logic/works"
	"timertracker/internel/repository"
	"timertracker/internel/service"
	"timertracker/internel/types"
)

// GetWorksHandler godoc
// @Summary Получает все работы из базы данных
// @Description Получает список всех работ в системе
// @Tags works
// @Accept json
// @Produce json
// @Success 200 {array} types.Work
// @Failure 500 {object} types.Errors
// @Router /v1/works [get]
func GetWorksHandler(ctx *fiber.Ctx, serviceContext service.ServiceContext) error {
	workRepository := repository.NewWork(serviceContext)
	workGet := works.NewGetWorkLogic(serviceContext, *workRepository)

	result, errMsg, status := workGet.GetWorks()
	if errMsg != nil {
		return ctx.Status(status).JSON(errMsg)
	}

	return ctx.Status(status).JSON(result)
}

// GetWorkHandler godoc
// @Summary Получает работу из базы данных по ID
// @Description Получает работу из системы по её ID
// @Tags works
// @Accept json
// @Produce json
// @Param workID path int true "Work ID"
// @Success 200 {object} types.Work
// @Failure 400 {object} types.Errors
// @Failure 404 {object} types.Errors
// @Failure 500 {object} types.Errors
// @Router /v1/work/{workID} [get]
func GetWorkHandler(ctx *fiber.Ctx, serviceContext service.ServiceContext) error {
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
	workGet := works.NewGetWorkLogic(serviceContext, *workRepository)

	result, errMsg, status := workGet.GetWork(workId)
	if errMsg != nil {
		return ctx.Status(status).JSON(errMsg)
	}

	return ctx.Status(status).JSON(result)
}
