package repo

import (
	"sync"
	"url_checker/internal/model"
)

type MemoryRepo struct {
	mu      sync.Mutex
	counter int64
	tasks   map[int64]model.Task
}

func NewMemoryRepo() *MemoryRepo {
	return &MemoryRepo{
		counter: 0,
		tasks:   make(map[int64]model.Task),
	}
}

func (r *MemoryRepo) CreateTask(urls []string) model.Task {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.counter++
	id := r.counter

	links := make([]model.LinkStruct, len(urls))

	for i, url := range urls {
		links[i] = model.LinkStruct{URL: url, Lstatus: model.LStatus(model.StatusUnavailable)}
	}

	task := model.Task{
		ID:         id,
		Links:      links,
		TaskStatus: model.TaskPending,
	}

	r.tasks[id] = task
	return task
}
func (r *MemoryRepo) GetTask(id int64) (model.Task, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()

	task, ok := r.tasks[id]
	return task, ok
}

func (r *MemoryRepo) UpdateTask(task model.Task) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.tasks[task.ID] = task
}

func (r *MemoryRepo) CreateTaskWithLinks(links []model.LinkStruct) model.Task {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.counter++
	id := r.counter

	task := model.Task{
		ID:         id,
		Links:      links,
		TaskStatus: model.TaskCompleted, // sync обработка завершена сразу
	}

	r.tasks[id] = task
	return task
}
