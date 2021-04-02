package http

import (
	"encoding/json"
	"net/http"
)

// ResponseMessage : wrapper struct for string response message
type ResponseMessage struct {
	Message string `json:"message"`
}

// WriteResponseJSON : helper method for writing response json
func respond(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// respondErr : default error response
func respondErr(w http.ResponseWriter, code int, err error) {
	errObj := struct {
		Error string `json:"error"`
	}{Error: err.Error()}
	respond(w, code, errObj)
}
