package transform

import (
	"fmt"
	"strings"
)

// Wrapper wraps long lines at a given column width, optionally indenting
// continuation lines.
type Wrapper struct {
	width  int
	indent string
}

// NewWrapper creates a Wrapper that hard-wraps lines at width columns.
// Continuation lines are prefixed with indent.
// width must be >= 1 and len(indent) must be less than width.
func NewWrapper(width int, indent string) (*Wrapper, error) {
	if width < 1 {
		return nil, fmt.Errorf("wrap: width must be >= 1, got %d", width)
	}
	if len(indent) >= width {
		return nil, fmt.Errorf("wrap: indent length %d must be less than width %d", len(indent), width)
	}
	return &Wrapper{width: width, indent: indent}, nil
}

// Format wraps line at w.width columns. The first segment uses the full width;
// subsequent segments are prefixed with w.indent. Returns the joined result.
func (w *Wrapper) Format(line string, _ int) (string, error) {
	if len(line) <= w.width {
		return line, nil
	}

	var segments []string
	remaining := line
	first := true

	for len(remaining) > 0 {
		var seg string
		if first {
			if len(remaining) <= w.width {
				seg = remaining
				remaining = ""
			} else {
				seg = remaining[:w.width]
				remaining = remaining[w.width:]
			}
			first = false
		} else {
			avail := w.width - len(w.indent)
			if len(remaining) <= avail {
				seg = w.indent + remaining
				remaining = ""
			} else {
				seg = w.indent + remaining[:avail]
				remaining = remaining[avail:]
			}
		}
		segments = append(segments, seg)
	}

	return strings.Join(segments, "\n"), nil
}
