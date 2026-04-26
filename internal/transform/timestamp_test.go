package transform

import (
	"testing"
	"time"
)

var fixedTime = time.Date(2024, 6, 1, 12, 0, 0, 0, time.UTC)

func fixedClock() time.Time { return fixedTime }

func TestNewTimestampPrepender_Valid(t *testing.T) {
	tp, err := NewTimestampPrepender(time.RFC3339)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if tp == nil {
		t.Fatal("expected non-nil prepender")
	}
}

func TestNewTimestampPrepender_EmptyFormat(t *testing.T) {
	_, err := NewTimestampPrepender("")
	if err == nil {
		t.Fatal("expected error for empty format")
	}
}

func TestTimestampPrepender_Format(t *testing.T) {
	tp, _ := NewTimestampPrepender(time.RFC3339, WithClock(fixedClock))
	out, err := tp.Format("hello world", 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := "2024-06-01T12:00:00Z hello world"
	if out != want {
		t.Errorf("got %q, want %q", out, want)
	}
}

func TestTimestampPrepender_CustomFormat(t *testing.T) {
	tp, _ := NewTimestampPrepender("2006-01-02", WithClock(fixedClock))
	out, err := tp.Format("line", 1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := "2024-06-01 line"
	if out != want {
		t.Errorf("got %q, want %q", out, want)
	}
}

func TestTimestampPrepender_IndexIgnored(t *testing.T) {
	tp, _ := NewTimestampPrepender("15:04:05", WithClock(fixedClock))
	out1, _ := tp.Format("a", 0)
	out2, _ := tp.Format("a", 99)
	if out1 != out2 {
		t.Errorf("index should be ignored: %q vs %q", out1, out2)
	}
}

func TestTimestampPrepender_EmptyLine(t *testing.T) {
	tp, _ := NewTimestampPrepender("15:04:05", WithClock(fixedClock))
	out, err := tp.Format("", 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out == "" {
		t.Error("expected non-empty output even for empty input")
	}
}
