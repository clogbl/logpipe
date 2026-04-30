package pipeline

import (
	"fmt"

	"github.com/user/logpipe/internal/transform"
)

// NewCountStage returns a Stage that counts every line passing through the
// pipeline. When appendCount is true each line is annotated with the running
// total in the form " [<n>]".
//
// The returned *transform.Counter is also exposed so callers can inspect or
// reset the count after the pipeline finishes.
func NewCountStage(appendCount bool) (Stage, *transform.Counter, error) {
	c, err := transform.NewCounter(appendCount)
	if err != nil {
		return nil, nil, fmt.Errorf("count stage: %w", err)
	}
	stage := func(line string, index int) (string, bool) {
		out, err := c.Format(line, index)
		if err != nil {
			return line, true
		}
		return out, true
	}
	return stage, c, nil
}
