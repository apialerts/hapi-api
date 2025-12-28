package webserver

import (
	"encoding/json"
	"hapi/internal/contract/responses"
	"log"
	"net/http"
)

func RespondJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

// RespondError sends a JSON-formatted error message with a given status code.
// If the message is an empty string, it only sends the status code with no body.
// It logs a server error if the JSON encoding fails.
func RespondError(w http.ResponseWriter, code int, message string) {
	if message == "" {
		w.WriteHeader(code)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	errPayload := responses.ErrorResponse{Error: message}
	if err := json.NewEncoder(w).Encode(errPayload); err != nil {
		log.Printf("RespondError: %v", err)
	}
}

func RespondHeader(w http.ResponseWriter, code int) {
	w.WriteHeader(code)
}

func RespondErrorBlank(w http.ResponseWriter, code int, issue string) {
	log.Printf("RespondErrorBlank: %v", issue)
	w.WriteHeader(code)
}
