package pipeline

import (
	"fmt"

	"github.com/user/logpipe/internal/transform"
)

// NewSqueezeStage returns a Stage that collapses consecutive whitespace runs
// into a single space on each line.
func NewSqueezeStage() Stage {
	s := transform.NewWhitespaceSqueezer()
	return formatterStage{formatter: s}
}

// NewSqueezeCharStage returns a Stage that collapses consecutive occurrences
// of char into a single occurrence on each line.
func NewSqueezeCharStage(char rune) (Stage, error) {
	if char == 0 {
		return nil, fmt.Errorf("squeeze: char must be non-zero")
	}
	s, err := transform.NewSqueezer(char, false)
	if err != nil {
		return nil, err
	}
	return formatterStage{formatter: s}, nil
}
