package pipeline

import (
	"bytes"
	"strings"
	"testing"
	"time"

	"github.com/user/logpipe/internal/transform"
)

var tsFixed = time.Date(2024, 6, 1, 12, 0, 0, 0, time.UTC)

func fixedTSClock() time.Time { return tsFixed }

func TestTimestampStage_InvalidFormat(t *testing.T) {
	_, err := NewTimestampStage("")
	if err == nil {
		t.Fatal("expected error for empty format")
	}
}

func TestTimestampStage_Valid(t *testing.T) {
	s, err := NewTimestampStage(time.RFC3339, transform.WithClock(fixedTSClock))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if s == nil {
		t.Fatal("expected non-nil stage")
	}
}

func TestTimestampStage_Output(t *testing.T) {
	s, _ := NewTimestampStage("2006-01-02", transform.WithClock(fixedTSClock))

	input := "error: disk full"
	r := strings.NewReader(input + "\n")
	var buf bytes.Buffer

	p, _ := New(r, &buf)
	p.AddStage(s)
	if err := p.Run(); err != nil {
		t.Fatalf("pipeline error: %v", err)
	}

	out := strings.TrimSpace(buf.String())
	want := "2024-06-01 error: disk full"
	if out != want {
		t.Errorf("got %q, want %q", out, want)
	}
}
