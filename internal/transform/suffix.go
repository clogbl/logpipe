package transform

import "fmt"

// LineSuffixer appends a fixed string to every line.
type LineSuffixer struct {
	suffix string
}

// NewLineSuffixer returns a LineSuffixer that appends suffix to each line.
// Returns an error if suffix is empty.
func NewLineSuffixer(suffix string) (*LineSuffixer, error) {
	if suffix == "" {
		return nil, fmt.Errorf("suffix must not be empty")
	}
	return &LineSuffixer{suffix: suffix}, nil
}

// Format appends the configured suffix to line. The index parameter is ignored.
func (s *LineSuffixer) Format(line string, _ int) (string, error) {
	return line + s.suffix, nil
}
