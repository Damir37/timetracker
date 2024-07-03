package repository

import (
	"errors"
	"gorm.io/gorm"
	"time"
	"timertracker/internel/service"
)

type WorkRepository interface {
	FindAll() ([]Work, error)
	FindByID(id int) (Work, error)
	GetWorkByUserAndPeriod(userID uint, startDate, endDate time.Time) ([]WorkSummary, error)
	Create(work Work) (Work, error)
	Update(work Work) (Work, error)
	Delete(id int) error
}

type WorkRep struct {
	ServiceContext service.ServiceContext
}

type Work struct {
	gorm.Model
	ID      uint `gorm:"primary_key;type:bigint"`
	TaskID  uint
	UserID  uint
	Hours   int
	Minutes int
	Date    time.Time
}

type WorkSummary struct {
	UserID               uint `gorm:"column:user_id"`
	TotalHours           int  `gorm:"column:total_hours"`
	TotalMinutes         int  `gorm:"column:total_minutes"`
	TotalMinutesCombined int  `gorm:"column:total_minutes_combined"`
}

func NewWork(serviceContext service.ServiceContext) *WorkRep {
	return &WorkRep{
		ServiceContext: serviceContext,
	}
}

func (w *WorkRep) FindAll() ([]Work, error) {
	var works []Work
	result := w.ServiceContext.DB.Find(&works)
	if result.Error != nil {
		return nil, result.Error
	}
	return works, nil
}

func (w *WorkRep) FindByID(id int) (Work, error) {
	var work Work
	result := w.ServiceContext.DB.First(&work, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return Work{}, ErrorMsg
		}
		return Work{}, result.Error
	}
	return work, nil
}

func (w *WorkRep) GetWorkByUserAndPeriod(userID uint, startDate, endDate time.Time) ([]WorkSummary, error) {
	var results []WorkSummary
	err := w.ServiceContext.DB.Model(&Work{}).
		Select("user_id, SUM(hours) as total_hours, SUM(minutes) as total_minutes, (SUM(hours) * 60 + SUM(minutes)) as total_minutes_combined").
		Where("date BETWEEN ? AND ? AND user_id = ?", startDate, endDate, userID).
		Group("user_id").
		Order("total_minutes_combined DESC").
		Find(&results).Error
	return results, err
}

func (w *WorkRep) Create(work Work) (Work, error) {
	result := w.ServiceContext.DB.Create(&work)
	if result.Error != nil {
		return Work{}, result.Error
	}
	return work, nil
}

func (w *WorkRep) Update(work Work) (Work, error) {
	result := w.ServiceContext.DB.Save(&work)
	if result.Error != nil {
		return Work{}, result.Error
	}
	return work, nil
}

func (w *WorkRep) Delete(id int) error {
	result := w.ServiceContext.DB.Delete(&Work{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrorMsg
	}
	return nil
}
