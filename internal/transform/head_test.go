package transform

import (
	"testing"
)

func TestNewHeadLimiter_Valid(t *testing.T) {
	h, err := NewHeadLimiter(5)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if h == nil {
		t.Fatal("expected non-nil HeadLimiter")
	}
}

func TestNewHeadLimiter_InvalidMaxLines(t *testing.T) {
	for _, n := range []int{0, -1, -100} {
		_, err := NewHeadLimiter(n)
		if err == nil {
			t.Errorf("expected error for maxLines=%d, got nil", n)
		}
	}
}

func TestHeadLimiter_PassesWithinLimit(t *testing.T) {
	h, _ := NewHeadLimiter(3)
	lines := []string{"line1", "line2", "line3"}
	for _, l := range lines {
		got := h.Format(l)
		if got != l {
			t.Errorf("expected %q, got %q", l, got)
		}
	}
}

func TestHeadLimiter_DropsAfterLimit(t *testing.T) {
	h, _ := NewHeadLimiter(2)
	h.Format("line1")
	h.Format("line2")

	got := h.Format("line3")
	if got != "" {
		t.Errorf("expected empty string after limit, got %q", got)
	}
}

func TestHeadLimiter_ExactlyAtLimit(t *testing.T) {
	h, _ := NewHeadLimiter(1)
	got := h.Format("only")
	if got != "only" {
		t.Errorf("expected %q, got %q", "only", got)
	}
	got = h.Format("dropped")
	if got != "" {
		t.Errorf("expected empty string, got %q", got)
	}
}

func TestHeadLimiter_Reset(t *testing.T) {
	h, _ := NewHeadLimiter(1)
	h.Format("first")

	if got := h.Format("dropped"); got != "" {
		t.Errorf("expected drop before reset, got %q", got)
	}

	h.Reset()

	if got := h.Format("after-reset"); got != "after-reset" {
		t.Errorf("expected line after reset, got %q", got)
	}
}
