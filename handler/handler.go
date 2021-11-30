package handler

import (
	"net/http"

	"github.com/cecobask/redhat-coding-challenge/db/files"
	"github.com/cecobask/redhat-coding-challenge/db/postgres"
	"github.com/go-chi/chi"
)

// NewHandler implements all routes for the application
func NewHandler(pg postgres.Database, fm files.FileManager) http.Handler {
	router := chi.NewRouter()
	objectsHandler := NewObjectsHandler(pg, fm)
	router.Route("/objects", objectsHandler.InitObjectsRoute)
	return router
}
