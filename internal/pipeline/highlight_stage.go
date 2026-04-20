package pipeline

import (
	"fmt"

	"github.com/user/logpipe/internal/transform"
)

// HighlightStage is a pipeline stage that wraps matched substrings with
// ANSI color codes before passing lines downstream.
type HighlightStage struct {
	highlighter *transform.Highlighter
}

// NewHighlightStage constructs a HighlightStage for the given pattern.
// opts are forwarded to transform.NewHighlighter.
func NewHighlightStage(pattern string, opts ...transform.HighlightOption) (*HighlightStage, error) {
	h, err := transform.NewHighlighter(pattern, opts...)
	if err != nil {
		return nil, fmt.Errorf("highlight stage: %w", err)
	}
	return &HighlightStage{highlighter: h}, nil
}

// Process applies highlighting to line and returns the result.
func (s *HighlightStage) Process(line string) (string, bool) {
	return s.highlighter.Format(line), true
}
