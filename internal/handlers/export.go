package handlers

import (
	"database/sql"
	"hapi/internal/contract/requests"
	"hapi/internal/contract/responses"
	"hapi/internal/util/webserver"
	"net/http"
	"time"

	"hapi/internal/util"
)

// Export Receives endpoints from the app and temporarily saves them to the database.
// Returns an import code that can be used on another device
func Export(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req requests.ExportRequest
		if err := webserver.DecodeJSON(r, &req); err != nil {
			webserver.RespondError(w, http.StatusBadRequest, "Invalid request body")
			return
		}

		if len(req.Endpoints) == 0 {
			webserver.RespondError(w, http.StatusBadRequest, "No endpoints to export")
			return
		}

		if len(req.Endpoints) > 1000 {
			webserver.RespondError(w, http.StatusBadRequest, "Too many endpoints to export")
			return
		}

		code, err := util.GenerateImportCode()
		if err != nil {
			webserver.RespondError(w, http.StatusInternalServerError, "Import code failed to generate, please try again")
			return
		}

		expiresAt := time.Now().UTC().Add(24 * time.Hour)

		_, err = db.Exec(`
			INSERT INTO endpoint_export (expires, code, data)
			VALUES ($1, $2, $3)
		`, expiresAt, code, req)
		if err != nil {
			webserver.RespondError(w, http.StatusInternalServerError, "Failed to export, please try again")
			return
		}

		response := responses.ExportResponse{
			Code:      code,
			ExpiresAt: expiresAt.String(),
		}
		webserver.RespondJSON(w, http.StatusOK, response)
	}
}
