package pipeline

import (
	"fmt"

	"github.com/user/logpipe/internal/transform"
)

// NewTimestampStage creates a pipeline stage that prepends a timestamp
// to every log line using the provided Go time format string.
func NewTimestampStage(format string, opts ...transform.TimestampOption) (*Stage, error) {
	if format == "" {
		return nil, fmt.Errorf("timestamp stage: format must not be empty")
	}
	tp, err := transform.NewTimestampPrepender(format, opts...)
	if err != nil {
		return nil, fmt.Errorf("timestamp stage: %w", err)
	}
	return &Stage{formatter: tp}, nil
}
