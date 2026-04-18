package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/user/logpipe/internal/filter"
	"github.com/user/logpipe/internal/pipeline"
	"github.com/user/logpipe/internal/transform"
)

func main() {
	var (
		include    = flag.String("include", "", "regex pattern to include lines")
		exclude    = flag.String("exclude", "", "regex pattern to exclude lines")
		prefix     = flag.String("prefix", "", "prefix to prepend to each line")
		format     = flag.String("format", "raw", "output format: raw or json")
		timestamp  = flag.Bool("timestamp", false, "add timestamp to each line")
	)
	flag.Parse()

	var filterOpts []filter.Option
	if *include != "" {
		p, err := filter.NewPattern(*include)
		if err != nil {
			fmt.Fprintf(os.Stderr, "invalid include pattern: %v\n", err)
			os.Exit(1)
		}
		filterOpts = append(filterOpts, filter.WithInclude(p))
	}
	if *exclude != "" {
		p, err := filter.NewPattern(*exclude)
		if err != nil {
			fmt.Fprintf(os.Stderr, "invalid exclude pattern: %v\n", err)
			os.Exit(1)
		}
		filterOpts = append(filterOpts, filter.WithExclude(p))
	}

	f, err := filter.NewFilter(filterOpts...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create filter: %v\n", err)
		os.Exit(1)
	}

	var fmtOpts []transform.Option
	if *prefix != "" {
		fmtOpts = append(fmtOpts, transform.WithPrefix(*prefix))
	}
	if *timestamp {
		fmtOpts = append(fmtOpts, transform.WithTimestamp())
	}

	fmt_, err := transform.NewFormatter(strings.ToLower(*format), fmtOpts...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create formatter: %v\n", err)
		os.Exit(1)
	}

	p, err := pipeline.New(os.Stdin, os.Stdout, f, fmt_)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create pipeline: %v\n", err)
		os.Exit(1)
	}

	if err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "pipeline error: %v\n", err)
		os.Exit(1)
	}
}
