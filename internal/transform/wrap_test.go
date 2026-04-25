package transform

import (
	"strings"
	"testing"
)

func TestNewWrapper_Valid(t *testing.T) {
	w, err := NewWrapper(80, "  ")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if w == nil {
		t.Fatal("expected non-nil Wrapper")
	}
}

func TestNewWrapper_ZeroWidth(t *testing.T) {
	_, err := NewWrapper(0, "")
	if err == nil {
		t.Fatal("expected error for width=0")
	}
}

func TestNewWrapper_NegativeWidth(t *testing.T) {
	_, err := NewWrapper(-5, "")
	if err == nil {
		t.Fatal("expected error for negative width")
	}
}

func TestNewWrapper_IndentTooLong(t *testing.T) {
	_, err := NewWrapper(4, "    ") // indent == width
	if err == nil {
		t.Fatal("expected error when indent >= width")
	}
}

func TestWrapper_ShortLine(t *testing.T) {
	w, _ := NewWrapper(20, "  ")
	out, err := w.Format("hello", 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "hello" {
		t.Errorf("expected %q, got %q", "hello", out)
	}
}

func TestWrapper_ExactWidth(t *testing.T) {
	w, _ := NewWrapper(5, "")
	out, err := w.Format("hello", 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "hello" {
		t.Errorf("expected %q, got %q", "hello", out)
	}
}

func TestWrapper_WrapsLongLine(t *testing.T) {
	w, _ := NewWrapper(5, "")
	out, err := w.Format("abcdefghij", 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	lines := strings.Split(out, "\n")
	if len(lines) != 2 {
		t.Fatalf("expected 2 lines, got %d: %v", len(lines), lines)
	}
	if lines[0] != "abcde" || lines[1] != "fghij" {
		t.Errorf("unexpected wrap result: %v", lines)
	}
}

func TestWrapper_WithIndent(t *testing.T) {
	w, _ := NewWrapper(6, "  ")
	// first seg: 6 chars, continuation avail = 6-2 = 4
	out, err := w.Format("abcdefghij", 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	lines := strings.Split(out, "\n")
	if len(lines) != 2 {
		t.Fatalf("expected 2 lines, got %d: %q", len(lines), out)
	}
	if lines[0] != "abcdef" {
		t.Errorf("first line: expected %q, got %q", "abcdef", lines[0])
	}
	if lines[1] != "  ghij" {
		t.Errorf("second line: expected %q, got %q", "  ghij", lines[1])
	}
}

func TestWrapper_IndexIgnored(t *testing.T) {
	w, _ := NewWrapper(10, "")
	out1, _ := w.Format("hello world!", 0)
	out2, _ := w.Format("hello world!", 99)
	if out1 != out2 {
		t.Errorf("index should not affect output")
	}
}
