package webhook

import (
	"fmt"
	"time"
)

// Error types for the webhook package
type (
	// RateLimitError is returned when Discord's rate limit is hit
	RateLimitError struct {
		RetryAfter time.Duration
	}

	// HTTPError is returned for non-2xx HTTP responses
	HTTPError struct {
		StatusCode int
		Message    string
	}
)

// Error returns the string representation of a RateLimitError
func (e RateLimitError) Error() string {
	return fmt.Sprintf("rate limit exceeded, retry after %v", e.RetryAfter)
}

// Error returns the string representation of an HTTPError
func (e HTTPError) Error() string {
	return fmt.Sprintf("HTTP error: %d - %s", e.StatusCode, e.Message)
}
