package worker

import "log"

type JobQueue struct {
	Jobs chan *Job
	Size int
}

func NewJobQueue(size int) *JobQueue {
	return &JobQueue{
		Jobs: make(chan *Job, size),
		Size: size,
	}

}

func (q *JobQueue) TrySubmit(job *Job) bool {
	select {
	case q.Jobs <- job:
		log.Printf("[jobqueue] Job %s submitted (type: %s)", job.ID, job.Type)
		return true
	default:
		log.Printf("[JobQueue] Queue full, job %s rejected", job.ID)
		return false
	}
}
