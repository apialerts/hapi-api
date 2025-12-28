package responses

// ErrorResponse defines the standard JSON payload for an API error.
type ErrorResponse struct {
	Error string `json:"error"`
}
