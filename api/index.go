package api

import (
	"io"
	"net/http"
)

// Handler transforms html form post data into function calls.
func Handler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Write(body)
}
