package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/jackc/pgx/v4"
)

// ObjectsHandler ...
type ObjectsHandler interface {
	GetAllObjects(http.ResponseWriter, *http.Request)
	GetObject(http.ResponseWriter, *http.Request)
	CreateOrUpdateObject(http.ResponseWriter, *http.Request)
	DeleteObject(http.ResponseWriter, *http.Request)
	InitObjectsRoute(chi.Router)
}

type objectsHandler struct {
	postgresConn *pgx.Conn
}

type key int

const (
	keyBucket key = iota
)

// NewObjectsHandler ...
func NewObjectsHandler(postgresConn *pgx.Conn) ObjectsHandler {
	return &objectsHandler{
		postgresConn: postgresConn,
	}
}

func (oh *objectsHandler) InitObjectsRoute(router chi.Router) {
	router.Route("/{bucket}", func(router chi.Router) {
		router.Use(contextObjects)
		router.Get("/", oh.GetAllObjects)
		router.Get("/{objectID}", oh.GetObject)
		router.Put("/{objectID}", oh.CreateOrUpdateObject)
		router.Delete("/{objectID}", oh.DeleteObject)
	})
}

func contextObjects(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bucket := chi.URLParam(r, "bucket")
		ctx := context.WithValue(r.Context(), keyBucket, bucket)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (oh *objectsHandler) GetAllObjects(w http.ResponseWriter, r *http.Request) {
	render.Render(w, r, errorRenderer(fmt.Errorf("TODO: implement getAllObjects handler"), http.StatusNotImplemented, nil))
}

func (oh *objectsHandler) GetObject(w http.ResponseWriter, r *http.Request) {
	render.Render(w, r, errorRenderer(fmt.Errorf("TODO: implement getObject handler"), http.StatusNotImplemented, nil))
}

func (oh *objectsHandler) CreateOrUpdateObject(w http.ResponseWriter, r *http.Request) {
	render.Render(w, r, errorRenderer(fmt.Errorf("TODO: implement createObject handler"), http.StatusNotImplemented, nil))
}

func (oh *objectsHandler) DeleteObject(w http.ResponseWriter, r *http.Request) {
	render.Render(w, r, errorRenderer(fmt.Errorf("TODO: implement deleteObject handler"), http.StatusNotImplemented, nil))
}
