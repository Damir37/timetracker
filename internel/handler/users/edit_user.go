package users

import (
	"github.com/gofiber/fiber/v2"
	"timertracker/internel/logic/users"
	"timertracker/internel/repository"
	"timertracker/internel/service"
	"timertracker/internel/types"
)

// EditUserHandler godoc
// @Summary Редактирует пользователя в базе данных
// @Description Редактирует существующего пользователя в системе
// @Tags users
// @Accept json
// @Produce json
// @Param task body types.User true "Объект пользователя, которую необходимо обновить"
// @Success 200 {object} types.User
// @Failure 400 {object} types.Errors
// @Failure 404 {object} types.Errors
// @Failure 500 {object} types.Errors
// @Router /v1/user/edit [put]
func EditUserHandler(ctx *fiber.Ctx, serviceContext service.ServiceContext) error {
	var user types.User
	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&types.Errors{
			Status:  fiber.StatusBadRequest,
			Message: "Ваш запрос пустой",
			Error:   err.Error(),
		})
	}

	if err := serviceContext.Valider.Struct(user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&types.Errors{
			Status:  fiber.StatusBadRequest,
			Message: "Данные не точные",
			Error:   err.Error(),
		})
	}

	userRepository := repository.NewUserRepository(serviceContext)
	userLogic := users.NewEditUserLogic(serviceContext, *userRepository)

	userEdit, err, status := userLogic.EditUser(&user)
	if err != nil {
		return ctx.Status(status).JSON(err)
	}

	return ctx.Status(status).JSON(userEdit)
}
