package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"hapi/internal/contract/requests"
	"hapi/internal/util/webserver"
	"net/http"
	"time"
)

func Import(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := r.PathValue("code")
		if code == "" {
			webserver.RespondError(w, http.StatusBadRequest, "Code is missing")
			return
		}

		ctx := r.Context()
		tx, err := db.BeginTx(ctx, nil)
		if err != nil {
			webserver.RespondError(w, http.StatusInternalServerError, "Failed to start transaction")
			return
		}
		// Defer a rollback in case anything fails. It's a no-op if Commit() was called.
		defer tx.Rollback()

		var id int
		var payload []byte
		var expiresAt time.Time

		// Query for the record and lock it for the transaction.
		err = tx.QueryRowContext(ctx, `
			SELECT id, expires, data
			FROM endpoint_export
			WHERE code = $1
			FOR UPDATE
		`, code).Scan(&id, &expiresAt, &payload)

		if errors.Is(err, sql.ErrNoRows) {
			webserver.RespondError(w, http.StatusNotFound, "Export is either expired, been used or does not exist")
			return
		}
		if err != nil {
			webserver.RespondError(w, http.StatusInternalServerError, "Failed to query data")
			return
		}

		// Always delete the record, whether it's expired or not.
		_, err = tx.ExecContext(ctx, `
			DELETE FROM endpoint_export
			WHERE id = $1
		`, id)
		if err != nil {
			webserver.RespondError(w, http.StatusInternalServerError, "Failed to delete record")
		}

		// Now, check if the record we just fetched was expired.
		if time.Now().UTC().After(expiresAt) {
			_ = tx.Commit() // Commit the deletion
			webserver.RespondError(w, http.StatusNotFound, "Export is either expired, been used or does not exist")
			return
		}

		// If we get here, the record was valid. Commit the transaction.
		if err := tx.Commit(); err != nil {
			webserver.RespondError(w, http.StatusInternalServerError, "Failed to commit transaction")
			return
		}

		var data requests.ExportRequest
		if err := json.Unmarshal(payload, &data); err != nil {
			// Handle the error, the JSON was malformed
			webserver.RespondError(w, http.StatusInternalServerError, "Failed to parse export data type, try a new export")
			return
		}

		// Return the record's data
		webserver.RespondJSON(w, http.StatusOK, data)
	}
}
