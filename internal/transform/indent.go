package transform

import (
	"fmt"
	"strings"
)

// Indenter prepends a fixed indentation string to each line.
type Indenter struct {
	prefix string
}

// NewIndenter creates an Indenter that prepends indent to every line.
// indent must be non-empty.
func NewIndenter(indent string) (*Indenter, error) {
	if indent == "" {
		return nil, fmt.Errorf("indent: indent string must not be empty")
	}
	return &Indenter{prefix: indent}, nil
}

// NewSpaceIndenter creates an Indenter using n spaces.
func NewSpaceIndenter(n int) (*Indenter, error) {
	if n <= 0 {
		return nil, fmt.Errorf("indent: number of spaces must be positive, got %d", n)
	}
	return NewIndenter(strings.Repeat(" ", n))
}

// Format prepends the indent prefix to line.
func (i *Indenter) Format(line string, _ int) (string, error) {
	return i.prefix + line, nil
}
