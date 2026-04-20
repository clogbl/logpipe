package transform

import (
	"strings"
	"testing"
)

func TestNewTruncator_Valid(t *testing.T) {
	tr, err := NewTruncator(80, "...")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if tr == nil {
		t.Fatal("expected non-nil Truncator")
	}
}

func TestNewTruncator_InvalidMaxLen(t *testing.T) {
	_, err := NewTruncator(0, "...")
	if err == nil {
		t.Fatal("expected error for maxLen=0")
	}
}

func TestNewTruncator_SuffixTooLong(t *testing.T) {
	_, err := NewTruncator(3, "...")
	if err == nil {
		t.Fatal("expected error when suffix length >= maxLen")
	}
}

func TestTruncator_ShortLine(t *testing.T) {
	tr, _ := NewTruncator(20, "...")
	line := "hello world"
	got := tr.Format(line)
	if got != line {
		t.Errorf("expected %q, got %q", line, got)
	}
}

func TestTruncator_ExactLength(t *testing.T) {
	tr, _ := NewTruncator(10, "...")
	line := "1234567890"
	got := tr.Format(line)
	if got != line {
		t.Errorf("expected %q, got %q", line, got)
	}
}

func TestTruncator_LongLine(t *testing.T) {
	tr, _ := NewTruncator(10, "...")
	line := "this line is definitely too long"
	got := tr.Format(line)
	if len(got) != 10 {
		t.Errorf("expected length 10, got %d", len(got))
	}
	if !strings.HasSuffix(got, "...") {
		t.Errorf("expected suffix '...', got %q", got)
	}
}

func TestTruncator_StripNewline(t *testing.T) {
	tr, _ := NewTruncator(20, "...")
	got := tr.Format("hello\n")
	if strings.Contains(got, "\n") {
		t.Errorf("expected newline stripped, got %q", got)
	}
}
