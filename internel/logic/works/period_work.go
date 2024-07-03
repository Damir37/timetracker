package works

import (
	"github.com/gofiber/fiber/v2"
	"time"
	"timertracker/internel/repository"
	"timertracker/internel/service"
	"timertracker/internel/types"
)

type GetPeriodWorkLogic struct {
	ServiceContext service.ServiceContext
	workRepository repository.WorkRep
}

func NewGetPeriodWorkLogic(serviceContext service.ServiceContext, workRepository repository.WorkRep) *GetPeriodWorkLogic {
	return &GetPeriodWorkLogic{
		ServiceContext: serviceContext,
		workRepository: workRepository,
	}
}

func (logic *GetPeriodWorkLogic) GetPeriodWork(userId uint, startDate, endDate time.Time) (result []types.WorkPeriod, errMsg *types.Errors, status int) {
	works, err := logic.workRepository.GetWorkByUserAndPeriod(userId, startDate, endDate)
	if err != nil {
		return nil, &types.Errors{
			Status:  fiber.StatusInternalServerError,
			Message: "База данных не работает",
			Error:   err.Error(),
		}, fiber.StatusInternalServerError
	}

	if len(works) == 0 {
		return nil, &types.Errors{
			Status:  fiber.StatusNotFound,
			Message: "Данные не найдены",
			Error:   "No data found for the given user and period",
		}, fiber.StatusNotFound
	}

	var resultWorks []types.WorkPeriod
	for _, work := range works {
		resultWorks = append(resultWorks, types.WorkPeriod{
			UserID:               work.UserID,
			TotalHours:           work.TotalHours,
			TotalMinutes:         work.TotalMinutes,
			TotalMinutesCombined: work.TotalMinutesCombined,
			StartDate:            startDate,
			EndDate:              endDate,
		})
	}

	return resultWorks, nil, fiber.StatusOK
}
