package pipeline

import (
	"fmt"

	"github.com/yourorg/logpipe/internal/transform"
)

// NewRateLimitStage returns a pipeline Option that adds a rate-limiting stage.
// Lines that exceed linesPerSec are silently dropped.
// A non-positive linesPerSec value returns an error option that surfaces at
// pipeline construction time via the existing error-propagation pattern.
func NewRateLimitStage(linesPerSec float64) Option {
	return func(p *Pipeline) error {
		rl, err := transform.NewRateLimiter(linesPerSec)
		if err != nil {
			return fmt.Errorf("ratelimit stage: %w", err)
		}
		p.stages = append(p.stages, func(line string) string {
			return rl.Format(line)
		})
		return nil
	}
}
