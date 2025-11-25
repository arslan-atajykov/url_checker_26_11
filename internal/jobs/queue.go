package jobs

import "url_checker/internal/repo"

type JobQueue struct {
	ch   chan int64
	repo repo.Repository
}

func NewJobQueue(r repo.Repository, size int) *JobQueue {
	return &JobQueue{
		ch:   make(chan int64, size),
		repo: r,
	}
}

func (q *JobQueue) Submit(jobID int64) {
	q.ch <- jobID
}
