package pipeline

import (
	"fmt"
	"strings"

	"github.com/user/logpipe/internal/transform"
)

// NewJSONFieldStage creates a pipeline Stage that extracts the given
// comma-separated field names from each JSON log line.
//
// Example:
//
//	stage, err := NewJSONFieldStage("level,msg", " | ")
func NewJSONFieldStage(fieldsCSV string, separator string) (Stage, error) {
	if strings.TrimSpace(fieldsCSV) == "" {
		return nil, fmt.Errorf("jsonfield stage: fields must not be empty")
	}
	fields := strings.Split(fieldsCSV, ",")
	for i, f := range fields {
		fields[i] = strings.TrimSpace(f)
	}

	opts := []transform.JSONFieldOption{}
	if separator != "" {
		opts = append(opts, transform.WithJSONSeparator(separator))
	}

	ext, err := transform.NewJSONFieldExtractor(fields, opts...)
	if err != nil {
		return nil, fmt.Errorf("jsonfield stage: %w", err)
	}
	return formatterStage{ext}, nil
}
