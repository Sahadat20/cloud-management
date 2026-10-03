package api

import (
	"net/http"

	"github.com/Sahadat20/cloud-management/internal/api/handlers"
	"github.com/Sahadat20/cloud-management/internal/worker"
)

type Router struct {
	Mux   *http.ServeMux
	queue *worker.JobQueue
}

func NewRouter(jobQueue *worker.JobQueue) *Router {
	router := &Router{
		Mux:   http.NewServeMux(),
		queue: jobQueue,
	}
	router.setupRoutes()
	return router
}

func (r *Router) setupRoutes() {
	jobHandler := handlers.NewJobHandler(r.queue)

	r.Mux.HandleFunc("/", jobHandler.Hello)
	r.Mux.HandleFunc("POST /instances", jobHandler.Launch)
}
