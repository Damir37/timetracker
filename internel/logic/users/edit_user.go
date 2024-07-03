package users

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"timertracker/internel/repository"
	"timertracker/internel/service"
	"timertracker/internel/types"
	"timertracker/internel/utils"
)

type EditUserLogic struct {
	ServiceContext service.ServiceContext
	userRepository repository.UserRep
}

func NewEditUserLogic(serviceContext service.ServiceContext, userRepository repository.UserRep) *EditUserLogic {
	return &EditUserLogic{
		ServiceContext: serviceContext,
		userRepository: userRepository,
	}
}

func (logic *EditUserLogic) EditUser(user *types.User) (result *types.User, errMsg *types.Errors, status int) {
	editUser := repository.User{
		ID:             user.Id,
		Name:           user.Name,
		PassportNumber: user.PassportNumber,
		Surname:        user.Surname,
		Patronymic:     user.Patronymic,
		Address:        user.Address,
		Tasks:          []repository.Task{},
	}

	resultDb, err := logic.userRepository.Update(editUser)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case pgerrcode.UniqueViolation:
				return nil, &types.Errors{
					Status:  fiber.StatusConflict,
					Message: "Конфликт данных",
					Error:   err.Error(),
				}, fiber.StatusConflict
			case pgerrcode.ForeignKeyViolation:
				return nil, &types.Errors{
					Status:  fiber.StatusBadRequest,
					Message: "Нарушение внешнего ключа",
					Error:   err.Error(),
				}, fiber.StatusBadRequest
			}
		}

		if err == repository.ErrorMsg {
			return nil, &types.Errors{
				Status:  fiber.StatusNotFound,
				Message: "Пользователь не найден",
				Error:   err.Error(),
			}, fiber.StatusNotFound
		}

		return nil, &types.Errors{
			Status:  fiber.StatusInternalServerError,
			Message: "База данных не работает",
			Error:   err.Error(),
		}, fiber.StatusInternalServerError
	}

	result = &types.User{
		Id:             resultDb.ID,
		Name:           resultDb.Name,
		PassportNumber: resultDb.PassportNumber,
		Surname:        resultDb.Surname,
		Patronymic:     resultDb.Patronymic,
		Address:        resultDb.Address,
		Tasks:          utils.ConvertTasks(resultDb.Tasks),
	}

	return result, nil, fiber.StatusOK
}
