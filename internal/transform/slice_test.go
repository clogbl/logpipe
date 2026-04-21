package transform

import (
	"testing"
)

func TestNewSlicer_Valid(t *testing.T) {
	_, err := NewSlicer(" ", 0, -1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestNewSlicer_EmptyDelimiter(t *testing.T) {
	_, err := NewSlicer("", 0, -1)
	if err == nil {
		t.Fatal("expected error for empty delimiter")
	}
}

func TestNewSlicer_NegativeStart(t *testing.T) {
	_, err := NewSlicer(" ", -1, -1)
	if err == nil {
		t.Fatal("expected error for negative start")
	}
}

func TestNewSlicer_EndBeforeStart(t *testing.T) {
	_, err := NewSlicer(" ", 3, 1)
	if err == nil {
		t.Fatal("expected error when end < start")
	}
}

func TestSlicer_FullLine(t *testing.T) {
	s, _ := NewSlicer(" ", 0, -1)
	out, err := s.Format("a b c d", 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "a b c d" {
		t.Errorf("expected 'a b c d', got %q", out)
	}
}

func TestSlicer_MiddleFields(t *testing.T) {
	s, _ := NewSlicer(" ", 1, 2)
	out, err := s.Format("a b c d", 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "b c" {
		t.Errorf("expected 'b c', got %q", out)
	}
}

func TestSlicer_StartBeyondFields(t *testing.T) {
	s, _ := NewSlicer(" ", 10, -1)
	_, err := s.Format("a b c", 0)
	if err != ErrSkip {
		t.Errorf("expected ErrSkip, got %v", err)
	}
}

func TestSlicer_EndClampedToLast(t *testing.T) {
	s, _ := NewSlicer(",", 1, 100)
	out, err := s.Format("x,y,z", 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "y,z" {
		t.Errorf("expected 'y,z', got %q", out)
	}
}

func TestSlicer_SingleField(t *testing.T) {
	s, _ := NewSlicer(":", 0, 0)
	out, err := s.Format("INFO:server started", 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "INFO" {
		t.Errorf("expected 'INFO', got %q", out)
	}
}
