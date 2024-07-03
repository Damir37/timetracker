package users

import (
	"github.com/gofiber/fiber/v2"
	"timertracker/internel/repository"
	"timertracker/internel/service"
	"timertracker/internel/types"
	"timertracker/internel/utils"
)

type GetUserLogic struct {
	ServiceContext service.ServiceContext
	userRepository repository.UserRep
}

func NewGetUserLogic(serviceContext service.ServiceContext, userRepository repository.UserRep) *GetUserLogic {
	return &GetUserLogic{
		ServiceContext: serviceContext,
		userRepository: userRepository,
	}
}

func (logic *GetUserLogic) GetUser(userId int) (result *types.User, errMsg *types.Errors, status int) {
	user, err := logic.userRepository.FindByID(userId)
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

	data := &types.User{
		Id:             user.ID,
		Name:           user.Name,
		PassportNumber: user.PassportNumber,
		Surname:        user.Surname,
		Patronymic:     user.Patronymic,
		Address:        user.Address,
		Tasks:          utils.ConvertTasks(user.Tasks),
	}

	return data, nil, status
}

func (logic *GetUserLogic) GetUsers() (result []*types.User, errMsg *types.Errors, status int) {
	all, err := logic.userRepository.FindAll()
	if err != nil {
		return nil, &types.Errors{
			Status:  fiber.StatusInternalServerError,
			Message: "Проблема с базой данных",
			Error:   err.Error(),
		}, fiber.StatusInternalServerError
	}

	for _, user := range all {
		data := &types.User{
			Id:             user.ID,
			Name:           user.Name,
			PassportNumber: user.PassportNumber,
			Surname:        user.Surname,
			Patronymic:     user.Patronymic,
			Address:        user.Address,
			Tasks:          utils.ConvertTasks(user.Tasks),
		}
		result = append(result, data)
	}

	return result, nil, fiber.StatusOK
}
