package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cecobask/redhat-coding-challenge/db/postgres"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// ObjectsHandler ...
type ObjectsHandler interface {
	GetAllObjectsInBucket(http.ResponseWriter, *http.Request)
	GetObjectByBucketNameAndID(http.ResponseWriter, *http.Request)
	CreateOrUpdateObject(http.ResponseWriter, *http.Request)
	DeleteObject(http.ResponseWriter, *http.Request)
	InitObjectsRoute(chi.Router)
}

type objectsHandler struct {
	pg postgres.Database
}

type key int

const (
	keyBucket key = iota
	keyObjectID
)

// NewObjectsHandler ...
func NewObjectsHandler(pg postgres.Database) ObjectsHandler {
	return &objectsHandler{
		pg: pg,
	}
}

func (oh *objectsHandler) InitObjectsRoute(router chi.Router) {
	router.Route("/{bucket}", func(router chi.Router) {
		router.Use(bucketContext)
		router.Get("/", oh.GetAllObjectsInBucket)
		router.Route("/{object-id}", func(router chi.Router) {
			router.Use(objectContext)
			router.Get("/", oh.GetObjectByBucketNameAndID)
			router.Put("/", oh.CreateOrUpdateObject)
			router.Delete("/", oh.DeleteObject)
		})
	})
}

func bucketContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), keyBucket, chi.URLParam(r, "bucket"))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func objectContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), keyObjectID, chi.URLParam(r, "object-id"))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (oh *objectsHandler) GetAllObjectsInBucket(w http.ResponseWriter, r *http.Request) {
	bucketName := r.Context().Value(keyBucket).(string)
	objects, err := oh.pg.GetAllObjectsInBucket(r.Context(), bucketName)
	if err != nil {
		render.Render(w, r, errorRenderer(err, http.StatusInternalServerError, nil))
		return
	}
	if err := render.Render(w, r, objects); err != nil {
		render.Render(w, r, errorRenderer(err, http.StatusInternalServerError, nil))
	}
	return
}

func (oh *objectsHandler) GetObjectByBucketNameAndID(w http.ResponseWriter, r *http.Request) {
	bucketName := r.Context().Value(keyBucket).(string)
	objectID := r.Context().Value(keyObjectID).(string)
	object, err := oh.pg.GetObjectByBucketNameAndID(r.Context(), bucketName, objectID)
	if err != nil {
		if err == postgres.ErrNoRows {
			message := fmt.Sprintf("No object with ID %s found in bucket %s", objectID, bucketName)
			render.Render(w, r, errorRenderer(postgres.ErrNoRows, http.StatusNotFound, &message))
		} else {
			render.Render(w, r, errorRenderer(err, http.StatusInternalServerError, nil))
		}
		return
	}
	if err := render.Render(w, r, object); err != nil {
		render.Render(w, r, errorRenderer(err, http.StatusInternalServerError, nil))
	}
	return
}

func (oh *objectsHandler) CreateOrUpdateObject(w http.ResponseWriter, r *http.Request) {
	render.Render(w, r, errorRenderer(fmt.Errorf("TODO: implement createObject handler"), http.StatusNotImplemented, nil))
}

func (oh *objectsHandler) DeleteObject(w http.ResponseWriter, r *http.Request) {
	render.Render(w, r, errorRenderer(fmt.Errorf("TODO: implement deleteObject handler"), http.StatusNotImplemented, nil))
}
