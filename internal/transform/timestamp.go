package transform

import (
	"fmt"
	"time"
)

// TimestampPrepender prepends the current timestamp to each log line.
type TimestampPrepender struct {
	format string
	clock  func() time.Time
}

// TimestampOption configures a TimestampPrepender.
type TimestampOption func(*TimestampPrepender)

// WithClock overrides the clock used for timestamping (useful in tests).
func WithClock(fn func() time.Time) TimestampOption {
	return func(t *TimestampPrepender) {
		t.clock = fn
	}
}

// NewTimestampPrepender creates a TimestampPrepender with the given time format.
// The format must be a non-empty Go time layout string.
func NewTimestampPrepender(format string, opts ...TimestampOption) (*TimestampPrepender, error) {
	if format == "" {
		return nil, fmt.Errorf("timestamp format must not be empty")
	}
	tp := &TimestampPrepender{
		format: format,
		clock:  time.Now,
	}
	for _, o := range opts {
		o(tp)
	}
	return tp, nil
}

// Format prepends the current timestamp to line.
func (tp *TimestampPrepender) Format(line string, _ int) (string, error) {
	return fmt.Sprintf("%s %s", tp.clock().Format(tp.format), line), nil
}
