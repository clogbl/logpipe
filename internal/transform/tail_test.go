package transform

import (
	"testing"
)

func TestNewTailLimiter_Valid(t *testing.T) {
	tl, err := NewTailLimiter(5)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if tl == nil {
		t.Fatal("expected non-nil TailLimiter")
	}
}

func TestNewTailLimiter_InvalidMaxLines(t *testing.T) {
	for _, n := range []int{0, -1, -100} {
		_, err := NewTailLimiter(n)
		if err == nil {
			t.Errorf("expected error for maxLines=%d", n)
		}
	}
}

func TestTailLimiter_FormatReturnsSkip(t *testing.T) {
	tl, _ := NewTailLimiter(3)
	_, err := tl.Format("line1")
	if err != ErrSkip {
		t.Errorf("expected ErrSkip, got %v", err)
	}
}

func TestTailLimiter_FlushFewerThanMax(t *testing.T) {
	tl, _ := NewTailLimiter(5)
	lines := []string{"a", "b", "c"}
	for _, l := range lines {
		tl.Format(l) //nolint:errcheck
	}
	got := tl.Flush()
	if len(got) != len(lines) {
		t.Fatalf("expected %d lines, got %d", len(lines), len(got))
	}
	for i, want := range lines {
		if got[i] != want {
			t.Errorf("line %d: want %q, got %q", i, want, got[i])
		}
	}
}

func TestTailLimiter_FlushExactMax(t *testing.T) {
	tl, _ := NewTailLimiter(3)
	for _, l := range []string{"x", "y", "z"} {
		tl.Format(l) //nolint:errcheck
	}
	got := tl.Flush()
	if len(got) != 3 {
		t.Fatalf("expected 3 lines, got %d", len(got))
	}
}

func TestTailLimiter_FlushOverflow(t *testing.T) {
	tl, _ := NewTailLimiter(3)
	for _, l := range []string{"1", "2", "3", "4", "5"} {
		tl.Format(l) //nolint:errcheck
	}
	got := tl.Flush()
	if len(got) != 3 {
		t.Fatalf("expected 3 lines, got %d", len(got))
	}
	want := []string{"3", "4", "5"}
	for i, w := range want {
		if got[i] != w {
			t.Errorf("line %d: want %q, got %q", i, w, got[i])
		}
	}
}

func TestTailLimiter_Reset(t *testing.T) {
	tl, _ := NewTailLimiter(3)
	for _, l := range []string{"a", "b", "c"} {
		tl.Format(l) //nolint:errcheck
	}
	tl.Reset()
	got := tl.Flush()
	if len(got) != 0 {
		t.Errorf("expected empty flush after reset, got %v", got)
	}
}
