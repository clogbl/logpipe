package transform

import (
	"fmt"
	"strings"
)

// Slicer extracts a range of fields from each log line using a delimiter.
type Slicer struct {
	delimiter string
	start     int
	end       int // -1 means to the end
}

// NewSlicer creates a Slicer that splits each line by delimiter and returns
// fields [start:end]. Use end = -1 to slice to the last field inclusive.
func NewSlicer(delimiter string, start, end int) (*Slicer, error) {
	if delimiter == "" {
		return nil, fmt.Errorf("slicer: delimiter must not be empty")
	}
	if start < 0 {
		return nil, fmt.Errorf("slicer: start index must be >= 0, got %d", start)
	}
	if end != -1 && end < start {
		return nil, fmt.Errorf("slicer: end index (%d) must be >= start index (%d) or -1", end, start)
	}
	return &Slicer{delimiter: delimiter, start: start, end: end}, nil
}

// Format splits the line and returns the requested field slice joined by the
// same delimiter. Returns ErrSkip if the line has fewer fields than start+1.
func (s *Slicer) Format(line string, _ int) (string, error) {
	parts := strings.Split(line, s.delimiter)
	if s.start >= len(parts) {
		return "", ErrSkip
	}
	end := s.end
	if end == -1 || end >= len(parts) {
		end = len(parts) - 1
	}
	return strings.Join(parts[s.start:end+1], s.delimiter), nil
}
