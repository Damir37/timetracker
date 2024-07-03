package repository

import (
	"errors"
	"gorm.io/gorm"
	"time"
	"timertracker/internel/service"
)

type TaskRepository interface {
	FindAll() ([]Task, error)
	FindByID(id int) (Task, error)
	Create(task Task) (Task, error)
	Update(task Task) (Task, error)
	Delete(id int) error
}

type TaskRep struct {
	ServiceContext service.ServiceContext
}

type Task struct {
	gorm.Model
	UserID      uint
	Description string `gorm:"type:varchar(255)"`
	StartTime   time.Time
	EndTime     time.Time
	Duration    time.Duration
	Works       []Work `gorm:"foreignKey:TaskID"`
}

func NewTask(serviceContext service.ServiceContext) *TaskRep {
	return &TaskRep{
		ServiceContext: serviceContext,
	}
}

func (t *TaskRep) FindAll() ([]Task, error) {
	var tasks []Task
	result := t.ServiceContext.DB.Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}
	return tasks, nil
}

func (t *TaskRep) FindByID(id int) (Task, error) {
	var task Task
	result := t.ServiceContext.DB.First(&task, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return Task{}, ErrorMsg
		}
		return Task{}, result.Error
	}
	return task, nil
}

func (t *TaskRep) Create(task Task) (Task, error) {
	result := t.ServiceContext.DB.Create(&task)
	if result.Error != nil {
		return Task{}, result.Error
	}
	return task, nil
}

func (t *TaskRep) Update(task Task) (Task, error) {
	result := t.ServiceContext.DB.Save(&task)
	if result.Error != nil {
		return Task{}, result.Error
	}
	return task, nil
}

func (t *TaskRep) Delete(id int) error {
	result := t.ServiceContext.DB.Delete(&Task{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrorMsg
	}
	return nil
}
