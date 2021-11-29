package handler

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jackc/pgx/v4"
)

// NewHandler implements all routes for the application
func NewHandler(postgresConn *pgx.Conn) http.Handler {
	router := chi.NewRouter()
	objectsHandler := NewObjectsHandler(postgresConn)
	router.Route("/objects", objectsHandler.InitObjectsRoute)
	return router
}
