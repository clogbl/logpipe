package transform

import "strings"

// Uppercaser transforms log lines to uppercase or lowercase.
type Uppercaser struct {
	upper bool
}

// NewUppercaser returns a transformer that converts lines to uppercase.
// If upper is false, lines are converted to lowercase instead.
func NewUppercaser(upper bool) *Uppercaser {
	return &Uppercaser{upper: upper}
}

// Format converts the given line to uppercase or lowercase.
func (u *Uppercaser) Format(line string, _ int) (string, error) {
	if u.upper {
		return strings.ToUpper(line), nil
	}
	return strings.ToLower(line), nil
}
