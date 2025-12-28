package webserver

import (
	"encoding/json"
	"net/http"
)

func DecodeJSON(r *http.Request, v any) error {
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(v)
}
