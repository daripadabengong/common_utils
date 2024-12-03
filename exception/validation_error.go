package exception

import (
	"encoding/json"
	"net/http"
)

// ValidationError represents an error for a specific field.
type ValidationError struct {
	Field   string `json:"field"`   // The name of the field that failed validation.
	Message string `json:"message"` // A descriptive error message for the field.
}

// InvalidRequestPayloadError represents multiple validation errors for a request payload.
type InvalidRequestPayloadError struct {
	Errors []*ValidationError `json:"errors"` // List of validation errors.
}

// NewInvalidRequestPayloadError creates a new instance of InvalidRequestPayloadError.
func NewInvalidRequestPayloadError() *InvalidRequestPayloadError {
	return &InvalidRequestPayloadError{
		Errors: []*ValidationError{},
	}
}

// AddError adds a new validation error to the InvalidRequestPayloadError.
func (e *InvalidRequestPayloadError) AddError(field, message string) {
	e.Errors = append(e.Errors, &ValidationError{
		Field:   field,
		Message: message,
	})
}

// Error implements the error interface.
func (e *InvalidRequestPayloadError) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// StatusCode returns the HTTP status code for this error.
func (e *InvalidRequestPayloadError) StatusCode() int {
	return http.StatusBadRequest
}

// HasErrors checks if there are any validation errors.
func (e *InvalidRequestPayloadError) HasErrors() bool {
	return len(e.Errors) > 0
}
