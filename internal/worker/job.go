package worker

type Job struct {
	Name         string `json:"name"`
	InstanceType string `json:"instance_type"`
}
