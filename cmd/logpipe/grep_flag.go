package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/user/logpipe/internal/pipeline"
)

var (
	grepPattern string
	grepInvert  bool
)

func registerGrepFlags() {
	flag.StringVar(&grepPattern, "grep", "", "keep lines matching `pattern` (regex)")
	flag.BoolVar(&grepInvert, "grep-v", false, "invert grep: drop lines matching --grep pattern")
}

// applyGrep adds a grep stage to p when --grep is provided.
// It returns an error if the pattern is invalid.
func applyGrep(p *pipeline.Pipeline) error {
	if grepPattern == "" {
		if grepInvert {
			fmt.Fprintln(os.Stderr, "logpipe: --grep-v has no effect without --grep")
		}
		return nil
	}
	stage, err := pipeline.NewGrepStage(grepPattern, grepInvert)
	if err != nil {
		return fmt.Errorf("grep: %w", err)
	}
	p.AddStage(stage)
	return nil
}
