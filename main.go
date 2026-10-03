package main

import (
	"fmt"
	"net/http"

	"github.com/Sahadat20/cloud-management/internal/api"
	"github.com/Sahadat20/cloud-management/internal/service"
	"github.com/Sahadat20/cloud-management/internal/worker"
)

func main() {
	client := service.NewMockClient()
	jobQueue := worker.NewJobQueue(100)

	go worker.Worker(client, jobQueue)

	// mux := http.NewServeMux()

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
