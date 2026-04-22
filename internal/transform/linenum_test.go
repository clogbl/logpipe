package transform

import (
	"testing"
)

func TestNewLineNumberer_Valid(t *testing.T) {
	ln, err := NewLineNumberer(0, "")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if ln == nil {
		t.Fatal("expected non-nil LineNumberer")
	}
}

func TestNewLineNumberer_CustomStart(t *testing.T) {
	ln, err := NewLineNumberer(10, "%d")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	out, _ := ln.Format("hello")
	if out != "10\thello" {
		t.Errorf("expected '10\\thello', got %q", out)
	}
}

func TestNewLineNumberer_NegativeStart(t *testing.T) {
	_, err := NewLineNumberer(-1, "")
	if err == nil {
		t.Fatal("expected error for negative start")
	}
}

func TestNewLineNumberer_InvalidFormat(t *testing.T) {
	_, err := NewLineNumberer(0, "noVerb")
	if err == nil {
		t.Fatal("expected error for format string with no verb")
	}
}

func TestLineNumberer_Format_Increments(t *testing.T) {
	ln, _ := NewLineNumberer(1, "%d")
	lines := []string{"alpha", "beta", "gamma"}
	expected := []string{"1\talpha", "2\tbeta", "3\tgamma"}
	for i, line := range lines {
		out, err := ln.Format(line)
		if err != nil {
			t.Fatalf("unexpected error at line %d: %v", i, err)
		}
		if out != expected[i] {
			t.Errorf("line %d: expected %q, got %q", i, expected[i], out)
		}
	}
}

func TestLineNumberer_Format_PaddedFormat(t *testing.T) {
	ln, _ := NewLineNumberer(0, "%04d")
	out, err := ln.Format("msg")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "0000\tmsg" {
		t.Errorf("expected '0000\\tmsg', got %q", out)
	}
}

func TestLineNumberer_Reset(t *testing.T) {
	ln, _ := NewLineNumberer(5, "%d")
	ln.Format("a")
	ln.Format("b")
	ln.Reset()
	out, _ := ln.Format("c")
	if out != "5\tc" {
		t.Errorf("expected '5\\tc' after reset, got %q", out)
	}
}
