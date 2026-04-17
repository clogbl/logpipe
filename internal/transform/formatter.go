package transform

import (
	"fmt"
	"strings"
	"time"
)

// Format defines the output format for log lines.
type Format int

const (
	FormatRaw Format = iota
	FormatJSON
	FormatPrefix
)

// Formatter transforms a log line into a desired output format.
type Formatter struct {
	format Format
	prefix string
	timestamp bool
}

// NewFormatter creates a Formatter with the given format string.
// Supported formats: "raw", "json", "prefix".
func NewFormatter(format string, opts ...Option) (*Formatter, error) {
	f := &Formatter{}
	switch strings.ToLower(format) {
	case "raw":
		f.format = FormatRaw
	case "json":
		f.format = FormatJSON
	case "prefix":
		f.format = FormatPrefix
	default:
		return nil, fmt.Errorf("unknown format %q: must be raw, json, or prefix", format)
	}
	for _, o := range opts {
		o(f)
	}
	return f, nil
}

// Option configures a Formatter.
type Option func(*Formatter)

// WithPrefix sets a prefix string used by the prefix format.
func WithPrefix(p string) Option {
	return func(f *Formatter) { f.prefix = p }
}

// WithTimestamp enables prepending a UTC timestamp to each line.
func WithTimestamp(enabled bool) Option {
	return func(f *Formatter) { f.timestamp = enabled }
}

// Format transforms the given log line according to the formatter's settings.
func (f *Formatter) Format(line string) string {
	ts := ""
	if f.timestamp {
		ts = time.Now().UTC().Format(time.RFC3339) + " "
	}
	switch f.format {
	case FormatJSON:
		escaped := strings.ReplaceAll(line, `"`, `\"`)
		return fmt.Sprintf(`{"ts":%q,"msg":%q}`, strings.TrimRight(ts, " "), escaped)
	case FormatPrefix:
		return fmt.Sprintf("%s%s%s", ts, f.prefix, line)
	default:
		return ts + line
	}
}
