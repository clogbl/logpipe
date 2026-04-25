package pipeline

import (
	"fmt"

	"github.com/user/logpipe/internal/transform"
)

// NewUppercaseStage adds a stage that converts all log lines to uppercase.
// Returns an error if the pipeline itself is nil.
func NewUppercaseStage(p *Pipeline) error {
	if p == nil {
		return fmt.Errorf("uppercase stage: pipeline must not be nil")
	}
	u := transform.NewUppercaser(true)
	p.AddStage(u.Format)
	return nil
}

// NewLowercaseStage adds a stage that converts all log lines to lowercase.
// Returns an error if the pipeline itself is nil.
func NewLowercaseStage(p *Pipeline) error {
	if p == nil {
		return fmt.Errorf("lowercase stage: pipeline must not be nil")
	}
	u := transform.NewUppercaser(false)
	p.AddStage(u.Format)
	return nil
}
