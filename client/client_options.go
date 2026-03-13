package client

import "time"

// Option is a functional option for configuring a [Client].
type Option func(*Client)

// WithUserAgent returns an [Option] that sets the User-Agent header sent with
// every request.
func WithUserAgent(ua string) Option {
	return func(c *Client) {
		c.UserAgent = ua
	}
}

// WithBaseURL returns an [Option] that overrides the API base URL.
func WithBaseURL(url string) Option {
	return func(c *Client) {
		c.BaseURL = url
	}
}

// WithCache returns an [Option] that enables in-memory response caching with
// the given TTL. Cached responses are keyed by full URL and served without
// making a network request until the entry expires.
func WithCache(ttl time.Duration) Option {
	return func(c *Client) {
		c.cache = NewCache(ttl)
	}
}

// WithHTTPClient returns an [Option] that replaces the HTTP client used to
// execute requests. This is useful for injecting a custom transport or a test
// double that implements [HTTPClient].
func WithHTTPClient(httpClient HTTPClient) Option {
	return func(c *Client) {
		c.HTTP = httpClient
	}
}
