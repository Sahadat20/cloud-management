package worker

import (
	"context"

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

func (p *Processor) ProcessLaunch(ctx context.Context, job *Job) error {
	_, err := p.ec2Client.LaunchInstance(ctx, job.Name, job.InstanceType)
	return err
}

type ProcessorFunc func(ctx context.Context, job *Job) error
