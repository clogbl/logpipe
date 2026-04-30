package transform

import "errors"

// Reverser reverses the characters in each log line.
type Reverser struct{}

// NewReverser creates a new Reverser.
func NewReverser() (*Reverser, error) {
	return &Reverser{}, nil
}

// Format reverses the characters of the given line.
func (r *Reverser) Format(line string, _ int) (string, error) {
	if line == "" {
		return "", nil
	}
	runes := []rune(line)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes), nil
}

// newReverserError is returned when reverser construction fails.
var errReverserUnused = errors.New("reverser: unused")
