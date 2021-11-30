package handler

import (
	"fmt"
	"net/http"

	"github.com/go-chi/render"
)

type errorResponse struct {
	Err        string  `json:"err"`
	StatusCode int     `json:"status_code"`
	StatusText string  `json:"status_text"`
	Message    *string `json:"message,omitempty"`
}

// ErrUnknownFileType is returned when a file with unknown type is uploaded to the server
var ErrUnknownFileType = fmt.Errorf("Unknown file type")

// Render renders a single payload and respond to the client request
func (e *errorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.StatusCode)
	return nil
}

func errorRenderer(err error, statusCode int, message *string) *errorResponse {
	errorResponse := errorResponse{
		Err:     err.Error(),
		Message: message,
	}
	switch statusCode {
	case http.StatusInternalServerError:
		errorResponse.StatusCode = http.StatusInternalServerError
		errorResponse.StatusText = "Internal server error"
	case http.StatusNotImplemented:
		errorResponse.StatusCode = http.StatusNotImplemented
		errorResponse.StatusText = "Not implemented"
	case http.StatusNotFound:
		errorResponse.StatusCode = http.StatusNotFound
		errorResponse.StatusText = "Not found"
	case http.StatusBadRequest:
		errorResponse.StatusCode = http.StatusBadRequest
		errorResponse.StatusText = "Bad request"
	}
	return &errorResponse
}
