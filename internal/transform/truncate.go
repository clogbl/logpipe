package transform

import (
	"fmt"
	"strings"
)

// Truncator truncates log lines that exceed a maximum length.
type Truncator struct {
	maxLen int
	suffix string
}

// NewTruncator creates a Truncator that clips lines to maxLen characters.
// A custom suffix (e.g. "...") is appended when a line is clipped.
// maxLen must be greater than the length of suffix.
func NewTruncator(maxLen int, suffix string) (*Truncator, error) {
	if maxLen <= 0 {
		return nil, fmt.Errorf("truncate: maxLen must be positive, got %d", maxLen)
	}
	if len(suffix) >= maxLen {
		return nil, fmt.Errorf("truncate: suffix length %d must be less than maxLen %d", len(suffix), maxLen)
	}
	return &Truncator{maxLen: maxLen, suffix: suffix}, nil
}

// Format returns the line unchanged if it fits within maxLen, otherwise
// truncates it and appends the suffix.
func (t *Truncator) Format(line string) string {
	line = strings.TrimRight(line, "\n")
	if len(line) <= t.maxLen {
		return line
	}
	cutAt := t.maxLen - len(t.suffix)
	return line[:cutAt] + t.suffix
}
