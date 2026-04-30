package pipeline_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/user/logpipe/internal/pipeline"
)

func TestNewSqueezeStage_Whitespace(t *testing.T) {
	stage := pipeline.NewSqueezeStage()
	if stage == nil {
		t.Fatal("expected non-nil stage")
	}
}

func TestNewSqueezeCharStage_Valid(t *testing.T) {
	stage, err := pipeline.NewSqueezeCharStage('-')
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if stage == nil {
		t.Fatal("expected non-nil stage")
	}
}

func TestNewSqueezeCharStage_ZeroChar(t *testing.T) {
	_, err := pipeline.NewSqueezeCharStage(0)
	if err == nil {
		t.Fatal("expected error for zero rune")
	}
}

func TestSqueezeStage_Output(t *testing.T) {
	input := "hello   world\nfoo  bar\n"
	reader := strings.NewReader(input)
	var buf bytes.Buffer

	p, err := pipeline.New(reader, &buf)
	if err != nil {
		t.Fatalf("pipeline.New: %v", err)
	}
	p.AddStage(pipeline.NewSqueezeStage())
	if err := p.Run(); err != nil {
		t.Fatalf("Run: %v", err)
	}

	lines := strings.Split(strings.TrimRight(buf.String(), "\n"), "\n")
	expected := []string{"hello world", "foo bar"}
	if len(lines) != len(expected) {
		t.Fatalf("expected %d lines, got %d", len(expected), len(lines))
	}
	for i, want := range expected {
		if lines[i] != want {
			t.Errorf("line %d: expected %q, got %q", i, want, lines[i])
		}
	}
}

func TestSqueezeCharStage_Output(t *testing.T) {
	input := "foo---bar\nbaz--qux\n"
	reader := strings.NewReader(input)
	var buf bytes.Buffer

	p, err := pipeline.New(reader, &buf)
	if err != nil {
		t.Fatalf("pipeline.New: %v", err)
	}
	stage, err := pipeline.NewSqueezeCharStage('-')
	if err != nil {
		t.Fatalf("NewSqueezeCharStage: %v", err)
	}
	p.AddStage(stage)
	if err := p.Run(); err != nil {
		t.Fatalf("Run: %v", err)
	}

	lines := strings.Split(strings.TrimRight(buf.String(), "\n"), "\n")
	expected := []string{"foo-bar", "baz-qux"}
	if len(lines) != len(expected) {
		t.Fatalf("expected %d lines, got %d", len(expected), len(lines))
	}
	for i, want := range expected {
		if lines[i] != want {
			t.Errorf("line %d: expected %q, got %q", i, want, lines[i])
		}
	}
}
