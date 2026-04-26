package transform_test

import (
	"testing"

	"github.com/yourorg/logpipe/internal/transform"
)

func TestNewMasker_Valid(t *testing.T) {
	m, err := transform.NewMasker(`\d+`, "[NUM]")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if m == nil {
		t.Fatal("expected non-nil Masker")
	}
}

func TestNewMasker_EmptyPattern(t *testing.T) {
	_, err := transform.NewMasker("", "[MASKED]")
	if err == nil {
		t.Fatal("expected error for empty pattern")
	}
}

func TestNewMasker_InvalidRegex(t *testing.T) {
	_, err := transform.NewMasker(`[invalid`, "[MASKED]")
	if err == nil {
		t.Fatal("expected error for invalid regex")
	}
}

func TestNewMasker_DefaultMask(t *testing.T) {
	m, err := transform.NewMasker(`\d+`, "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	out, err := m.Format("user123", 0)
	if err != nil {
		t.Fatalf("unexpected format error: %v", err)
	}
	if out != "user***" {
		t.Errorf("expected 'user***', got %q", out)
	}
}

func TestMasker_NoMatch(t *testing.T) {
	m, _ := transform.NewMasker(`\d+`, "[NUM]")
	out, err := m.Format("no digits here", 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "no digits here" {
		t.Errorf("expected unchanged line, got %q", out)
	}
}

func TestMasker_SingleMatch(t *testing.T) {
	m, _ := transform.NewMasker(`\d+`, "[NUM]")
	out, err := m.Format("error code 404", 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "error code [NUM]" {
		t.Errorf("expected 'error code [NUM]', got %q", out)
	}
}

func TestMasker_MultipleMatches(t *testing.T) {
	m, _ := transform.NewMasker(`\d+`, "#")
	out, err := m.Format("ids: 1 2 3", 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "ids: # # #" {
		t.Errorf("expected 'ids: # # #', got %q", out)
	}
}

func TestMasker_SensitiveData(t *testing.T) {
	m, _ := transform.NewMasker(`password=\S+`, "password=[REDACTED]")
	out, err := m.Format("login password=secret123 ok", 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "login password=[REDACTED] ok" {
		t.Errorf("got %q", out)
	}
}
