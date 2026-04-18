package transform

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
)

// TemplateFormatter formats log lines using Go text/template syntax.
// Available fields: .Line, .Index
type TemplateFormatter struct {
	tmpl *template.Template
}

// TemplateData holds data passed to the template during execution.
type TemplateData struct {
	Line  string
	Index int64
}

// NewTemplateFormatter parses and validates the given template string.
func NewTemplateFormatter(tmplStr string) (*TemplateFormatter, error) {
	if strings.TrimSpace(tmplStr) == "" {
		return nil, fmt.Errorf("template string must not be empty")
	}
	tmpl, err := template.New("logpipe").Parse(tmplStr)
	if err != nil {
		return nil, fmt.Errorf("invalid template: %w", err)
	}
	return &TemplateFormatter{tmpl: tmpl}, nil
}

// Format applies the template to the given line and index.
func (t *TemplateFormatter) Format(line string, index int64) (string, error) {
	var buf bytes.Buffer
	data := TemplateData{Line: line, Index: index}
	if err := t.tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("template execution failed: %w", err)
	}
	return buf.String(), nil
}
