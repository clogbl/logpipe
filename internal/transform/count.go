package transform

import (
	"fmt"
	"sync/atomic"
)

// Counter counts lines passing through and optionally appends the running
// total to each line in the form " [n]".
type Counter struct {
	count   atomic.Int64
	append  bool
}

// NewCounter creates a Counter. If appendCount is true, each formatted line
// will have " [<n>]" appended where <n> is the current running total.
func NewCounter(appendCount bool) (*Counter, error) {
	return &Counter{append: appendCount}, nil
}

// Format increments the internal counter and, if appendCount was set,
// returns the line with the running total appended. Otherwise it returns
// the line unchanged (the side-effect of counting still occurs).
func (c *Counter) Format(line string, _ int) (string, error) {
	n := c.count.Add(1)
	if c.append {
		return fmt.Sprintf("%s [%d]", line, n), nil
	}
	return line, nil
}

// Count returns the number of lines that have passed through so far.
func (c *Counter) Count() int64 {
	return c.count.Load()
}

// Reset resets the counter to zero.
func (c *Counter) Reset() {
	c.count.Store(0)
}
