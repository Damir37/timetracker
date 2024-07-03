package works

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
	"strings"
	"time"
	"timertracker/internel/logic/works"
	"timertracker/internel/repository"
	"timertracker/internel/service"
	"timertracker/internel/types"
)

// GetPeriodWorkHandler godoc
// @Summary Получает работы пользователя за период
// @Description Получает работы пользователя за указанный период времени
// @Tags works
// @Accept json
// @Produce json
// @Param userID path int true "User ID"
// @Param startDate query string true "Start Date Пример: 2023-10-05T21:48:00+07:00" Format(date-time)
// @Param endDate query string true "End Date Пример: 2023-10-05T21:48:00+07:00" Format(date-time)
// @Success 200 {array} types.WorkPeriod
// @Failure 400 {object} types.Errors
// @Failure 505 {object} types.Errors
// @Failure 500 {object} types.Errors
// @Router /v1/work/period/{userID} [get]
func GetPeriodWorkHandler(ctx *fiber.Ctx, serviceContext service.ServiceContext) error {
	userIDParam := ctx.Params("userID")

	userID, err := strconv.ParseUint(userIDParam, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&types.Errors{
			Status:  fiber.StatusBadRequest,
			Message: "Неверный формат userID",
			Error:   err.Error(),
		})
	}

	startDateParam := ctx.Query("startDate")
	endDateParam := ctx.Query("endDate")

	startDate, err := time.Parse("2006-01-02T15:04:05Z07:00", strings.Replace(startDateParam, " ", "+", 1))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&types.Errors{
			Status:  fiber.StatusBadRequest,
			Message: "Неверный формат startDate",
			Error:   err.Error(),
		})
	}

	endDate, err := time.Parse("2006-01-02T15:04:05Z07:00", strings.Replace(endDateParam, " ", "+", 1))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&types.Errors{
			Status:  fiber.StatusBadRequest,
			Message: "Неверный формат endDate",
			Error:   err.Error(),
		})
	}

	workRepository := repository.NewWork(serviceContext)
	getPeriodWorkLogic := works.NewGetPeriodWorkLogic(serviceContext, *workRepository)

	result, errMsg, status := getPeriodWorkLogic.GetPeriodWork(uint(userID), startDate, endDate)
	if errMsg != nil {
		return ctx.Status(status).JSON(errMsg)
	}

	return ctx.Status(status).JSON(result)
}
