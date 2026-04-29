package pipeline

import (
	"bytes"
	"strings"
	"testing"
)

func TestNewColumnStage_EmptyIndices(t *testing.T) {
	_, err := NewColumnStage([]int{}, "", "")
	if err == nil {
		t.Fatal("expected error for empty indices")
	}
}

func TestNewColumnStage_Valid(t *testing.T) {
	_, err := NewColumnStage([]int{0, 1}, "", "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestColumnStage_Output(t *testing.T) {
	stage, err := NewColumnStage([]int{0, 2}, "", " ")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	input := "alpha beta gamma\ndelta epsilon zeta\n"
	reader := strings.NewReader(input)
	var buf bytes.Buffer

	p, err := New(reader, &buf)
	if err != nil {
		t.Fatalf("pipeline creation failed: %v", err)
	}
	p.AddStage(stage)

	if err := p.Run(); err != nil {
		t.Fatalf("pipeline run failed: %v", err)
	}

	lines := strings.Split(strings.TrimRight(buf.String(), "\n"), "\n")
	if len(lines) != 2 {
		t.Fatalf("expected 2 lines, got %d", len(lines))
	}
	if lines[0] != "alpha gamma" {
		t.Errorf("line 0: expected 'alpha gamma', got %q", lines[0])
	}
	if lines[1] != "delta zeta" {
		t.Errorf("line 1: expected 'delta zeta', got %q", lines[1])
	}
}

func TestColumnStage_CustomSeparator(t *testing.T) {
	stage, err := NewColumnStage([]int{0, 2}, "|", "-")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	input := "a|b|c\n"
	reader := strings.NewReader(input)
	var buf bytes.Buffer

	p, err := New(reader, &buf)
	if err != nil {
		t.Fatalf("pipeline creation failed: %v", err)
	}
	p.AddStage(stage)

	if err := p.Run(); err != nil {
		t.Fatalf("pipeline run failed: %v", err)
	}

	got := strings.TrimRight(buf.String(), "\n")
	if got != "a-c" {
		t.Errorf("expected 'a-c', got %q", got)
	}
}
