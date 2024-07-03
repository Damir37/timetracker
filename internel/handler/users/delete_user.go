package users

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
	"timertracker/internel/logic/users"
	"timertracker/internel/repository"
	"timertracker/internel/service"
	"timertracker/internel/types"
)

// DeleteUserHandler godoc
// @Summary Удаляет пользователя из базы данных
// @Description Удаляет пользователя из системы по его ID
// @Tags users
// @Accept json
// @Produce json
// @Param userId path int true "User ID"
// @Success 204 {object} types.Result
// @Failure 400 {object} types.Errors
// @Failure 404 {object} types.Errors
// @Failure 500 {object} types.Errors
// @Router /v1/user/delete/{userId} [delete]
func DeleteUserHandler(ctx *fiber.Ctx, serviceContext service.ServiceContext) error {
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
	userDel := users.NewDeleteUserLogic(serviceContext, *userRepository)

	result, errMsg, status := userDel.DeleteUser(userId)
	if errMsg != nil {
		return ctx.Status(status).JSON(errMsg)
	}

	return ctx.Status(status).JSON(result)
}
