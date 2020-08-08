package response

import (
	"encoding/json"
	"net/http"
)

func WithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Context-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func WithError(w http.ResponseWriter, code int, message string) {
	WithJSON(w, code, map[string]string{"error": message})
}
