package pipeline

import (
	"bytes"
	"strings"
	"testing"
)

func TestNewIndentStage_Valid(t *testing.T) {
	_, err := NewIndentStage("  ")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestNewIndentStage_Empty(t *testing.T) {
	_, err := NewIndentStage("")
	if err == nil {
		t.Fatal("expected error for empty indent, got nil")
	}
}

func TestNewSpaceIndentStage_Valid(t *testing.T) {
	_, err := NewSpaceIndentStage(4)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestNewSpaceIndentStage_Invalid(t *testing.T) {
	_, err := NewSpaceIndentStage(0)
	if err == nil {
		t.Fatal("expected error for zero spaces, got nil")
	}
}

func TestIndentStage_Output(t *testing.T) {
	stage, err := NewIndentStage(">> ")
	if err != nil {
		t.Fatalf("setup error: %v", err)
	}

	input := "foo\nbar\nbaz\n"
	reader := strings.NewReader(input)
	var buf bytes.Buffer

	if err := stage.Run(reader, &buf); err != nil {
		t.Fatalf("Run error: %v", err)
	}

	lines := strings.Split(strings.TrimRight(buf.String(), "\n"), "\n")
	expected := []string{">> foo", ">> bar", ">> baz"}
	if len(lines) != len(expected) {
		t.Fatalf("expected %d lines, got %d", len(expected), len(lines))
	}
	for i, want := range expected {
		if lines[i] != want {
			t.Errorf("line %d: expected %q, got %q", i, want, lines[i])
		}
	}
}

func TestSpaceIndentStage_Output(t *testing.T) {
	stage, err := NewSpaceIndentStage(3)
	if err != nil {
		t.Fatalf("setup error: %v", err)
	}

	input := "hello\nworld\n"
	reader := strings.NewReader(input)
	var buf bytes.Buffer

	if err := stage.Run(reader, &buf); err != nil {
		t.Fatalf("Run error: %v", err)
	}

	lines := strings.Split(strings.TrimRight(buf.String(), "\n"), "\n")
	for _, line := range lines {
		if !strings.HasPrefix(line, "   ") {
			t.Errorf("expected 3-space indent, got %q", line)
		}
	}
}
