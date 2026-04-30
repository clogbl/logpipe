package transform_test

import (
	"testing"

	"github.com/user/logpipe/internal/transform"
)

func TestNewSqueezer_WhitespaceOnly(t *testing.T) {
	s := transform.NewWhitespaceSqueezer()
	if s == nil {
		t.Fatal("expected non-nil squeezer")
	}
}

func TestNewSqueezer_ValidChar(t *testing.T) {
	s, err := transform.NewSqueezer('-', false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if s == nil {
		t.Fatal("expected non-nil squeezer")
	}
}

func TestNewSqueezer_ZeroCharNonWhitespace(t *testing.T) {
	_, err := transform.NewSqueezer(0, false)
	if err == nil {
		t.Fatal("expected error for zero char with whitespaceOnly=false")
	}
}

func TestSqueezer_CollapseSpaces(t *testing.T) {
	s := transform.NewWhitespaceSqueezer()
	out, err := s.Format("hello   world", 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "hello world" {
		t.Errorf("expected %q, got %q", "hello world", out)
	}
}

func TestSqueezer_MixedWhitespace(t *testing.T) {
	s := transform.NewWhitespaceSqueezer()
	out, err := s.Format("foo\t\t  bar", 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "foo bar" {
		t.Errorf("expected %q, got %q", "foo bar", out)
	}
}

func TestSqueezer_NoRepeats(t *testing.T) {
	s := transform.NewWhitespaceSqueezer()
	out, err := s.Format("clean line", 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "clean line" {
		t.Errorf("expected %q, got %q", "clean line", out)
	}
}

func TestSqueezer_SpecificChar(t *testing.T) {
	s, _ := transform.NewSqueezer('-', false)
	out, err := s.Format("foo---bar--baz", 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "foo-bar-baz" {
		t.Errorf("expected %q, got %q", "foo-bar-baz", out)
	}
}

func TestSqueezer_SpecificChar_NoMatch(t *testing.T) {
	s, _ := transform.NewSqueezer('*', false)
	out, err := s.Format("no dashes here", 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "no dashes here" {
		t.Errorf("expected %q, got %q", "no dashes here", out)
	}
}

func TestSqueezer_EmptyLine(t *testing.T) {
	s := transform.NewWhitespaceSqueezer()
	out, err := s.Format("", 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "" {
		t.Errorf("expected empty string, got %q", out)
	}
}
