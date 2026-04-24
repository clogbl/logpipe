package pipeline

import (
	"fmt"

	"github.com/user/logpipe/internal/transform"
)

// NewIndentStage creates a pipeline Stage that indents every line by the
// given indent string. Use NewSpaceIndentStage for space-based indentation.
func NewIndentStage(indent string) (Stage, error) {
	ind, err := transform.NewIndenter(indent)
	if err != nil {
		return nil, fmt.Errorf("indent stage: %w", err)
	}
	return newFormatterStage(ind), nil
}

// NewSpaceIndentStage creates a pipeline Stage that indents every line with
// n spaces.
func NewSpaceIndentStage(n int) (Stage, error) {
	ind, err := transform.NewSpaceIndenter(n)
	if err != nil {
		return nil, fmt.Errorf("space indent stage: %w", err)
	}
	return newFormatterStage(ind), nil
}
