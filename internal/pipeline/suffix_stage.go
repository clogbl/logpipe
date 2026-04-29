package pipeline

import (
	"fmt"

	"github.com/yourorg/logpipe/internal/transform"
)

// NewSuffixStage creates a pipeline Stage that appends suffix to every line.
// Returns an error if suffix is empty.
func NewSuffixStage(suffix string) (Stage, error) {
	s, err := transform.NewLineSuffixer(suffix)
	if err != nil {
		return nil, fmt.Errorf("suffix stage: %w", err)
	}
	return formatterStage{formatter: s}, nil
}
