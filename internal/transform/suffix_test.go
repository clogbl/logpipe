package transform_test

import (
	"testing"

	"github.com/yourorg/logpipe/internal/transform"
)

func TestNewLineSuffixer_Valid(t *testing.T) {
	s, err := transform.NewLineSuffixer(" [end]")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if s == nil {
		t.Fatal("expected non-nil suffixer")
	}
}

func TestNewLineSuffixer_Empty(t *testing.T) {
	_, err := transform.NewLineSuffixer("")
	if err == nil {
		t.Fatal("expected error for empty suffix")
	}
}

func TestLineSuffixer_Format_Basic(t *testing.T) {
	s, _ := transform.NewLineSuffixer(" [ok]")
	out, err := s.Format("hello", 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "hello [ok]" {
		t.Errorf("expected %q, got %q", "hello [ok]", out)
	}
}

func TestLineSuffixer_Format_EmptyLine(t *testing.T) {
	s, _ := transform.NewLineSuffixer("!!")
	out, err := s.Format("", 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "!!" {
		t.Errorf("expected %q, got %q", "!!", out)
	}
}

func TestLineSuffixer_Format_IndexIgnored(t *testing.T) {
	s, _ := transform.NewLineSuffixer("-X")
	out0, _ := s.Format("line", 0)
	out5, _ := s.Format("line", 5)
	if out0 != out5 {
		t.Errorf("index should be ignored: %q vs %q", out0, out5)
	}
}

func TestLineSuffixer_Format_Whitespace(t *testing.T) {
	s, _ := transform.NewLineSuffixer("  ")
	out, err := s.Format("trim", 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "trim  " {
		t.Errorf("expected %q, got %q", "trim  ", out)
	}
}
