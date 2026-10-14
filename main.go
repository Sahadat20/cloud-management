package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Sahadat20/cloud-management/internal/api"
	"github.com/Sahadat20/cloud-management/internal/service"
	"github.com/Sahadat20/cloud-management/internal/worker"
)

func main() {
	client := service.NewMockClient()

	// start jobqueue and worker pool

	workerCtx, workerCancel := context.WithCancel(context.Background())
	defer workerCancel()
	jobQueue := worker.NewJobQueue(100)
	processor := worker.NewProcessor(*client)
	workerPoll := worker.NewPool(5, jobQueue, processor.ProcessLaunch)

	workerPoll.Start(workerCtx)

	router := api.NewRouter(jobQueue)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router.Mux,
	}
	fmt.Println("Server running on http://localhost:8080")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Server failed:", err)
	}

}
