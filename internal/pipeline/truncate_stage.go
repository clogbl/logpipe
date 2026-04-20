package pipeline

import (
	"fmt"

	"github.com/user/logpipe/internal/transform"
)

// TruncateStage is a pipeline stage that truncates long lines.
type TruncateStage struct {
	truncator *transform.Truncator
}

// NewTruncateStage creates a pipeline stage that clips each line to maxLen
// characters, appending suffix when truncation occurs.
func NewTruncateStage(maxLen int, suffix string) (*TruncateStage, error) {
	tr, err := transform.NewTruncator(maxLen, suffix)
	if err != nil {
		return nil, fmt.Errorf("truncate stage: %w", err)
	}
	return &TruncateStage{truncator: tr}, nil
}

// Process applies truncation to every line emitted by the upstream channel
// and forwards results to the returned channel.
func (s *TruncateStage) Process(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for line := range in {
			out <- s.truncator.Format(line)
		}
	}()
	return out
}
