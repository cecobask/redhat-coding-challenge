package handler

import (
	"net/http"

	"github.com/go-chi/chi"
)

// NewHandler implements all routes for the application
func NewHandler() http.Handler {
	router := chi.NewRouter()
	router.Route("/objects", routeObjects)
	return router
}
