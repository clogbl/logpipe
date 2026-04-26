package transform

import (
	"fmt"
	"regexp"
)

// Masker replaces matched substrings with a fixed mask string.
type Masker struct {
	pattern *regexp.Regexp
	mask    string
}

// NewMasker returns a Masker that replaces all matches of pattern with mask.
// If mask is empty, it defaults to "***".
func NewMasker(pattern, mask string) (*Masker, error) {
	if pattern == "" {
		return nil, fmt.Errorf("mask: pattern must not be empty")
	}
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, fmt.Errorf("mask: invalid pattern: %w", err)
	}
	if mask == "" {
		mask = "***"
	}
	return &Masker{pattern: re, mask: mask}, nil
}

// Format replaces all occurrences of the pattern in line with the mask string.
// The index parameter is unused.
func (m *Masker) Format(line string, _ int) (string, error) {
	return m.pattern.ReplaceAllLiteralString(line, m.mask), nil
}
