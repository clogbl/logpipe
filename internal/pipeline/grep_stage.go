package pipeline

import (
	"fmt"

	"github.com/user/logpipe/internal/transform"
)

// NewGrepStage creates a pipeline Stage that keeps only lines matching the
// given regular expression pattern. If invert is true, matching lines are
// dropped instead (equivalent to grep -v).
func NewGrepStage(pattern string, invert bool) (Stage, error) {
	if pattern == "" {
		return nil, fmt.Errorf("grep stage: pattern must not be empty")
	}
	g, err := transform.NewGrepper(pattern, invert)
	if err != nil {
		return nil, fmt.Errorf("grep stage: %w", err)
	}
	return formatterStage{formatter: g}, nil
}
