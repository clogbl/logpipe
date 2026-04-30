package transform

import (
	"fmt"
	"strings"
	"unicode"
)

// Squeezer collapses consecutive repeated characters (or whitespace) into one.
type Squeezer struct {
	whitespaceOnly bool
	target         rune
}

// NewSqueezer returns a Squeezer that collapses runs of the given character
// into a single occurrence. Pass 0 to squeeze all whitespace runs.
func NewSqueezer(char rune, whitespaceOnly bool) (*Squeezer, error) {
	if !whitespaceOnly && char == 0 {
		return nil, fmt.Errorf("squeeze: char must be non-zero when whitespaceOnly is false")
	}
	return &Squeezer{whitespaceOnly: whitespaceOnly, target: char}, nil
}

// NewWhitespaceSqueezer returns a Squeezer that collapses any run of
// whitespace characters into a single space.
func NewWhitespaceSqueezer() *Squeezer {
	return &Squeezer{whitespaceOnly: true}
}

// Format implements transform.Formatter.
func (s *Squeezer) Format(line string, _ int) (string, error) {
	if s.whitespaceOnly {
		return squeezeWhitespace(line), nil
	}
	return squeezeRune(line, s.target), nil
}

func squeezeWhitespace(line string) string {
	var b strings.Builder
	b.Grow(len(line))
	prevSpace := false
	for _, r := range line {
		if unicode.IsSpace(r) {
			if !prevSpace {
				b.WriteRune(' ')
			}
			prevSpace = true
		} else {
			b.WriteRune(r)
			prevSpace = false
		}
	}
	return b.String()
}

func squeezeRune(line string, target rune) string {
	var b strings.Builder
	b.Grow(len(line))
	prevMatch := false
	for _, r := range line {
		if r == target {
			if !prevMatch {
				b.WriteRune(r)
			}
			prevMatch = true
		} else {
			b.WriteRune(r)
			prevMatch = false
		}
	}
	return b.String()
}
