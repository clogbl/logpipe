package transform

import (
	"errors"
	"fmt"
)

// LineNumberer prepends a line number to each log line.
type LineNumberer struct {
	start  int
	counter int
	format string
}

// NewLineNumberer creates a LineNumberer starting at the given start index.
// start must be >= 0. format is a printf-style format string for the number
// (e.g. "%4d"); if empty, "%d" is used.
func NewLineNumberer(start int, format string) (*LineNumberer, error) {
	if start < 0 {
		return nil, errors.New("linenum: start index must be >= 0")
	}
	if format == "" {
		format = "%d"
	}
	// Validate the format string by doing a trial format.
	trial := fmt.Sprintf(format, 0)
	if trial == format {
		// No substitution happened — format string has no verb.
		return nil, errors.New("linenum: format string must contain a numeric verb (e.g. %d)")
	}
	return &LineNumberer{
		start:   start,
		counter: start,
		format:  format,
	}, nil
}

// Format prepends the current line number to line and increments the counter.
// It never returns ErrSkip.
func (l *LineNumberer) Format(line string) (string, error) {
	numStr := fmt.Sprintf(l.format, l.counter)
	l.counter++
	return numStr + "\t" + line, nil
}

// Reset resets the counter back to the start index.
func (l *LineNumberer) Reset() {
	l.counter = l.start
}
