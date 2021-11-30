package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/cecobask/redhat-coding-challenge/db/files"
	"github.com/cecobask/redhat-coding-challenge/db/postgres"
	"github.com/cecobask/redhat-coding-challenge/model"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// ObjectsHandler ...
type ObjectsHandler interface {
	GetAllObjectsInBucket(http.ResponseWriter, *http.Request)
	GetObjectByBucketNameAndID(http.ResponseWriter, *http.Request)
	CreateOrUpdateObject(http.ResponseWriter, *http.Request)
	DeleteObjectByBucketNameAndID(http.ResponseWriter, *http.Request)
	InitObjectsRoute(chi.Router)
}

type objectsHandler struct {
	pg postgres.Database
	fm files.FileManager
}

type key int

const (
	keyBucket key = iota
	keyObjectID
)
const keyUploadObject = "uploadObject"

// NewObjectsHandler ...
func NewObjectsHandler(pg postgres.Database, fm files.FileManager) ObjectsHandler {
	return &objectsHandler{
		pg: pg,
		fm: fm,
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
			router.Delete("/", oh.DeleteObjectByBucketNameAndID)
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
	log.Println("objectsHandler.GetAllObjectsInBucket() invoked")
	bucketName := r.Context().Value(keyBucket).(string)
	objects, err := oh.pg.GetAllObjectsInBucket(r.Context(), bucketName)
	if err != nil {
		render.Render(w, r, errorRenderer(err, http.StatusInternalServerError, nil))
		return
	}
	if err := render.Render(w, r, objects); err != nil {
		render.Render(w, r, errorRenderer(err, http.StatusInternalServerError, nil))
		return
	}
	return
}

func (oh *objectsHandler) GetObjectByBucketNameAndID(w http.ResponseWriter, r *http.Request) {
	log.Println("objectsHandler.GetObjectByBucketNameAndID() invoked")
	bucketName := r.Context().Value(keyBucket).(string)
	objectID := r.Context().Value(keyObjectID).(string)
	object, err := oh.pg.GetObjectByBucketNameAndID(r.Context(), bucketName, objectID)
	if err != nil {
		if err == postgres.ErrNoRows {
			message := fmt.Sprintf("Object with ID %s not found in bucket %s", objectID, bucketName)
			render.Render(w, r, errorRenderer(postgres.ErrNoRows, http.StatusNotFound, &message))
		} else {
			render.Render(w, r, errorRenderer(err, http.StatusInternalServerError, nil))
		}
		return
	}
	path, err := oh.fm.RetrieveFile(*object)
	if err != nil {
		render.Render(w, r, errorRenderer(err, http.StatusInternalServerError, nil))
		return
	}
	http.ServeFile(w, r, path)
	return
}

func (oh *objectsHandler) CreateOrUpdateObject(w http.ResponseWriter, r *http.Request) {
	log.Println("objectsHandler.CreateOrUpdateObject() invoked")
	bucketName := r.Context().Value(keyBucket).(string)
	objectID := r.Context().Value(keyObjectID).(string)
	r.ParseMultipartForm(10 << 20)
	file, fileHeader, err := r.FormFile(keyUploadObject)
	if err != nil {
		message := fmt.Sprintf("No form key with value %s found in the request body", keyUploadObject)
		render.Render(w, r, errorRenderer(err, http.StatusBadRequest, &message))
		return
	}
	defer file.Close()
	lastDotIndex := strings.LastIndexByte(fileHeader.Filename, '.')
	objectName := fileHeader.Filename[:lastDotIndex]
	objectExtension := fileHeader.Filename[lastDotIndex+1:]
	objectModel := model.Object{
		ID:              objectID,
		ObjectName:      objectName,
		ObjectExtension: objectExtension,
		ObjectPath:      fmt.Sprintf("uploads/%s/%s.%s", bucketName, objectName, objectExtension),
		BucketName:      bucketName,
	}
	_, err = oh.pg.CreateOrUpdateObject(r.Context(), objectModel)
	if err != nil {
		render.Render(w, r, errorRenderer(err, http.StatusInternalServerError, nil))
		return
	}
	err = oh.fm.CreateFile(file, objectModel)
	if err != nil {
		render.Render(w, r, errorRenderer(err, http.StatusInternalServerError, nil))
		return
	}
	render.Status(r, http.StatusCreated)
	if err := render.Render(w, r, &model.Object{ID: objectID}); err != nil {
		render.Render(w, r, errorRenderer(err, http.StatusInternalServerError, nil))
	}
	return
}

func (oh *objectsHandler) DeleteObjectByBucketNameAndID(w http.ResponseWriter, r *http.Request) {
	log.Println("objectsHandler.DeleteObjectByBucketNameAndID() invoked")
	bucketName := r.Context().Value(keyBucket).(string)
	objectID := r.Context().Value(keyObjectID).(string)
	object, err := oh.pg.DeleteObjectByBucketNameAndID(r.Context(), bucketName, objectID)
	if err != nil {
		if err == postgres.ErrNoRows {
			message := fmt.Sprintf("Object with ID %s not found in bucket %s", objectID, bucketName)
			render.Render(w, r, errorRenderer(postgres.ErrNoRows, http.StatusNotFound, &message))
		} else {
			render.Render(w, r, errorRenderer(err, http.StatusInternalServerError, nil))
		}
		return
	}
	err = oh.fm.DeleteFile(*object)
	if err != nil {
		render.Render(w, r, errorRenderer(err, http.StatusInternalServerError, nil))
		return
	}
	if err := render.Render(w, r, object); err != nil {
		render.Render(w, r, errorRenderer(err, http.StatusInternalServerError, nil))
	}
	return
}
