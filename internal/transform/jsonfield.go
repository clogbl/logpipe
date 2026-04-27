package transform

import (
	"encoding/json"
	"fmt"
	"strings"
)

// JSONFieldExtractor extracts one or more fields from a JSON log line and
// formats them as a new line. This is useful for narrowing verbose structured
// log output to only the fields you care about.
//
// Example input:  {"level":"info","msg":"started","ts":"2024-01-01T00:00:00Z"}
// Fields ["level","msg"] → "info started"
type JSONFieldExtractor struct {
	fields    []string
	separator string
}

// JSONFieldOption configures a JSONFieldExtractor.
type JSONFieldOption func(*JSONFieldExtractor)

// WithJSONSeparator sets the string used to join extracted field values.
// Defaults to a single space.
func WithJSONSeparator(sep string) JSONFieldOption {
	return func(e *JSONFieldExtractor) {
		e.separator = sep
	}
}

// NewJSONFieldExtractor creates a JSONFieldExtractor that pulls the given
// top-level JSON keys from each line. At least one field must be provided.
// Returns an error if fields is empty.
func NewJSONFieldExtractor(fields []string, opts ...JSONFieldOption) (*JSONFieldExtractor, error) {
	if len(fields) == 0 {
		return nil, fmt.Errorf("jsonfield: at least one field name is required")
	}
	for i, f := range fields {
		if strings.TrimSpace(f) == "" {
			return nil, fmt.Errorf("jsonfield: field at index %d is empty or whitespace", i)
		}
	}
	e := &JSONFieldExtractor{
		fields:    fields,
		separator: " ",
	}
	for _, o := range opts {
		o(e)
	}
	return e, nil
}

// Format parses line as a JSON object and returns the requested field values
// joined by the configured separator. If the line is not valid JSON, or none
// of the requested fields are present, the original line is returned unchanged
// so that non-JSON lines pass through gracefully.
func (e *JSONFieldExtractor) Format(line string, _ int) (string, error) {
	var obj map[string]interface{}
	if err := json.Unmarshal([]byte(line), &obj); err != nil {
		// Not JSON — pass through unchanged.
		return line, nil
	}

	parts := make([]string, 0, len(e.fields))
	for _, f := range e.fields {
		val, ok := obj[f]
		if !ok {
			continue
		}
		parts = append(parts, fmt.Sprintf("%v", val))
	}

	if len(parts) == 0 {
		// None of the requested fields were found — pass through unchanged.
		return line, nil
	}

	return strings.Join(parts, e.separator), nil
}
