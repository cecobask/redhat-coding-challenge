package model

import (
	"net/http"
)

// Object ...
type Object struct {
	ID              string `json:"id"`
	ObjectName      string `json:"object_name,omitempty"`
	ObjectExtension string `json:"object_extension,omitempty"`
	ObjectPath      string `json:"object_path,omitempty"`
	BucketName      string `json:"bucket_name,omitempty"`
	CreatedAt       string `json:"created_at,omitempty"`
}

// ObjectList ...
type ObjectList struct {
	Objects []Object `json:"objects"`
}

// Render is an implementation of the Renderer interface https://pkg.go.dev/github.com/go-chi/render#Renderer
func (*ObjectList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// Render is an implementation of the Renderer interface https://pkg.go.dev/github.com/go-chi/render#Renderer
func (*Object) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
