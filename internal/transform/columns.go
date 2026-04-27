// Package transform provides line-by-line transformation primitives for log streams.
package transform

import (
	"fmt"
	"strings"
)

// ColumnExtractor extracts specific whitespace-delimited columns from each log line.
// Columns are zero-indexed. If a requested column index exceeds the number of
// fields in a line, it is silently omitted from the output.
type ColumnExtractor struct {
	columns   []int
	separator string
	joiner    string
}

// ColumnOption configures a ColumnExtractor.
type ColumnOption func(*ColumnExtractor)

// WithColumnSeparator sets the field delimiter used when splitting lines.
// Defaults to whitespace splitting (strings.Fields behaviour) when empty.
func WithColumnSeparator(sep string) ColumnOption {
	return func(c *ColumnExtractor) {
		c.separator = sep
	}
}

// WithColumnJoiner sets the string used to join selected columns in the output.
// Defaults to a single space.
func WithColumnJoiner(join string) ColumnOption {
	return func(c *ColumnExtractor) {
		c.joiner = join
	}
}

// NewColumnExtractor creates a ColumnExtractor that selects the given column
// indices (zero-based) from each line. At least one column index must be
// provided, and no index may be negative.
func NewColumnExtractor(columns []int, opts ...ColumnOption) (*ColumnExtractor, error) {
	if len(columns) == 0 {
		return nil, fmt.Errorf("columns: at least one column index is required")
	}
	for _, c := range columns {
		if c < 0 {
			return nil, fmt.Errorf("columns: negative column index %d is not allowed", c)
		}
	}

	ext := &ColumnExtractor{
		columns: columns,
		joiner:  " ",
	}
	for _, o := range opts {
		o(ext)
	}
	return ext, nil
}

// Format implements the Formatter interface. It splits the line into fields,
// selects the configured columns, joins them, and returns the result.
// Lines that yield no selected columns are returned as an empty string rather
// than being dropped, so callers can decide how to handle them.
func (c *ColumnExtractor) Format(line string, _ int) (string, error) {
	var fields []string
	if c.separator == "" {
		fields = strings.Fields(line)
	} else {
		fields = strings.Split(line, c.separator)
	}

	selected := make([]string, 0, len(c.columns))
	for _, idx := range c.columns {
		if idx < len(fields) {
			selected = append(selected, fields[idx])
		}
	}

	return strings.Join(selected, c.joiner), nil
}
