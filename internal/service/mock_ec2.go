package service

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

type MockEC2Client struct {
	// latency
	minLatency time.Duration
	maxLatency time.Duration
}

func NewMockClient() *MockEC2Client {
	return &MockEC2Client{
		minLatency: 100 * time.Millisecond,
		maxLatency: 500 * time.Millisecond,
	}
}

func (m *MockEC2Client) generateInstanceID() string {
	const chars = "0123456789abcdef"
	id := make([]byte, 17)
	for i := range id {
		id[i] = chars[rand.Intn(len(chars))]
	}
	return "i-" + string(id)
}

func (m *MockEC2Client) LaunchInstance(ctx context.Context, instanceType, name string) (string, error) {
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	default:
	}
	instanceID := m.generateInstanceID()
	time.Sleep(3 * time.Second)
	fmt.Printf("[MockEC2] Launched instance %s (type: %s, name: %s)\n", instanceID, instanceType, name)
	return instanceID, nil
}
