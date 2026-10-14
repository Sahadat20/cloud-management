package worker

import (
	"context"
	"fmt"
	"log"
	"sync"
)

type Pool struct {
	numOfWorkers int
	queue        *JobQueue
	processor    ProcessorFunc
	wg           sync.WaitGroup
	started      bool
	mu           sync.Mutex
}

func NewPool(numWorkers int, queue *JobQueue, processor ProcessorFunc) *Pool {
	return &Pool{
		numOfWorkers: numWorkers,
		queue:        queue,
		processor:    processor,
	}
}

func (p *Pool) Start(ctx context.Context) {
	p.mu.Lock()
	if p.started {
		p.mu.Unlock()
		return
	}
	p.started = true
	p.mu.Unlock()
	log.Printf("[Workerpool] starting %d workers", p.numOfWorkers)
	for i := 0; i < p.numOfWorkers; i++ {
		p.wg.Add(1)
		go p.Worker(ctx, i)
	}
}

func (p *Pool) Worker(ctx context.Context, id int) {
	defer p.wg.Done()
	log.Printf("[Worker-%d] started", id)

	for {
		select {
		case <-ctx.Done():
			log.Printf("[Worker-%d] stopping context cancelled", id)
		case job, ok := <-p.queue.Jobs:
			if !ok {
				// Channel closed - shutdown signal
				log.Printf("[Worker-%d] Stopping (channel closed)", id)
				return
			}
			err := p.processor(ctx, job)
			if err != nil {
				fmt.Println("Instance lunch failed")
			}

		}
	}

}
