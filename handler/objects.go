package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type key int

const (
	keyBucket key = iota
)

func routeObjects(router chi.Router) {
	router.Route("/{bucket}", func(router chi.Router) {
		router.Use(contextObjects)
		router.Get("/", getAllObjects)
		router.Get("/{objectID}", getObject)
		router.Put("/{objectID}", createObject)
		router.Delete("/{objectID}", deleteObject)
	})
}

func contextObjects(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bucket := chi.URLParam(r, "bucket")
		ctx := context.WithValue(r.Context(), keyBucket, bucket)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getAllObjects(w http.ResponseWriter, r *http.Request) {
	render.Render(w, r, errorRenderer(fmt.Errorf("TODO: implement getAllObjects handler"), http.StatusNotImplemented, nil))
}

func getObject(w http.ResponseWriter, r *http.Request) {
	render.Render(w, r, errorRenderer(fmt.Errorf("TODO: implement getObject handler"), http.StatusNotImplemented, nil))
}

func createObject(w http.ResponseWriter, r *http.Request) {
	render.Render(w, r, errorRenderer(fmt.Errorf("TODO: implement createObject handler"), http.StatusNotImplemented, nil))
}

func deleteObject(w http.ResponseWriter, r *http.Request) {
	render.Render(w, r, errorRenderer(fmt.Errorf("TODO: implement deleteObject handler"), http.StatusNotImplemented, nil))
}
