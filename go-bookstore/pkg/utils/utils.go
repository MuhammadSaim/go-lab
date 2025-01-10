package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

// a simple functio to parse the request body
func ParseBody(r *http.Request, x interface{}) {
	if body, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}
