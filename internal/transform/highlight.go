package transform

import (
	"fmt"
	"strings"
)

// Highlighter wraps matched substrings in ANSI escape codes.
type Highlighter struct {
	pattern string
	colorCode string
	reset     string
}

// HighlightOption configures a Highlighter.
type HighlightOption func(*Highlighter)

// WithColor sets the ANSI color code used for highlighting.
// code should be a valid SGR parameter, e.g. "31" for red.
func WithColor(code string) HighlightOption {
	return func(h *Highlighter) {
		h.colorCode = "\033[" + code + "m"
	}
}

// NewHighlighter creates a Highlighter that marks occurrences of pattern
// in each line with ANSI color escapes.
// Returns an error if pattern is empty.
func NewHighlighter(pattern string, opts ...HighlightOption) (*Highlighter, error) {
	if pattern == "" {
		return nil, fmt.Errorf("highlight: pattern must not be empty")
	}
	h := &Highlighter{
		pattern:   pattern,
		colorCode: "\033[33m", // default: yellow
		reset:     "\033[0m",
	}
	for _, o := range opts {
		o(h)
	}
	return h, nil
}

// Format replaces every occurrence of the pattern in line with a
// color-wrapped version and returns the result.
func (h *Highlighter) Format(line string) string {
	if !strings.Contains(line, h.pattern) {
		return line
	}
	replacement := h.colorCode + h.pattern + h.reset
	return strings.ReplaceAll(line, h.pattern, replacement)
}
