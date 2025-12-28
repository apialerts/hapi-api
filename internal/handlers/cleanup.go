package handlers

import (
	"database/sql"
	"hapi/internal/util/webserver"
	"log"
	"net/http"
)

// CleanupExpired deletes all expired exports from the database
func CleanupExpired(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		result, err := db.ExecContext(ctx, `
			DELETE FROM endpoint_export
			WHERE expires < NOW()
		`)
		if err != nil {
			log.Printf("ERROR: Failed to delete expired records: %v", err)
			webserver.RespondError(w, http.StatusInternalServerError, "failed to execute cleanup task")
			return
		}

		rowsAffected, _ := result.RowsAffected()
		log.Printf("Cleanup successful: %d expired records deleted.", rowsAffected)

		// Send a success response back to the cron job.
		response := map[string]interface{}{
			"status":         "success",
			"recordsDeleted": rowsAffected,
		}
		webserver.RespondJSON(w, http.StatusOK, response)
	}
}
