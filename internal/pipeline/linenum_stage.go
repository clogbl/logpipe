package pipeline

import (
	"fmt"

	"github.com/user/logpipe/internal/transform"
)

// NewLineNumStage adds a line-number stage to the pipeline.
// start is the first line number (>= 0), format is a printf-style numeric
// format string (e.g. "%4d"); pass an empty string to use the default "%d".
func NewLineNumStage(p *Pipeline, start int, format string) error {
	ln, err := transform.NewLineNumberer(start, format)
	if err != nil {
		return fmt.Errorf("linenum stage: %w", err)
	}
	p.AddStage(func(line string) (string, error) {
		return ln.Format(line)
	})
	return nil
}
