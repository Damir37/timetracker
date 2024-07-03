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

type AddUserLogic struct {
	ServiceContext service.ServiceContext
	userRepository repository.UserRep
}

func NewAddUserLogic(serviceContext service.ServiceContext, userRepository repository.UserRep) *AddUserLogic {
	return &AddUserLogic{
		ServiceContext: serviceContext,
		userRepository: userRepository,
	}
}

func (logic *AddUserLogic) AddUser(user *types.User) (result *types.User, errMsg *types.Errors, status int) {
	createUser := repository.User{
		Name:           user.Name,
		PassportNumber: user.PassportNumber,
		Surname:        user.Surname,
		Patronymic:     user.Patronymic,
		Address:        user.Address,
		Tasks:          []repository.Task{},
	}

	resultDb, err := logic.userRepository.Create(createUser)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			return nil, &types.Errors{
				Status:  fiber.StatusConflict,
				Message: "Пользователь уже существует",
				Error:   err.Error(),
			}, fiber.StatusConflict
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

	return result, nil, fiber.StatusCreated
}
