package webhook

import (
	"net/http"
	"strconv"
	"sync"
	"time"
)

// rateLimiter handles Discord's rate limiting
type rateLimiter struct {
	mutex     sync.Mutex
	resetTime time.Time
	remaining int
}

// newRateLimiter creates a new rate limiter
func newRateLimiter() *rateLimiter {
	return &rateLimiter{}
}

// wait blocks until the rate limit allows a request
func (r *rateLimiter) wait() error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if r.remaining <= 0 && time.Now().Before(r.resetTime) {
		waitTime := time.Until(r.resetTime)
		return &RateLimitError{RetryAfter: waitTime}
	}

	return nil
}

// update updates the rate limiter based on Discord's response headers
func (r *rateLimiter) update(headers http.Header) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if remaining := headers.Get("X-RateLimit-Remaining"); remaining != "" {
		if val, err := strconv.Atoi(remaining); err == nil {
			r.remaining = val
		}
	}

	if reset := headers.Get("X-RateLimit-Reset"); reset != "" {
		if val, err := strconv.ParseInt(reset, 10, 64); err == nil {
			r.resetTime = time.Unix(val, 0)
		}
	}

	if retryAfter := headers.Get("Retry-After"); retryAfter != "" {
		if val, err := strconv.ParseInt(retryAfter, 10, 64); err == nil {
			r.resetTime = time.Now().Add(time.Duration(val) * time.Millisecond)
			r.remaining = 0
		}
	}
}
