package pipeline

import (
	"fmt"

	"github.com/user/logpipe/internal/transform"
)

// NewSliceStage adds a field-slicing stage to the pipeline.
// It splits each line by delimiter and keeps fields [start:end].
// Use end = -1 to keep all fields from start to the last.
func NewSliceStage(p *Pipeline, delimiter string, start, end int) error {
	if p == nil {
		return fmt.Errorf("slice stage: pipeline must not be nil")
	}
	slicer, err := transform.NewSlicer(delimiter, start, end)
	if err != nil {
		return fmt.Errorf("slice stage: %w", err)
	}
	p.AddStage(slicer.Format)
	return nil
}
