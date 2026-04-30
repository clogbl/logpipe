package transform

import (
	"fmt"
	"sync/atomic"
)

// Sampler passes through every Nth line and drops the rest.
type Sampler struct {
	n       int
	counter atomic.Int64
}

// NewSampler creates a Sampler that keeps every nth line (1-based).
// n must be >= 1.
func NewSampler(n int) (*Sampler, error) {
	if n < 1 {
		return nil, fmt.Errorf("sample rate must be >= 1, got %d", n)
	}
	return &Sampler{n: n}, nil
}

// Format returns the line if it falls on a multiple of n, otherwise ErrSkip.
func (s *Sampler) Format(line string, _ int) (string, error) {
	count := s.counter.Add(1)
	if count%int64(s.n) == 0 {
		return line, nil
	}
	return "", ErrSkip
}
