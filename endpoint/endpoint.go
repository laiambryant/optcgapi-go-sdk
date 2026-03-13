package endpoint

import (
	"context"
	"encoding/json"

	"github.com/laiambryant/optcgapi-go-sdk/client"
)

// Endpoint is a generic HTTP endpoint that fetches a path and unmarshals the
// JSON response body into a value of type T.
type Endpoint[T any] struct {
	// Client is the underlying HTTP client used to execute requests.
	Client *client.Client
}

// New creates an Endpoint backed by c.
func New[T any](c *client.Client) *Endpoint[T] {
	return &Endpoint[T]{
		Client: c,
	}
}

// Fetch performs a GET request for path and unmarshals the response body into
// a value of type T. A [DecodeError] is returned if the body cannot be parsed
// as JSON; other errors propagate from [client.Client.Get].
func (e *Endpoint[T]) Fetch(ctx context.Context, path string) (T, error) {
	var item T
	data, err := e.Client.Get(ctx, path)
	if err != nil {
		return item, err
	}
	if err := json.Unmarshal(data, &item); err != nil {
		return item, &DecodeError{Resource: path, Err: err}
	}
	return item, nil
}
