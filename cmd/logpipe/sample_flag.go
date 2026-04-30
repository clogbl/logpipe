package main

import (
	"flag"
	"fmt"

	"github.com/user/logpipe/internal/pipeline"
)

var sampleRate int

func registerSampleFlag(fs *flag.FlagSet) {
	fs.IntVar(&sampleRate, "sample", 0, "keep every Nth line (e.g. --sample 3); 0 disables")
}

func applySample(p *pipeline.Pipeline) error {
	if sampleRate == 0 {
		return nil
	}
	if sampleRate < 1 {
		return fmt.Errorf("--sample: rate must be >= 1, got %d", sampleRate)
	}
	stage, err := pipeline.NewSampleStage(sampleRate)
	if err != nil {
		return fmt.Errorf("--sample: %w", err)
	}
	p.AddStage(stage)
	return nil
}
