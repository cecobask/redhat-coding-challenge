package handler

import (
	"net/http"

	"github.com/cecobask/redhat-coding-challenge/db/postgres"
	"github.com/go-chi/chi"
)

// NewHandler implements all routes for the application
func NewHandler(pg postgres.Database) http.Handler {
	router := chi.NewRouter()
	objectsHandler := NewObjectsHandler(pg)
	router.Route("/objects", objectsHandler.InitObjectsRoute)
	return router
}
