package transform

import (
	"errors"
	"fmt"
)

// LinePrefixer prepends a fixed string to every log line.
type LinePrefixer struct {
	prefix string
}

// NewLinePrefixer returns a LinePrefixer that prepends prefix to each line.
// Returns an error if prefix is empty.
func NewLinePrefixer(prefix string) (*LinePrefixer, error) {
	if prefix == "" {
		return nil, errors.New("prefix must not be empty")
	}
	return &LinePrefixer{prefix: prefix}, nil
}

// Format prepends the configured prefix to line.
// The index parameter is accepted for interface compatibility but unused.
func (p *LinePrefixer) Format(line string, _ int) (string, error) {
	return fmt.Sprintf("%s%s", p.prefix, line), nil
}
