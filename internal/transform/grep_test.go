package transform

import (
	"testing"
)

func TestNewGrepper_Valid(t *testing.T) {
	g, err := NewGrepper(`ERROR`, false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if g == nil {
		t.Fatal("expected non-nil Grepper")
	}
}

func TestNewGrepper_EmptyPattern(t *testing.T) {
	_, err := NewGrepper("", false)
	if err == nil {
		t.Fatal("expected error for empty pattern")
	}
}

func TestNewGrepper_InvalidPattern(t *testing.T) {
	_, err := NewGrepper(`[invalid`, false)
	if err == nil {
		t.Fatal("expected error for invalid regex")
	}
}

func TestGrepper_Match(t *testing.T) {
	g, _ := NewGrepper(`ERROR`, false)
	out, err := g.Format("ERROR: something failed", 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "ERROR: something failed" {
		t.Errorf("expected line passed through, got %q", out)
	}
}

func TestGrepper_NoMatch(t *testing.T) {
	g, _ := NewGrepper(`ERROR`, false)
	_, err := g.Format("INFO: all good", 0)
	if err != ErrSkip {
		t.Errorf("expected ErrSkip, got %v", err)
	}
}

func TestGrepper_Invert_Match(t *testing.T) {
	g, _ := NewGrepper(`DEBUG`, true)
	out, err := g.Format("INFO: something", 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "INFO: something" {
		t.Errorf("expected line passed through, got %q", out)
	}
}

func TestGrepper_Invert_NoMatch(t *testing.T) {
	g, _ := NewGrepper(`DEBUG`, true)
	_, err := g.Format("DEBUG: verbose output", 0)
	if err != ErrSkip {
		t.Errorf("expected ErrSkip for inverted match, got %v", err)
	}
}

func TestGrepper_CaseInsensitive(t *testing.T) {
	g, _ := NewGrepper(`(?i)error`, false)
	out, err := g.Format("Error: case test", 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out == "" {
		t.Error("expected non-empty output for case-insensitive match")
	}
}
