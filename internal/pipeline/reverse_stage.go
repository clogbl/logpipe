package pipeline

import (
	"fmt"

	"github.com/user/logpipe/internal/transform"
)

// NewReverseStage creates a pipeline stage that reverses each log line's characters.
// It returns an error only for API consistency; reverser construction cannot fail.
func NewReverseStage(p *Pipeline) error {
	if p == nil {
		return fmt.Errorf("reverse stage: nil pipeline")
	}
	r, err := transform.NewReverser()
	if err != nil {
		return fmt.Errorf("reverse stage: %w", err)
	}
	p.AddStage(func(line string, index int) (string, error) {
		return r.Format(line, index)
	})
	return nil
}
