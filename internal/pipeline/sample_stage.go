package pipeline

import (
	"fmt"

	"github.com/user/logpipe/internal/transform"
)

// NewSampleStage creates a pipeline stage that keeps every nth line.
// n must be >= 1.
func NewSampleStage(n int) (Stage, error) {
	if n < 1 {
		return nil, fmt.Errorf("sample: rate must be >= 1, got %d", n)
	}
	s, err := transform.NewSampler(n)
	if err != nil {
		return nil, fmt.Errorf("sample: %w", err)
	}
	return formatterStage{s}, nil
}
