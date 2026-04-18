package pipeline

import (
	"bufio"
	"fmt"
	"io"

	"github.com/user/logpipe/internal/transform"
)

// TemplateStage applies a TemplateFormatter to each line in a pipeline.
type TemplateStage struct {
	formatter *transform.TemplateFormatter
}

// NewTemplateStage creates a TemplateStage from the given template string.
func NewTemplateStage(tmplStr string) (*TemplateStage, error) {
	f, err := transform.NewTemplateFormatter(tmplStr)
	if err != nil {
		return nil, err
	}
	return &TemplateStage{formatter: f}, nil
}

// Run reads lines from r, applies the template, and writes results to w.
func (ts *TemplateStage) Run(r io.Reader, w io.Writer) error {
	scanner := bufio.NewScanner(r)
	var index int64
	for scanner.Scan() {
		line := scanner.Text()
		out, err := ts.formatter.Format(line, index)
		if err != nil {
			return fmt.Errorf("template stage error at line %d: %w", index, err)
		}
		if _, err := fmt.Fprintln(w, out); err != nil {
			return err
		}
		index++
	}
	return scanner.Err()
}
