package repo

import "url_checker/internal/model"

type Repository interface {
	CreateTask(links []string) model.Task
	GetTask(id int64) (model.Task, bool)
}
