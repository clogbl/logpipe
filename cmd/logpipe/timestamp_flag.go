package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/user/logpipe/internal/pipeline"
)

const defaultTimestampFormat = time.RFC3339

var timestampFormat string

// registerTimestampFlag registers the --timestamp flag.
func registerTimestampFlag(fs *flag.FlagSet) {
	fs.StringVar(&timestampFormat, "timestamp", "",
		`prepend a timestamp to each line using the given Go time layout (e.g. "2006-01-02T15:04:05Z07:00"); use "default" for RFC3339`)
}

// applyTimestamp adds a timestamp stage to p when the flag is set.
func applyTimestamp(p *pipeline.Pipeline) error {
	if timestampFormat == "" {
		return nil
	}
	fmt_ := timestampFormat
	if fmt_ == "default" {
		fmt_ = defaultTimestampFormat
	}
	s, err := pipeline.NewTimestampStage(fmt_)
	if err != nil {
		return fmt.Errorf("--timestamp: %w", err)
	}
	p.AddStage(s)
	return nil
}
