package pipeline

import (
	"bytes"
	"strings"
	"testing"
)

func TestTemplateStage_Basic(t *testing.T) {
	stage, err := NewTemplateStage("[{{.Index}}] {{.Line}}")
	if err != nil {
		t.Fatalf("setup error: %v", err)
	}
	input := strings.NewReader("foo\nbar\n")
	var out bytes.Buffer
	if err := stage.Run(input, &out); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	lines := strings.Split(strings.TrimRight(out.String(), "\n"), "\n")
	if len(lines) != 2 {
		t.Fatalf("expected 2 lines, got %d", len(lines))
	}
	if lines[0] != "[0] foo" {
		t.Errorf("line 0: expected '[0] foo', got %q", lines[0])
	}
	if lines[1] != "[1] bar" {
		t.Errorf("line 1: expected '[1] bar', got %q", lines[1])
	}
}

func TestTemplateStage_InvalidTemplate(t *testing.T) {
	_, err := NewTemplateStage("{{.Broken")
	if err == nil {
		t.Fatal("expected error for invalid template")
	}
}

func TestTemplateStage_EmptyInput(t *testing.T) {
	stage, _ := NewTemplateStage("{{.Line}}")
	var out bytes.Buffer
	if err := stage.Run(strings.NewReader(""), &out); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.Len() != 0 {
		t.Errorf("expected empty output, got %q", out.String())
	}
}
