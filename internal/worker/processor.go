package worker

import (
	"context"
	"fmt"

	"github.com/Sahadat20/cloud-management/internal/service"
	"github.com/google/uuid"
)

type Processor struct {
	ec2Client service.MockEC2Client
}

func NewProcessor(ec2Clint service.MockEC2Client) *Processor {
	return &Processor{
		ec2Client: ec2Clint,
	}
}

type ProcessorFunc func(ctx context.Context, job *Job) error

func (p *Processor) Process(ctx context.Context, job *Job) error {
	jobType := job.Type
	switch jobType {
	case JobTypeLaunch:
		return p.ProcessLaunch(ctx, job)
	case JobTypeStart:
		return p.ProcessStart(ctx, job)
	case JobTypeStop:
		return p.ProcessStop(ctx, job)
	case JobTypeTerminate:
		return p.ProcessTerminate(ctx, job)
	default:
		return fmt.Errorf("unknown job type %s", jobType)
	}

}
func (p *Processor) ProcessLaunch(ctx context.Context, job *Job) error {
	payload, err := job.ParseLaunchPayload()
	if err != nil {
		return fmt.Errorf("invalid payload: %w", err)
	}
	_, err = p.ec2Client.LaunchInstance(ctx, payload.InstanceType, payload.Name)
	return err
}
func (p *Processor) ProcessStart(ctx context.Context, job *Job) error {
	payload, err := job.ParseInstancePayload()
	if err != nil {
		return fmt.Errorf("invalid payload: %w", err)
	}
	err = p.ec2Client.StartInstance(ctx, payload.InstanceID)
	return err
}
func (p *Processor) ProcessStop(ctx context.Context, job *Job) error {
	payload, err := job.ParseInstancePayload()
	if err != nil {
		return fmt.Errorf("invalid payload: %w", err)
	}
	err = p.ec2Client.StopInstance(ctx, payload.InstanceID)
	return err
}
func (p *Processor) ProcessTerminate(ctx context.Context, job *Job) error {
	payload, err := job.ParseInstancePayload()
	if err != nil {
		return fmt.Errorf("invalid payload: %w", err)
	}
	err = p.ec2Client.TerminateInstance(ctx, payload.InstanceID)
	return err
}

func CrateLaunchJob(name string, instanceType string) (*Job, error) {
	return NewJob(JobTypeLaunch, LaunchPayload{
		Name:         name,
		InstanceType: instanceType,
	})
}

func CreateStartJob(instanceID uuid.UUID) (*Job, error) {
	return NewJob(JobTypeStart, InstancePayload{
		InstanceID: instanceID,
	})
}
func CreateStopJob(instanceID uuid.UUID) (*Job, error) {
	return NewJob(JobTypeStop, InstancePayload{
		InstanceID: instanceID,
	})
}
func CreateTerminateJob(instanceID uuid.UUID) (*Job, error) {
	return NewJob(JobTypeTerminate, InstancePayload{
		InstanceID: instanceID,
	})
}
