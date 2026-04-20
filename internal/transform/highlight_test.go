package transform

import (
	"strings"
	"testing"
)

func TestNewHighlighter_Valid(t *testing.T) {
	h, err := NewHighlighter("ERROR")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if h == nil {
		t.Fatal("expected non-nil Highlighter")
	}
}

func TestNewHighlighter_EmptyPattern(t *testing.T) {
	_, err := NewHighlighter("")
	if err == nil {
		t.Fatal("expected error for empty pattern")
	}
}

func TestHighlighter_NoMatch(t *testing.T) {
	h, _ := NewHighlighter("ERROR")
	line := "everything is fine"
	got := h.Format(line)
	if got != line {
		t.Errorf("expected unchanged line, got %q", got)
	}
}

func TestHighlighter_Match(t *testing.T) {
	h, _ := NewHighlighter("ERROR")
	line := "2024/01/01 ERROR something bad"
	got := h.Format(line)
	if !strings.Contains(got, "ERROR") {
		t.Error("expected ERROR to still appear in output")
	}
	// Should contain ANSI escape for yellow by default.
	if !strings.Contains(got, "\033[33m") {
		t.Error("expected ANSI yellow escape in output")
	}
	if !strings.Contains(got, "\033[0m") {
		t.Error("expected ANSI reset in output")
	}
}

func TestHighlighter_MultipleMatches(t *testing.T) {
	h, _ := NewHighlighter("WARN")
	line := "WARN: disk WARN: memory"
	got := h.Format(line)
	count := strings.Count(got, "\033[33m")
	if count != 2 {
		t.Errorf("expected 2 highlights, got %d", count)
	}
}

func TestHighlighter_WithColor(t *testing.T) {
	h, _ := NewHighlighter("CRIT", WithColor("31"))
	line := "CRIT: system failure"
	got := h.Format(line)
	if !strings.Contains(got, "\033[31m") {
		t.Errorf("expected red ANSI code, got %q", got)
	}
}
