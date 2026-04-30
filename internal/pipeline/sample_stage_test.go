package pipeline_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/user/logpipe/internal/pipeline"
)

func TestNewSampleStage_InvalidRate(t *testing.T) {
	_, err := pipeline.NewSampleStage(0)
	if err == nil {
		t.Fatal("expected error for rate=0")
	}
}

func TestNewSampleStage_Valid(t *testing.T) {
	_, err := pipeline.NewSampleStage(2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestSampleStage_EveryOther(t *testing.T) {
	stage, _ := pipeline.NewSampleStage(2)
	input := "line1\nline2\nline3\nline4\n"
	var out bytes.Buffer
	p, _ := pipeline.New(strings.NewReader(input), &out)
	p.AddStage(stage)
	if err := p.Run(); err != nil {
		t.Fatalf("pipeline error: %v", err)
	}
	got := strings.TrimSpace(out.String())
	lines := strings.Split(got, "\n")
	if len(lines) != 2 {
		t.Fatalf("expected 2 lines, got %d: %v", len(lines), lines)
	}
	if lines[0] != "line2" || lines[1] != "line4" {
		t.Errorf("expected [line2 line4], got %v", lines)
	}
}

func TestSampleStage_RateOne_PassAll(t *testing.T) {
	stage, _ := pipeline.NewSampleStage(1)
	input := "a\nb\nc\n"
	var out bytes.Buffer
	p, _ := pipeline.New(strings.NewReader(input), &out)
	p.AddStage(stage)
	if err := p.Run(); err != nil {
		t.Fatalf("pipeline error: %v", err)
	}
	got := strings.TrimSpace(out.String())
	if got != strings.TrimSpace(input) {
		t.Errorf("expected all lines, got %q", got)
	}
}
