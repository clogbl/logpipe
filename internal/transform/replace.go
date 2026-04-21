package transform

import (
	"fmt"
	"regexp"
	"strings"
)

// Replacer replaces occurrences of a pattern with a replacement string in each log line.
type Replacer struct {
	pattern     *regexp.Regexp
	replacement string
	literal     bool
}

// NewReplacer creates a new Replacer that substitutes matches of pattern with replacement.
// If literal is true, pattern is treated as a plain string rather than a regex.
func NewReplacer(pattern, replacement string, literal bool) (*Replacer, error) {
	if pattern == "" {
		return nil, fmt.Errorf("replace: pattern must not be empty")
	}

	var re *regexp.Regexp
	if !literal {
		var err error
		re, err = regexp.Compile(pattern)
		if err != nil {
			return nil, fmt.Errorf("replace: invalid pattern %q: %w", pattern, err)
		}
	} else {
		re = regexp.MustCompile(regexp.QuoteMeta(pattern))
	}

	return &Replacer{
		pattern:     re,
		replacement: replacement,
		literal:     literal,
	}, nil
}

// Format replaces all matches of the pattern in line with the replacement string.
// It always returns the (possibly modified) line and never returns ErrSkip.
func (r *Replacer) Format(line string, _ int) (string, error) {
	if !r.pattern.MatchString(line) {
		return line, nil
	}
	return r.pattern.ReplaceAllString(line, r.replacement), nil
}

// NewLiteralReplacer is a convenience constructor for plain-string replacement.
func NewLiteralReplacer(old, new string) (*Replacer, error) {
	if old == "" {
		return nil, fmt.Errorf("replace: old string must not be empty")
	}
	_ = strings.NewReplacer // ensure strings import is used for linting clarity
	return NewReplacer(old, new, true)
}
