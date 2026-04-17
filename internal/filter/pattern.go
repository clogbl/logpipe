package filter

import (
	"fmt"
	"regexp"
)

// Pattern holds a compiled regex and optional label for log line matching.
type Pattern struct {
	Label  string
	Regexp *regexp.Regexp
}

// NewPattern compiles a regex pattern and returns a Pattern.
func NewPattern(label, expr string) (*Pattern, error) {
	re, err := regexp.Compile(expr)
	if err != nil {
		return nil, fmt.Errorf("invalid pattern %q: %w", expr, err)
	}
	return &Pattern{Label: label, Regexp: re}, nil
}

// Match reports whether the line matches the pattern.
func (p *Pattern) Match(line string) bool {
	return p.Regexp.MatchString(line)
}

// Filter holds a list of patterns and a mode (include/exclude).
type Filter struct {
	Patterns []*Pattern
	Exclude  bool // if true, matching lines are dropped
}

// NewFilter creates a Filter from a slice of raw regex expressions.
func NewFilter(exprs []string, exclude bool) (*Filter, error) {
	patterns := make([]*Pattern, 0, len(exprs))
	for i, expr := range exprs {
		p, err := NewPattern(fmt.Sprintf("p%d", i), expr)
		if err != nil {
			return nil, err
		}
		patterns = append(patterns, p)
	}
	return &Filter{Patterns: patterns, Exclude: exclude}, nil
}

// Keep returns true if the line should be kept according to filter rules.
// With no patterns configured every line is kept.
func (f *Filter) Keep(line string) bool {
	if len(f.Patterns) == 0 {
		return true
	}
	for _, p := range f.Patterns {
		if p.Match(line) {
			return !f.Exclude
		}
	}
	return f.Exclude
}
