package users

import (
	"github.com/gofiber/fiber/v2"
	"timertracker/internel/logic/users"
	"timertracker/internel/repository"
	"timertracker/internel/service"
	"timertracker/internel/types"
)

// AddUserHandler godoc
// @Summary Добавляет в базу данных нового пользователя
// @Description Добавляет нового пользователя в систему
// @Tags users
// @Accept json
// @Produce json
// @Param user body types.User true "Объект пользователя, который нужно добавить"
// @Success 201 {object} types.User
// @Failure 400 {object} types.Errors
// @Failure 409 {object} types.Errors
// @Failure 500 {object} types.Errors
// @Router /v1/user [post]
func AddUserHandler(ctx *fiber.Ctx, serviceContext service.ServiceContext) error {
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
	userLogic := users.NewAddUserLogic(serviceContext, *userRepository)

	result, errMsg, status := userLogic.AddUser(&user)
	if errMsg != nil {
		return ctx.Status(status).JSON(errMsg)
	}

	return ctx.Status(status).JSON(result)
}
