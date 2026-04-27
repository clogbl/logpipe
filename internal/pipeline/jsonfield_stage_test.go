package pipeline

import (
	"bytes"
	"strings"
	"testing"
)

func TestNewJSONFieldStage_EmptyFields(t *testing.T) {
	_, err := NewJSONFieldStage("", " ")
	if err == nil {
		t.Fatal("expected error for empty fields")
	}
}

func TestNewJSONFieldStage_Valid(t *testing.T) {
	s, err := NewJSONFieldStage("level,msg", " ")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if s == nil {
		t.Fatal("expected non-nil stage")
	}
}

func TestJSONFieldStage_Output(t *testing.T) {
	s, err := NewJSONFieldStage("level,msg", " | ")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	input := `{"level":"error","msg":"disk full","host":"srv1"}
{"level":"info","msg":"ok"}
`
	var out bytes.Buffer
	p, _ := New(strings.NewReader(input), &out)
	p.AddStage(s)
	if err := p.Run(); err != nil {
		t.Fatalf("pipeline error: %v", err)
	}

	lines := strings.Split(strings.TrimRight(out.String(), "\n"), "\n")
	if len(lines) != 2 {
		t.Fatalf("expected 2 lines, got %d: %v", len(lines), lines)
	}
	if lines[0] != "error | disk full" {
		t.Errorf("line 0: got %q, want %q", lines[0], "error | disk full")
	}
	if lines[1] != "info | ok" {
		t.Errorf("line 1: got %q, want %q", lines[1], "info | ok")
	}
}

func TestJSONFieldStage_InvalidJSONPassthrough(t *testing.T) {
	s, _ := NewJSONFieldStage("level", " ")
	input := "plain text line\n"
	var out bytes.Buffer
	p, _ := New(strings.NewReader(input), &out)
	p.AddStage(s)
	if err := p.Run(); err != nil {
		t.Fatalf("pipeline error: %v", err)
	}
	if strings.TrimRight(out.String(), "\n") != "plain text line" {
		t.Errorf("expected passthrough, got %q", out.String())
	}
}
