package web

import (
	"encoding/json"
	"net/http"

	"github.com/anonychun/go-blog-api/internal/app/model"
)

func MarshalPayload(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}

func MarshalError(w http.ResponseWriter, code int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(model.ErrorResponse{Message: err.Error()})
}
