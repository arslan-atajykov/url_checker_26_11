package repo

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sync"

	"url_checker/internal/model"
)

type fileState struct {
	Counter int64        `json:"counter"`
	Tasks   []model.Task `json:"tasks"`
}

type FileRepo struct {
	mu       sync.Mutex
	counter  int64
	tasks    map[int64]model.Task
	filePath string
}

func NewFileRepo(path string) (*FileRepo, error) {
	r := &FileRepo{
		tasks:    make(map[int64]model.Task),
		filePath: path,
	}
	if err := r.loadFromDisk(); err != nil {
		return nil, err
	}
	return r, nil
}

func (r *FileRepo) CreateTaskWithLinks(links []model.LinkStruct) model.Task {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.counter++
	id := r.counter

	task := model.Task{
		ID:         id,
		Links:      links,
		TaskStatus: model.TaskCompleted,
	}

	r.tasks[id] = task
	_ = r.saveToDisk()
	return task
}

func (r *FileRepo) GetTask(id int64) (model.Task, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()

	task, ok := r.tasks[id]
	return task, ok
}

func (r *FileRepo) UpdateTask(task model.Task) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.tasks[task.ID] = task
	_ = r.saveToDisk()
}

func (r *FileRepo) saveToDisk() error {
	state := fileState{
		Counter: r.counter,
	}

	state.Tasks = make([]model.Task, 0, len(r.tasks))
	for _, t := range r.tasks {
		state.Tasks = append(state.Tasks, t)
	}

	if err := os.MkdirAll(filepath.Dir(r.filePath), 0o755); err != nil {
		return err
	}

	tmpPath := r.filePath + ".tmp"

	f, err := os.Create(tmpPath)
	if err != nil {
		return err
	}

	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	if err := enc.Encode(state); err != nil {
		f.Close()
		return err
	}

	if err := f.Close(); err != nil {
		return err
	}

	return os.Rename(tmpPath, r.filePath)
}

func (r *FileRepo) loadFromDisk() error {
	data, err := os.ReadFile(r.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	var state fileState
	if err := json.Unmarshal(data, &state); err != nil {
		return err
	}

	r.counter = state.Counter

	for _, t := range state.Tasks {
		r.tasks[t.ID] = t
	}

	return nil
}
