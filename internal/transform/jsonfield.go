package transform

import (
	"encoding/json"
	"fmt"
	"strings"
)

// JSONFieldExtractor extracts one or more fields from a JSON log line.
type JSONFieldExtractor struct {
	fields    []string
	separator string
}

// JSONFieldOption configures a JSONFieldExtractor.
type JSONFieldOption func(*JSONFieldExtractor)

// WithJSONSeparator sets the separator used between extracted field values.
func WithJSONSeparator(sep string) JSONFieldOption {
	return func(e *JSONFieldExtractor) {
		e.separator = sep
	}
}

// NewJSONFieldExtractor creates a JSONFieldExtractor that extracts the given
// top-level fields from each JSON line. Fields are joined with separator
// (default: " "). Returns an error if no fields are provided.
func NewJSONFieldExtractor(fields []string, opts ...JSONFieldOption) (*JSONFieldExtractor, error) {
	if len(fields) == 0 {
		return nil, fmt.Errorf("jsonfield: at least one field name is required")
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

// Format extracts the configured fields from the JSON line and returns them
// joined by the separator. If the line is not valid JSON or a field is missing
// the raw line is returned unchanged.
func (e *JSONFieldExtractor) Format(line string, _ int) (string, error) {
	var obj map[string]interface{}
	if err := json.Unmarshal([]byte(line), &obj); err != nil {
		return line, nil
	}
	parts := make([]string, 0, len(e.fields))
	for _, f := range e.fields {
		v, ok := obj[f]
		if !ok {
			return line, nil
		}
		parts = append(parts, fmt.Sprintf("%v", v))
	}
	return strings.Join(parts, e.separator), nil
}
