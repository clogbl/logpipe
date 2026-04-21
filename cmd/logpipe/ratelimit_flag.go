package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/yourorg/logpipe/internal/pipeline"
)

// ratelimitFlag wires the --rate flag into the pipeline option set.
// It is called from main after all flags have been parsed.
func ratelimitFlag(opts *[]pipeline.Option) {
	rate := flag.Lookup("rate")
	if rate == nil {
		// flag not registered — nothing to do
		return
	}
	v := rate.Value.String()
	var lps float64
	if _, err := fmt.Sscanf(v, "%f", &lps); err != nil || lps <= 0 {
		return
	}
	*opts = append(*opts, pipeline.NewRateLimitStage(lps))
}

// registerRateLimitFlag registers the --rate CLI flag.
// Call this before flag.Parse().
func registerRateLimitFlag() *float64 {
	return flag.Float64("rate", 0,
		"max lines per second to forward (0 = unlimited)")
}

// applyRateLimit appends a rate-limit stage when linesPerSec > 0.
func applyRateLimit(linesPerSec float64, opts *[]pipeline.Option) {
	if linesPerSec <= 0 {
		return
	}
	*opts = append(*opts, pipeline.NewRateLimitStage(linesPerSec))
	fmt.Fprintf(os.Stderr, "[logpipe] rate limit: %.2f lines/sec\n", linesPerSec)
}
