package bank

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Body    any    `json:"body,omitempty"`
}

func JsonResponse(w http.ResponseWriter, status int, message string, body ...any) {
	w.WriteHeader(status)
	if len(body) == 0 {
		body = nil
	}
	jsonResponse, err := json.Marshal(Response{status, message, body})
	if err != nil {
		fmt.Fprintf(w, "{\"status\":500, \"message\":\"Error interno del servidor\"}")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(jsonResponse)
	return
}
