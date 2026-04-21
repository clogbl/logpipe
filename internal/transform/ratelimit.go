package transform

import (
	"errors"
	"time"
)

// RateLimiter drops lines that arrive faster than the configured rate.
// It uses a token-bucket approach: one token is consumed per line, and
// tokens refill at the given rate (lines per second).
type RateLimiter struct {
	tokens   float64
	max      float64
	refill   float64 // tokens per nanosecond
	lastTick time.Time
	now      func() time.Time
}

// NewRateLimiter creates a RateLimiter that allows at most linesPerSec lines
// per second with a burst capacity equal to linesPerSec.
func NewRateLimiter(linesPerSec float64) (*RateLimiter, error) {
	if linesPerSec <= 0 {
		return nil, errors.New("ratelimit: linesPerSec must be positive")
	}
	now := time.Now()
	return &RateLimiter{
		tokens:   linesPerSec,
		max:      linesPerSec,
		refill:   linesPerSec / float64(time.Second),
		lastTick: now,
		now:      time.Now,
	}, nil
}

// Allow returns true if the line should be forwarded, false if it should be
// dropped. It is safe to call from a single goroutine only.
func (r *RateLimiter) Allow() bool {
	now := r.now()
	elapsed := now.Sub(r.lastTick)
	r.lastTick = now

	r.tokens += float64(elapsed) * r.refill
	if r.tokens > r.max {
		r.tokens = r.max
	}

	if r.tokens >= 1.0 {
		r.tokens -= 1.0
		return true
	}
	return false
}

// Format satisfies a single-method transform interface used by pipeline stages.
// It returns the line unchanged when allowed, or an empty string when dropped.
func (r *RateLimiter) Format(line string) string {
	if r.Allow() {
		return line
	}
	return ""
}
