package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Sahadat20/cloud-management/internal/api"
	"github.com/Sahadat20/cloud-management/internal/config"
	"github.com/Sahadat20/cloud-management/internal/db"
	"github.com/Sahadat20/cloud-management/internal/service"
	"github.com/Sahadat20/cloud-management/internal/worker"
)

func main() {
	log.Println("Starting cloud management server....")

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}
	log.Printf("Configuration loaded:  db=%s:%s/%s",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.DBName)

	// conntect to database
	database, err := db.New(&cfg.Database)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close()
	log.Println("Connected to PostgreSQL database")

	// Verify database health
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := database.Health(ctx); err != nil {
		log.Fatalf("Database health check failed: %v", err)
	}
	log.Println("Database health check passed")
	// Initialize ec2 client
	client := service.NewMockClient()

	// start jobqueue and worker pool

	workerCtx, workerCancel := context.WithCancel(context.Background())
	defer workerCancel()
	jobQueue := worker.NewJobQueue(100)
	processor := worker.NewProcessor(*client)
	workerPoll := worker.NewPool(5, jobQueue, processor.Process)

	workerPoll.Start(workerCtx)

	router := api.NewRouter(jobQueue)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router.Mux,
	}
	fmt.Println("Server running on http://localhost:8080")
	err = server.ListenAndServe()
	if err != nil {
		fmt.Println("Server failed:", err)
	}

}
