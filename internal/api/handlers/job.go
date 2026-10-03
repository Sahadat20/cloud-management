package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Sahadat20/cloud-management/internal/worker"
)

type JobHandler struct {
	queue *worker.JobQueue
}

func NewJobHandler(queue *worker.JobQueue) *JobHandler {
	return &JobHandler{
		queue: queue,
	}
}

type JobRequest struct {
	Name         string `json:"name"`
	InstanceType string `json:"instance_type"`
}

func (h *JobHandler) Launch(w http.ResponseWriter, r *http.Request) {
	var req JobRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Fprintln(w, "job submition faild")
		return
	}
	job := worker.Job{
		Name:         req.Name,
		InstanceType: req.InstanceType,
	}
	h.queue.Jobs <- &job
	fmt.Fprintln(w, "job submitted")
}
func (h *JobHandler) Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, Cloud Manager!")
}
