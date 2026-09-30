package main

import (
	"context"
	"fmt"
)

func main() {
	client := NewMockClient()
	instanceID, err := client.LaunchInstance(context.Background(), "ec2-mini", "ec2-mini-instance")
	if err != nil {
		fmt.Println("Instance lunch failed")
	}
	fmt.Printf("Intance created succefully %s", instanceID)
}
