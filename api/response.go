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
