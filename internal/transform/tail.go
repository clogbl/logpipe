package transform

import "fmt"

// TailLimiter buffers the last N lines and emits them once the stream ends.
// During streaming it acts as a pass-through gate: it always returns ErrSkip
// until Flush is called, which drains the ring buffer in order.
type TailLimiter struct {
	buf      []string
	maxLines int
	head     int // index of the oldest entry in the ring
	count    int // number of valid entries
}

// NewTailLimiter creates a TailLimiter that keeps the last maxLines lines.
func NewTailLimiter(maxLines int) (*TailLimiter, error) {
	if maxLines <= 0 {
		return nil, fmt.Errorf("tail: maxLines must be > 0, got %d", maxLines)
	}
	return &TailLimiter{
		buf:      make([]string, maxLines),
		maxLines: maxLines,
	}, nil
}

// Format records the line in the ring buffer and returns ErrSkip so the
// pipeline suppresses the line during normal streaming.
func (t *TailLimiter) Format(line string) (string, error) {
	slot := (t.head + t.count) % t.maxLines
	if t.count < t.maxLines {
		t.buf[slot] = line
		t.count++
	} else {
		// Ring is full — overwrite oldest entry and advance head.
		t.buf[t.head] = line
		t.head = (t.head + 1) % t.maxLines
	}
	return "", ErrSkip
}

// Flush returns all buffered lines in insertion order (oldest first).
func (t *TailLimiter) Flush() []string {
	out := make([]string, t.count)
	for i := 0; i < t.count; i++ {
		out[i] = t.buf[(t.head+i)%t.maxLines]
	}
	return out
}

// Reset clears the internal buffer so the limiter can be reused.
func (t *TailLimiter) Reset() {
	t.head = 0
	t.count = 0
}
