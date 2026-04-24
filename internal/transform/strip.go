package transform

import (
	"fmt"
	"strings"
)

// Stripper removes a fixed prefix or suffix from each log line.
type Stripper struct {
	prefix string
	suffix string
}

// StripOption configures a Stripper.
type StripOption func(*Stripper)

// WithStripPrefix sets the prefix to remove.
func WithStripPrefix(prefix string) StripOption {
	return func(s *Stripper) {
		s.prefix = prefix
	}
}

// WithStripSuffix sets the suffix to remove.
func WithStripSuffix(suffix string) StripOption {
	return func(s *Stripper) {
		s.suffix = suffix
	}
}

// NewStripper creates a Stripper with at least one of prefix or suffix set.
func NewStripper(opts ...StripOption) (*Stripper, error) {
	s := &Stripper{}
	for _, o := range opts {
		o(s)
	}
	if s.prefix == "" && s.suffix == "" {
		return nil, fmt.Errorf("stripper: at least one of prefix or suffix must be non-empty")
	}
	return s, nil
}

// Format removes the configured prefix and/or suffix from line.
// It always returns FormatKeep so the (possibly modified) line passes through.
func (s *Stripper) Format(line string, _ int) (string, FormatAction) {
	if s.prefix != "" {
		line = strings.TrimPrefix(line, s.prefix)
	}
	if s.suffix != "" {
		line = strings.TrimSuffix(line, s.suffix)
	}
	return line, FormatKeep
}
