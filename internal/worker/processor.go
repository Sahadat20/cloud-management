package worker

import (
	"context"
	"fmt"

	"github.com/Sahadat20/cloud-management/internal/service"
)

type Processor struct {
	ec2Client service.MockEC2Client
}

func NewProcessor(ec2Clint service.MockEC2Client) *Processor {
	return &Processor{
		ec2Client: ec2Clint,
	}
}

func (p *Processor) Process(ctx context.Context, job *Job) error {
	jobType := job.Type
	switch jobType {
	case JobTypeLaunch:
		return p.ProcessLaunch(ctx, job)
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

type ProcessorFunc func(ctx context.Context, job *Job) error

func CrateLaunchJob(name string, instanceType string) (*Job, error) {
	return NewJob(JobTypeLaunch, LaunchPayload{
		Name:         name,
		InstanceType: instanceType,
	})
}
