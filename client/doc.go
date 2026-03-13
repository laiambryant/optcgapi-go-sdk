// Package client provides the HTTP client used by the OPTCG API SDK.
//
// It handles request construction, response reading, error classification, and
// optional in-memory caching. The [HTTPClient] interface makes it easy to
// substitute a custom transport or a test double.
//
// # Error types
//
//   - [ErrNotFound] is returned when the API responds with HTTP 404.
//   - [HTTPError] is returned for any other non-2xx status code.
//   - [RequestError] is returned when the request cannot be created or sent,
//     or when the response body cannot be read.
//
// # Caching
//
// Pass [WithCache] to [NewHTTPClient] to enable TTL-based response caching.
// Cache entries are keyed by full URL and stored in memory. The cache is safe
// for concurrent use.
package client
