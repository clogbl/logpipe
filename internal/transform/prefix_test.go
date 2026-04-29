package transform

import (
	"testing"
)

func TestNewLinePrefixer_Valid(t *testing.T) {
	p, err := NewLinePrefixer(">>> ")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if p == nil {
		t.Fatal("expected non-nil LinePrefixer")
	}
}

func TestNewLinePrefixer_Empty(t *testing.T) {
	_, err := NewLinePrefixer("")
	if err == nil {
		t.Fatal("expected error for empty prefix, got nil")
	}
}

func TestLinePrefixer_Format_Basic(t *testing.T) {
	p, err := NewLinePrefixer("[INFO] ")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	got, err := p.Format("hello world", 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != "[INFO] hello world" {
		t.Errorf("expected %q, got %q", "[INFO] hello world", got)
	}
}

func TestLinePrefixer_Format_EmptyLine(t *testing.T) {
	p, err := NewLinePrefixer(">> ")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	got, err := p.Format("", 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != ">> " {
		t.Errorf("expected %q, got %q", ">> ", got)
	}
}

func TestLinePrefixer_Format_IndexIgnored(t *testing.T) {
	p, err := NewLinePrefixer("- ")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	got0, _ := p.Format("line", 0)
	got5, _ := p.Format("line", 5)
	if got0 != got5 {
		t.Errorf("index should be ignored: got %q vs %q", got0, got5)
	}
}

func TestLinePrefixer_Format_SpecialChars(t *testing.T) {
	p, err := NewLinePrefixer("\t")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	got, err := p.Format("tabbed", 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != "\ttabbed" {
		t.Errorf("expected tab-prefixed line, got %q", got)
	}
}
