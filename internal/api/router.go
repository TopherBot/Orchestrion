package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

// RegisterRoutes mounts all API endpoints onto the supplied router.
func RegisterRoutes(r chi.Router) {
	// Versioned subrouter (v1)
	v1 := chi.NewRouter()
	v1.Get("/pipelines", listPipelines)
	v1.Post("/pipelines", createPipeline)
	v1.Get("/pipelines/{id}", getPipeline)
	v1.Post("/pipelines/{id}/run", runPipeline)

	r.Mount("/api/v1", v1)
}

func listPipelines(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusOK); w.Write([]byte(`[]`)) }
func createPipeline(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusCreated) }
func getPipeline(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusOK) }
func runPipeline(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusAccepted) }
