package users

import (
	"github.com/gofiber/fiber/v2"
	"timertracker/internel/repository"
	"timertracker/internel/service"
	"timertracker/internel/types"
)

type DeleteUserLogic struct {
	ServiceContext service.ServiceContext
	userRepository repository.UserRep
}

func NewDeleteUserLogic(serviceContext service.ServiceContext, userRepository repository.UserRep) *DeleteUserLogic {
	return &DeleteUserLogic{
		ServiceContext: serviceContext,
		userRepository: userRepository,
	}
}

func (logic *DeleteUserLogic) DeleteUser(userId int) (result *types.Result, errMsg *types.Errors, status int) {
	err := logic.userRepository.Delete(userId)

	if err != nil {
		if err == repository.ErrorMsg {
			return nil, &types.Errors{
				Status:  fiber.StatusNotFound,
				Message: "Пользователь не найден",
				Error:   err.Error(),
			}, fiber.StatusNotFound
		}
		return nil, &types.Errors{
			Status:  fiber.StatusInternalServerError,
			Message: "Проблема с базой данных",
			Error:   err.Error(),
		}, fiber.StatusInternalServerError
	}

	return &types.Result{
		Status:  fiber.StatusNoContent,
		Message: "Пользователь удалён",
	}, nil, fiber.StatusNoContent
}
