package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Sahadat20/cloud-management/internal/api/response"
	"github.com/Sahadat20/cloud-management/internal/worker"
	"github.com/google/uuid"
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
type JobResponse struct {
	JobID  string `json:"job_id"`
	Status string `json:"status"`
}

func (h *JobHandler) Launch(w http.ResponseWriter, r *http.Request) {
	var req JobRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.BadRequest(w, "Invalid JSON body")
		return
	}
	job, err := worker.CrateLaunchJob(req.Name, req.InstanceType)
	if err != nil {
		response.InternalError(w)
		return
	}
	if !h.queue.TrySubmit(job) {
		response.Error(w, http.StatusServiceUnavailable, "Job queu is full, try again later", "QUEUE_FULL")
		return
	}
	response.JSON(w, http.StatusAccepted, JobResponse{
		JobID:  job.ID.String(),
		Status: string(job.Status),
	})
}
func (h *JobHandler) Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, Cloud Manager!")
}

func (h *JobHandler) Start(w http.ResponseWriter, r *http.Request) {
	instanceID, err := uuid.Parse(r.PathValue("id"))

	job, err := worker.CreateStartJob(instanceID)
	if err != nil {
		response.InternalError(w)
		return
	}
	if !h.queue.TrySubmit(job) {
		response.Error(w, http.StatusServiceUnavailable, "Job queu is full, try again later", "QUEUE_FULL")
		return
	}
	response.JSON(w, http.StatusAccepted, JobResponse{
		JobID:  job.ID.String(),
		Status: string(job.Status),
	})
}
func (h *JobHandler) Stop(w http.ResponseWriter, r *http.Request) {
	instanceID, err := uuid.Parse(r.PathValue("id"))

	job, err := worker.CreateStopJob(instanceID)
	if err != nil {
		response.InternalError(w)
		return
	}
	if !h.queue.TrySubmit(job) {
		response.Error(w, http.StatusServiceUnavailable, "Job queu is full, try again later", "QUEUE_FULL")
		return
	}
	response.JSON(w, http.StatusAccepted, JobResponse{
		JobID:  job.ID.String(),
		Status: string(job.Status),
	})
}
func (h *JobHandler) Terminate(w http.ResponseWriter, r *http.Request) {
	instanceID, err := uuid.Parse(r.PathValue("id"))

	job, err := worker.CreateTerminateJob(instanceID)
	if err != nil {
		response.InternalError(w)
		return
	}
	if !h.queue.TrySubmit(job) {
		response.Error(w, http.StatusServiceUnavailable, "Job queu is full, try again later", "QUEUE_FULL")
		return
	}
	response.JSON(w, http.StatusAccepted, JobResponse{
		JobID:  job.ID.String(),
		Status: string(job.Status),
	})
}
