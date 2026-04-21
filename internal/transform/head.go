package transform

import "fmt"

// HeadLimiter passes through only the first N lines, dropping all subsequent lines.
type HeadLimiter struct {
	maxLines int
	seen     int
}

// NewHeadLimiter creates a HeadLimiter that allows at most maxLines lines through.
// Returns an error if maxLines is less than 1.
func NewHeadLimiter(maxLines int) (*HeadLimiter, error) {
	if maxLines < 1 {
		return nil, fmt.Errorf("head: maxLines must be at least 1, got %d", maxLines)
	}
	return &HeadLimiter{maxLines: maxLines}, nil
}

// Format returns the line unchanged if the limit has not been reached,
// or an empty string once the limit is exceeded.
func (h *HeadLimiter) Format(line string) string {
	if h.seen >= h.maxLines {
		return ""
	}
	h.seen++
	return line
}

// Reset resets the line counter, allowing lines to pass through again from the start.
func (h *HeadLimiter) Reset() {
	h.seen = 0
}
