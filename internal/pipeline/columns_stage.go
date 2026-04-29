package pipeline

import (
	"fmt"

	"github.com/user/logpipe/internal/transform"
)

// NewColumnStage creates a pipeline stage that extracts specific whitespace-
// delimited columns from each log line. Columns are zero-indexed. An optional
// separator overrides the default whitespace splitting, and an optional joiner
// controls how selected columns are re-joined in the output.
//
// Example:
//
//	stage, err := NewColumnStage([]int{0, 2}, ",", " ")
func NewColumnStage(indices []int, separator, joiner string) (*Stage, error) {
	if len(indices) == 0 {
		return nil, fmt.Errorf("columns stage: at least one column index is required")
	}

	opts := []transform.ColumnOption{}
	if separator != "" {
		opts = append(opts, transform.WithColumnSeparator(separator))
	}
	if joiner != "" {
		opts = append(opts, transform.WithColumnJoiner(joiner))
	}

	ext, err := transform.NewColumnExtractor(indices, opts...)
	if err != nil {
		return nil, fmt.Errorf("columns stage: %w", err)
	}

	return &Stage{formatter: ext}, nil
}
