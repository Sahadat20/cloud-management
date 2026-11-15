package worker

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type JobType string

const (
	JobTypeLaunch    JobType = "launch"
	JobTypeStart     JobType = "start"
	JobTypeStop      JobType = "stop"
	JobTypeTerminate JobType = "terminate"
)

type JobStatus string

const (
	JobStatusPending   JobStatus = "pending"
	JobStatusStart     JobStatus = "start"
	JobStatusStop      JobStatus = "stop"
	JobStatusTerminate JobStatus = "terminate"
)

type Job struct {
	ID         uuid.UUID       `json:"id"`
	Type       JobType         `json:"type"`
	Status     JobStatus       `json:"status"`
	Payload    json.RawMessage `json:"payload"`
	Result     json.RawMessage `json:"result,omitempty"`
	Error      string          `json:"error,omitempty"`
	CreatedAt  time.Time       `json:"crated_at"`
	UpdatedAt  time.Time       `json:"updated_at"`
	InstanceID *uuid.UUID      `json:"instance_id,omitempty"`
}

type LaunchPayload struct {
	Name         string `json:"name"`
	InstanceType string `json:"instance_type"`
}

type InstancePayload struct {
	InstanceID uuid.UUID `json:"instance_id"`
}

func NewJob(jobType JobType, payload interface{}) (*Job, error) {
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	return &Job{
		ID:        uuid.New(),
		Type:      jobType,
		Status:    JobStatusPending,
		Payload:   payloadBytes,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}

func (j *Job) ParseLaunchPayload() (*LaunchPayload, error) {
	var payload LaunchPayload
	if err := json.Unmarshal(j.Payload, &payload); err != nil {
		return nil, err
	}
	return &payload, nil
}
func (j *Job) ParseInstancePayload() (*InstancePayload, error) {
	var payload InstancePayload
	if err := json.Unmarshal(j.Payload, &payload); err != nil {
		return nil, err
	}
	return &payload, nil
}
