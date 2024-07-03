package users

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
	"timertracker/internel/logic/users"
	"timertracker/internel/repository"
	"timertracker/internel/service"
	"timertracker/internel/types"
)

// GetUsersHandler godoc
// @Summary Получает всех пользователей из базы данных
// @Description Получает список всех пользователей в системе
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} types.User
// @Failure 500 {object} types.Errors
// @Router /v1/users [get]
func GetUsersHandler(ctx *fiber.Ctx, serviceContext service.ServiceContext) error {
	userRepository := repository.NewUserRepository(serviceContext)
	userGet := users.NewGetUserLogic(serviceContext, *userRepository)

	result, errMsg, status := userGet.GetUsers()
	if errMsg != nil {
		return ctx.Status(status).JSON(errMsg)
	}

	return ctx.Status(status).JSON(result)
}

// GetUserHandler godoc
// @Summary Получает пользователя из базы данных по ID
// @Description Получает пользователя из системы по его ID
// @Tags users
// @Accept json
// @Produce json
// @Param userId path int true "User ID"
// @Success 200 {object} types.User
// @Failure 400 {object} types.Errors
// @Failure 404 {object} types.Errors
// @Failure 500 {object} types.Errors
// @Router /v1/user/{userId} [get]
func GetUserHandler(ctx *fiber.Ctx, serviceContext service.ServiceContext) error {
	userIdParam := ctx.Params("userId")

	userId, err := strconv.Atoi(userIdParam)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&types.Errors{
			Status:  fiber.StatusBadRequest,
			Message: "Неверный формат ID",
			Error:   err.Error(),
		})
	}

	userRepository := repository.NewUserRepository(serviceContext)
	userGet := users.NewGetUserLogic(serviceContext, *userRepository)

	result, errMsg, status := userGet.GetUser(userId)
	if errMsg != nil {
		return ctx.Status(status).JSON(errMsg)
	}

	return ctx.Status(status).JSON(result)
}
