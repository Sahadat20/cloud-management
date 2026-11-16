package worker

import (
	"log"
	"sync"

	"github.com/google/uuid"
)

type JobQueue struct {
	Jobs    chan *Job
	results map[uuid.UUID]*Job
	mu      sync.RWMutex
	Size    int
}

func NewJobQueue(size int) *JobQueue {
	return &JobQueue{
		Jobs:    make(chan *Job, size),
		results: make(map[uuid.UUID]*Job),
		Size:    size,
	}

}

func (q *JobQueue) TrySubmit(job *Job) bool {
	q.mu.Lock()
	q.results[job.ID] = job
	q.mu.Unlock()

	select {
	case q.Jobs <- job:
		log.Printf("[jobqueue] Job %s submitted (type: %s)", job.ID, job.Type)
		return true
	default:
		//queue is full remove from results
		q.mu.Lock()
		delete(q.results, job.ID)
		q.mu.Unlock()
		log.Printf("[JobQueue] Queue full, job %s rejected", job.ID)
		return false
	}
}

func (q *JobQueue) UpdateJob(job *Job) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.results[job.ID] = job
}
