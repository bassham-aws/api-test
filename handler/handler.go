package handler

import (
	"encoding/json"
	"net/http"

	"github.com/rickbassham/goapi/middleware"
)

type ErrorResponse struct {
	Error string `json:"error,omitempty"`
}

type FooResponse struct {
	Message string `json:"message"`
	ErrorResponse
}

type HandlerService struct {
	log middleware.RequestLogger
}

func NewHandlerService(log middleware.RequestLogger) *HandlerService {
	return &HandlerService{
		log: log,
	}
}

func (svc *HandlerService) writeJSONResponse(w http.ResponseWriter, statusCode int, body interface{}) (err error) {
	w.WriteHeader(statusCode)
	enc := json.NewEncoder(w)
	err = enc.Encode(body)
	return
}

func (svc *HandlerService) Foo(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	svc.writeJSONResponse(w, http.StatusOK, &FooResponse{
		Message: "bar",
	})

	return
}
