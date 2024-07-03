package types

import (
	"encoding/json"
	"fmt"
	"time"
)

type User struct {
	Id             uint   `json:"id"`
	Name           string `json:"name"`
	PassportNumber string `json:"passportNumber" validate:"required,passport"`
	Surname        string `json:"surname"`
	Patronymic     string `json:"patronymic"`
	Address        string `json:"address"`
	Tasks          []Task `json:"-"`
}

type Task struct {
	Id          uint          `json:"id"`
	UserID      uint          `json:"userID"`
	Description string        `json:"description"`
	StartTime   time.Time     `json:"startTime"`
	EndTime     time.Time     `json:"endTime"`
	Duration    time.Duration `json:"duration"`
	Works       []Work        `json:"-"`
}

type TaskSwagger struct {
	Id          uint      `json:"id"`
	UserID      uint      `json:"userID"`
	Description string    `json:"description"`
	StartTime   time.Time `json:"startTime"`
	EndTime     time.Time `json:"endTime"`
	Duration    string    `json:"duration"`
	Works       []Work    `json:"-"`
}

type Work struct {
	Id      uint      `json:"id"`
	TaskID  uint      `json:"taskID"`
	UserID  uint      `json:"userID"`
	Hours   int       `json:"hours"`
	Minutes int       `json:"minutes"`
	Date    time.Time `json:"date"`
}

type WorkPeriod struct {
	UserID               uint      `json:"userID"`
	TotalHours           int       `json:"totalHours"`
	TotalMinutes         int       `json:"totalMinutes"`
	TotalMinutesCombined int       `json:"totalMinutesCombined"`
	StartDate            time.Time `json:"startDate"`
	EndDate              time.Time `json:"endDate"`
}

type Errors struct {
	Status  uint   `json:"status"`
	Message string `json:"message"`
	Error   string `json:"errors"`
}

type Result struct {
	Status  uint   `json:"status"`
	Message string `json:"message"`
}

func (t *Task) UnmarshalJSON(data []byte) error {
	type Alias Task
	aux := &struct {
		Duration interface{} `json:"duration"`
		*Alias
	}{
		Alias: (*Alias)(t),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	switch v := aux.Duration.(type) {
	case string:
		dur, err := time.ParseDuration(v)
		if err != nil {
			return err
		}
		t.Duration = dur
	case float64:
		t.Duration = time.Duration(v)
	default:
		return fmt.Errorf("unexpected type for duration: %T", v)
	}

	return nil
}
