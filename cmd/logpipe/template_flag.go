package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/user/logpipe/internal/pipeline"
)

// runWithTemplate is invoked when the --template flag is provided.
// It bypasses the standard pipeline and uses TemplateStage directly.
func runWithTemplate(tmplStr string) {
	stage, err := pipeline.NewTemplateStage(tmplStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "logpipe: invalid template: %v\n", err)
		os.Exit(1)
	}
	if err := stage.Run(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "logpipe: template stage error: %v\n", err)
		os.Exit(1)
	}
}

// templateFlag returns the value of the --template flag if present.
func templateFlag() string {
	fs := flag.NewFlagSet("logpipe-template", flag.ContinueOnError)
	tmpl := fs.String("template", "", "Go template for formatting each log line")
	// parse only known flags; ignore errors from unknown flags
	_ = fs.Parse(os.Args[1:])
	return *tmpl
}
