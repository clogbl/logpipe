package transform

import (
	"fmt"
	"regexp"
)

// Grepper filters lines by matching against a regular expression,
// optionally inverting the match (like grep -v).
type Grepper struct {
	pattern *regexp.Regexp
	invert  bool
}

// NewGrepper creates a Grepper that keeps lines matching pattern.
// If invert is true, lines NOT matching the pattern are kept.
func NewGrepper(pattern string, invert bool) (*Grepper, error) {
	if pattern == "" {
		return nil, fmt.Errorf("grep pattern must not be empty")
	}
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, fmt.Errorf("invalid grep pattern %q: %w", pattern, err)
	}
	return &Grepper{pattern: re, invert: invert}, nil
}

// Format returns the line if it matches the configured pattern (or does not
// match when invert is true), otherwise it returns ErrSkip.
func (g *Grepper) Format(line string, index int) (string, error) {
	matched := g.pattern.MatchString(line)
	if g.invert {
		matched = !matched
	}
	if !matched {
		return "", ErrSkip
	}
	return line, nil
}
