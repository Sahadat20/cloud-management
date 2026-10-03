package worker

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
