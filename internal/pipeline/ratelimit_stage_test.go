package pipeline

import (
	"strings"
	"testing"
)

func TestRateLimitStage_InvalidRate(t *testing.T) {
	_, err := New(
		strings.NewReader(""),
		new(strings.Builder),
		NewRateLimitStage(0),
	)
	if err == nil {
		t.Fatal("expected error for rate=0, got nil")
	}
}

func TestRateLimitStage_AllowsLines(t *testing.T) {
	// With a high rate every line should pass through.
	input := "line1\nline2\nline3\n"
	var out strings.Builder
	p, err := New(strings.NewReader(input), &out, NewRateLimitStage(1000))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if err := p.Run(); err != nil {
		t.Fatalf("Run error: %v", err)
	}
	got := out.String()
	for _, want := range []string{"line1", "line2", "line3"} {
		if !strings.Contains(got, want) {
			t.Errorf("expected %q in output, got: %q", want, got)
		}
	}
}

func TestRateLimitStage_DropsExcessLines(t *testing.T) {
	// Rate=1 with burst=1: only the first line should pass; subsequent lines
	// arrive "instantly" (same tick) so tokens are exhausted.
	// We send enough lines that at least some must be dropped.
	lines := make([]string, 20)
	for i := range lines {
		lines[i] = "msg"
	}
	input := strings.Join(lines, "\n") + "\n"
	var out strings.Builder
	p, err := New(strings.NewReader(input), &out, NewRateLimitStage(1))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if err := p.Run(); err != nil {
		t.Fatalf("Run error: %v", err)
	}
	outLines := strings.Count(out.String(), "msg")
	if outLines >= 20 {
		t.Errorf("expected some lines to be dropped, but all %d passed", outLines)
	}
}
