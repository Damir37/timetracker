package utils

import (
	"timertracker/internel/repository"
	"timertracker/internel/types"
)

func ConvertTasks(repoTasks []repository.Task) []types.Task {
	tasks := make([]types.Task, len(repoTasks))
	for i, task := range repoTasks {
		tasks[i] = types.Task{
			Description: task.Description,
			StartTime:   task.StartTime,
			EndTime:     task.EndTime,
			Duration:    task.Duration,
			Works:       ConvertWorks(task.Works),
		}
	}
	return tasks
}

func ConvertWorks(repoWorks []repository.Work) []types.Work {
	works := make([]types.Work, len(repoWorks))
	for i, work := range repoWorks {
		works[i] = types.Work{
			UserID:  work.UserID,
			Hours:   work.Hours,
			Minutes: work.Minutes,
			Date:    work.Date,
		}
	}
	return works
}
