package client

import (
	"errors"
	"fmt"
)

// ErrNotFound is returned by [Client.Get] when the API responds with HTTP 404.
var (
	ErrNotFound = errors.New("resource not found")
)

// HTTPError is returned when the API responds with a non-2xx, non-404 status
// code. It carries the status code, the request URL, and the raw response body
// to aid debugging.
type HTTPError struct {
	// Status is the HTTP response status code.
	Status int
	// URL is the full request URL that produced the error.
	URL string
	// Body is the raw response body returned by the server.
	Body string
	// Cause is the underlying sentinel error ("api error").
	Cause error
}

// Error implements the error interface.
func (e *HTTPError) Error() string {
	return fmt.Sprintf("http %d for %s: %s", e.Status, e.URL, e.Body)
}

// Unwrap returns the underlying cause, enabling errors.Is and errors.As.
func (e *HTTPError) Unwrap() error {
	return e.Cause
}

// RequestError is returned when a request cannot be created, executed, or when
// the response body cannot be read. Op describes which step failed.
type RequestError struct {
	// Op is a short description of the failed operation (e.g. "create request",
	// "do request", "read body").
	Op string
	// Err is the underlying error.
	Err error
}

// Error implements the error interface.
func (e *RequestError) Error() string {
	return fmt.Sprintf("request error during %s: %v", e.Op, e.Err)
}

// Unwrap returns the underlying error, enabling errors.Is and errors.As.
func (e *RequestError) Unwrap() error {
	return e.Err
}
