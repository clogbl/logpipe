package pipeline

import (
	"fmt"

	"github.com/yourorg/logpipe/internal/transform"
)

// NewMaskStage creates a pipeline stage that masks substrings matching the
// given regular expression pattern, replacing them with the provided mask
// string. This is useful for redacting sensitive data such as tokens,
// passwords, or PII from log streams before they are written to output.
//
// pattern must be a valid regular expression. If mask is empty, the
// transform package default mask ("***") is used.
func NewMaskStage(pattern, mask string) (*Stage, error) {
	if pattern == "" {
		return nil, fmt.Errorf("mask stage: pattern must not be empty")
	}

	m, err := transform.NewMasker(pattern, mask)
	if err != nil {
		return nil, fmt.Errorf("mask stage: %w", err)
	}

	return &Stage{formatter: m}, nil
}
