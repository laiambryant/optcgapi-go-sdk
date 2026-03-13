package endpoint

import "fmt"

// DecodeError is returned by [Endpoint.Fetch] when the API response body
// cannot be unmarshalled into the target type.
type DecodeError struct {
	// Resource is the request path that produced the undecoded response.
	Resource string
	// Err is the underlying JSON decode error.
	Err error
}

// Error implements the error interface.
func (e *DecodeError) Error() string {
	return fmt.Sprintf("decode error for %s: %v", e.Resource, e.Err)
}

// Unwrap returns the underlying error, enabling errors.Is and errors.As.
func (e *DecodeError) Unwrap() error {
	return e.Err
}
