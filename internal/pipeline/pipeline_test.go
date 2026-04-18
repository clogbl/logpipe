package pipeline_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/logpipe/internal/filter"
	"github.com/logpipe/internal/pipeline"
	"github.com/logpipe/internal/transform"
)

func TestPipeline_PassThrough(t *testing.T) {
	input := "hello world\nfoo bar\n"
	reader := strings.NewReader(input)
	var writer bytes.Buffer

	p, err := pipeline.New(pipeline.Config{Reader: reader, Writer: &writer})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if err := p.Run(); err != nil {
		t.Fatalf("run error: %v", err)
	}
	if got := writer.String(); got != input {
		t.Errorf("expected %q, got %q", input, got)
	}
}

func TestPipeline_WithFilter(t *testing.T) {
	pat, _ := filter.NewPattern("error")
	f := filter.NewFilter(filter.WithInclude(pat))

	reader := strings.NewReader("info: ok\nerror: bad\nwarn: meh\n")
	var writer bytes.Buffer

	p, _ := pipeline.New(pipeline.Config{Filter: f, Reader: reader, Writer: &writer})
	p.Run()

	if got := writer.String(); got != "error: bad\n" {
		t.Errorf("unexpected output: %q", got)
	}
}

func TestPipeline_WithFormatter(t *testing.T) {
	fmt, _ := transform.NewFormatter(transform.WithPrefix("[LOG] "))
	reader := strings.NewReader("hello\n")
	var writer bytes.Buffer

	p, _ := pipeline.New(pipeline.Config{Formatter: fmt, Reader: reader, Writer: &writer})
	p.Run()

	if got := writer.String(); got != "[LOG] hello\n" {
		t.Errorf("unexpected output: %q", got)
	}
}

func TestPipeline_NilReader(t *testing.T) {
	_, err := pipeline.New(pipeline.Config{Writer: &bytes.Buffer{}})
	if err == nil {
		t.Error("expected error for nil reader")
	}
}

func TestPipeline_NilWriter(t *testing.T) {
	_, err := pipeline.New(pipeline.Config{Reader: strings.NewReader("")})
	if err == nil {
		t.Error("expected error for nil writer")
	}
}
