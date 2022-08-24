package controller

import (
	"encoding/json"
	"net/http"
)

// ErrorResponse represents error response data object
type ErrorResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// ErrorResponse represents success response data object
type SuccessResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// RespondWithError calls RespondWithJSON and writes error response data object to response writer
func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, ErrorResponse{
		// We hardcoded this error value
		Code: 1,
		Msg:  message,
	})
}

// RespondWithJSON write json response data object to response writer
func RespondWithJSON(w http.ResponseWriter, code int, payload any) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
