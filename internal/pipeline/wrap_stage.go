package pipeline

import (
	"fmt"

	"github.com/user/logpipe/internal/transform"
)

// NewWrapStage creates a pipeline Stage that wraps long lines at width columns.
// Continuation lines are prefixed with indent.
func NewWrapStage(width int, indent string) (Stage, error) {
	if width < 1 {
		return nil, fmt.Errorf("wrap stage: width must be >= 1, got %d", width)
	}
	w, err := transform.NewWrapper(width, indent)
	if err != nil {
		return nil, fmt.Errorf("wrap stage: %w", err)
	}
	return stageFunc(w.Format), nil
}
