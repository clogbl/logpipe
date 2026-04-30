package transform

import (
	"testing"
)

func TestNewReverser_Valid(t *testing.T) {
	r, err := NewReverser()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if r == nil {
		t.Fatal("expected non-nil Reverser")
	}
}

func TestReverser_EmptyLine(t *testing.T) {
	r, _ := NewReverser()
	out, err := r.Format("", 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "" {
		t.Errorf("expected empty string, got %q", out)
	}
}

func TestReverser_SimpleASCII(t *testing.T) {
	r, _ := NewReverser()
	out, err := r.Format("hello", 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "olleh" {
		t.Errorf("expected %q, got %q", "olleh", out)
	}
}

func TestReverser_SingleChar(t *testing.T) {
	r, _ := NewReverser()
	out, err := r.Format("x", 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "x" {
		t.Errorf("expected %q, got %q", "x", out)
	}
}

func TestReverser_Unicode(t *testing.T) {
	r, _ := NewReverser()
	out, err := r.Format("héllo", 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "olléh" {
		t.Errorf("expected %q, got %q", "olléh", out)
	}
}

func TestReverser_IndexIgnored(t *testing.T) {
	r, _ := NewReverser()
	out1, _ := r.Format("abcd", 0)
	out2, _ := r.Format("abcd", 99)
	if out1 != out2 {
		t.Errorf("index should be ignored: got %q and %q", out1, out2)
	}
}

func TestReverser_Palindrome(t *testing.T) {
	r, _ := NewReverser()
	out, err := r.Format("racecar", 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "racecar" {
		t.Errorf("expected palindrome unchanged, got %q", out)
	}
}
