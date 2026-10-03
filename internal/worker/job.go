package worker

import (
	"context"
	"fmt"

	"github.com/Sahadat20/cloud-management/internal/service"
)

type Job struct {
	Name         string `json:"name"`
	InstanceType string `json:"instance_type"`
}

func Worker(client *service.MockEC2Client, jobs_queue *JobQueue) {
	for job := range jobs_queue.Jobs {
		instanceID, err := client.LaunchInstance(context.Background(), job.Name, job.InstanceType)
		if err != nil {
			fmt.Println("Instance lunch failed")
		}
		fmt.Printf("Intance created succefully %s \n", instanceID)
	}

}
