package api_errors

import (
	"net/http"
	"strings"
)

const (
	// InvalidMissingCallerIDMessage is the default message when the caller.id is not present in the request.
	InvalidMissingCallerIDMessage = "Invalid or missing caller.id."
	// InvalidMissingCallerScopeMessage is the default message when the caller.scope is not present in the request.
	InvalidMissingCallerScopeMessage = "Invalid or missing caller.scopes."
	// BadRequestMessage is the default message when the input parameters on a request are wrong or it is malformed.
	BadRequestMessage = "Invalid request parameters."
	// ResourceNotFoundMessage is the default message when a requested resource is not available.
	ResourceNotFoundMessage = "Resource not found."
	// ResourceNotOwnedMessage is the default message when a user is requesting for a resource that he doesn't own.
	ResourceNotOwnedMessage = "You are not allowed to access this resource."
	// MethodNotAllowedMessage is the default message when a HTTP verb is forbidden on a resource.
	MethodNotAllowedMessage = "Method not allowed on the current resource."
	// StatusUnavailableForLegalReasonsMessage is the default message when a request over a LockDown user occurs.
	StatusUnavailableForLegalReasonsMessage = "The requested user is not available due to legal reasons"
	// InternalServerErrorMessage is the default message when an unexpected condition occurs.
	InternalServerErrorMessage = "Internal Server Error."
)

type APIError struct {
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Err     string   `json:"error"`
	Cause   []string `json:"cause"`
}

func newAPIError(code int, message string, err string) *APIError {
	return &APIError{
		Status:  code,
		Message: message,
		Err:     err,
		Cause:   make([]string, 0),
	}
}

func NewBadRequestError(messages ...string) *APIError {
	message := BadRequestMessage
	if len(messages) > 0 {
		message = strings.Join(messages, " - ")
	}
	return newAPIError(http.StatusBadRequest, message, "bad_request")
}

func NewUnauthorizedRequest(messages ...string) *APIError {
	message := InvalidMissingCallerIDMessage

	if len(messages) > 0 {
		message = strings.Join(messages, " - ")
	}
	return newAPIError(http.StatusUnauthorized, message, "unauthorized")
}

func NewInternalServerError(messages ...string) *APIError {
	message := InternalServerErrorMessage

	if len(messages) > 0 {
		message = strings.Join(messages, " - ")
	}
	return newAPIError(http.StatusInternalServerError, message, "internal_error")
}

func NewResourceNotFound(messages ...string) *APIError {
	message := ResourceNotFoundMessage

	if len(messages) > 0 {
		message = strings.Join(messages, " - ")
	}
	return newAPIError(http.StatusNotFound, message, "found")
}
