package handlers

import (
	"database/sql"
	"errors"
	"hapi/internal/contract/responses"
	"hapi/internal/util/webserver"
	"net/http"
)

func Health(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var value string

		err := db.QueryRow(`
			SELECT value
			FROM setting
			WHERE key = $1
		`, "health_status").Scan(&value)

		if errors.Is(err, sql.ErrNoRows) {
			webserver.RespondError(w, http.StatusNotFound, "Unknown server status")
			return
		}
		if err != nil {
			webserver.RespondError(w, http.StatusInternalServerError, "Failed to query data")
			return
		}

		response := responses.HealthResponse{
			Status: value,
		}
		webserver.RespondJSON(w, http.StatusOK, response)
	}
}
