package jobs

import (
	"url_checker/internal/checker"
	"url_checker/internal/model"
)

func (q *JobQueue) StartWorker() {
	go func() {
		for jobID := range q.ch {
			task, ok := q.repo.GetTask(jobID)
			if !ok {
				continue
			}

			task.TaskStatus = model.TaskRunning
			q.repo.UpdateTask(task)

			for i := range task.Links {
				status := checker.CheckURL(task.Links[i].URL)
				task.Links[i].Lstatus = model.LStatus(status)
			}

			task.TaskStatus = model.TaskCompleted
			q.repo.UpdateTask(task)
		}
	}()
}
