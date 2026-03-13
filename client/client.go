package client

import (
	"context"
	"errors"
	"io"
	"net/http"
)

// HTTPClient is the interface satisfied by any HTTP client that can be injected
// into [Client], including *http.Client and test doubles.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client is the HTTP client used by the SDK. BaseURL, HTTP, and UserAgent are
// publicly accessible so that callers can inspect the effective configuration,
// but they should normally be set through the functional options passed to
// [NewHTTPClient].
type Client struct {
	// BaseURL is prepended to every path passed to [Client.Get].
	BaseURL string
	// HTTP is the underlying HTTP client used to execute requests.
	HTTP HTTPClient
	// UserAgent is sent as the User-Agent header on every request.
	UserAgent string
	cache     *Cache
}

// NewHTTPClient creates a Client with the OPTCG API base URL and the
// optcgapi-go-sdk user-agent. If httpClient is nil, http.DefaultClient is
// used. Provide [Option] values to override any of these defaults.
func NewHTTPClient(httpClient HTTPClient, opts ...Option) *Client {
	c := &Client{
		BaseURL:   "https://www.optcgapi.com/api",
		HTTP:      httpClient,
		UserAgent: "optcgapi-go-sdk",
	}
	if c.HTTP == nil {
		c.HTTP = http.DefaultClient
	}

	for _, opt := range opts {
		opt(c)
	}
	return c
}

// Get performs a GET request for BaseURL+path. It returns the raw response
// body on success. Responses are served from the in-memory cache when caching
// is enabled and the entry has not expired.
//
// Errors:
//   - [ErrNotFound] for HTTP 404
//   - [*HTTPError] for any other non-2xx status
//   - [*RequestError] for network or I/O failures
func (c *Client) Get(ctx context.Context, path string) ([]byte, error) {
	fullURL := c.BaseURL + path
	if c.cache != nil {
		if data, ok := c.cache.Get(fullURL); ok {
			return data, nil
		}
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fullURL, nil)
	if err != nil {
		return nil, &RequestError{Op: "create request", Err: err}
	}
	req.Header.Set("User-Agent", c.UserAgent)

	resp, err := c.HTTP.Do(req)
	if err != nil {
		return nil, &RequestError{Op: "do request", Err: err}
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, &RequestError{Op: "read body", Err: err}
	}

	if resp.StatusCode == http.StatusNotFound {
		return nil, ErrNotFound
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, &HTTPError{
			Status: resp.StatusCode,
			URL:    fullURL,
			Body:   string(body),
			Cause:  errors.New("api error"),
		}
	}

	if c.cache != nil {
		c.cache.Set(fullURL, body)
	}

	return body, nil
}
