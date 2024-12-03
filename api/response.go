package api

import (
	"encoding/json"
	"net/http"
)

type ErrorResponseBody struct {
	Message string `json:"message"`
}

type ErrorClientResponse struct {
	Code int
	ErrorResponseBody
}

func JSONResponse(w http.ResponseWriter, responsePayload interface{}, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if responsePayload != nil {
		json.NewEncoder(w).Encode(responsePayload)
	}
}

func JSONError(w http.ResponseWriter, errMessage string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	errResponseBody := ErrorResponseBody{
		Message: errMessage,
	}
	json.NewEncoder(w).Encode(errResponseBody)
}

func JSONObjectError(w http.ResponseWriter, message interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	// Serialize and write the response
	if err := json.NewEncoder(w).Encode(message); err != nil {
		// In case of an encoding failure, write a fallback error response
		http.Error(w, "Failed to encode error response", http.StatusInternalServerError)
	}
}
