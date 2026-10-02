package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type JobHandler struct {
	queue chan Job
}

func NewJobHandler(queue chan Job) *JobHandler {
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
	job := Job{
		Name:         req.Name,
		InstanceType: req.InstanceType,
	}
	h.queue <- job
	fmt.Fprintln(w, "job submitted")
}
func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, Cloud Manager!")
}

func main() {
	client := NewMockClient()
	jobs_queue := make(chan Job, 10)
	jobHandler := NewJobHandler(jobs_queue)
	http.HandleFunc("/", Hello)
	http.HandleFunc("POST /instances", jobHandler.Launch)
	go worker(client, jobs_queue)

	fmt.Println("Server running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server failed:", err)
	}

}

type Job struct {
	Name         string `json:"name"`
	InstanceType string `json:"instance_type"`
}

func worker(client *MockEC2Client, jobs_queue chan Job) {
	for job := range jobs_queue {
		instanceID, err := client.LaunchInstance(context.Background(), job.Name, job.InstanceType)
		if err != nil {
			fmt.Println("Instance lunch failed")
		}
		fmt.Printf("Intance created succefully %s \n", instanceID)
	}

}
