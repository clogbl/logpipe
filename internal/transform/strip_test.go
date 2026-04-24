package transform

import (
	"testing"
)

func TestNewStripper_NeitherSet(t *testing.T) {
	_, err := NewStripper()
	if err == nil {
		t.Fatal("expected error when neither prefix nor suffix is set")
	}
}

func TestNewStripper_PrefixOnly(t *testing.T) {
	s, err := NewStripper(WithStripPrefix("INFO "))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if s == nil {
		t.Fatal("expected non-nil Stripper")
	}
}

func TestNewStripper_SuffixOnly(t *testing.T) {
	_, err := NewStripper(WithStripSuffix(" [ok]"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestNewStripper_BothSet(t *testing.T) {
	_, err := NewStripper(WithStripPrefix(">> "), WithStripSuffix(" <<"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestStripper_RemovesPrefix(t *testing.T) {
	s, _ := NewStripper(WithStripPrefix("INFO "))
	out, action := s.Format("INFO hello world", 0)
	if action != FormatKeep {
		t.Errorf("expected FormatKeep, got %v", action)
	}
	if out != "hello world" {
		t.Errorf("expected %q, got %q", "hello world", out)
	}
}

func TestStripper_RemovesSuffix(t *testing.T) {
	s, _ := NewStripper(WithStripSuffix(" [done]"))
	out, action := s.Format("task complete [done]", 0)
	if action != FormatKeep {
		t.Errorf("expected FormatKeep, got %v", action)
	}
	if out != "task complete" {
		t.Errorf("expected %q, got %q", "task complete", out)
	}
}

func TestStripper_RemovesBoth(t *testing.T) {
	s, _ := NewStripper(WithStripPrefix(">> "), WithStripSuffix(" <<"))
	out, _ := s.Format(">> message <<", 0)
	if out != "message" {
		t.Errorf("expected %q, got %q", "message", out)
	}
}

func TestStripper_NoMatch(t *testing.T) {
	s, _ := NewStripper(WithStripPrefix("DEBUG "))
	out, action := s.Format("INFO nothing to strip", 0)
	if action != FormatKeep {
		t.Errorf("expected FormatKeep, got %v", action)
	}
	if out != "INFO nothing to strip" {
		t.Errorf("expected line unchanged, got %q", out)
	}
}

func TestStripper_EmptyLine(t *testing.T) {
	s, _ := NewStripper(WithStripPrefix("INFO "))
	out, action := s.Format("", 0)
	if action != FormatKeep {
		t.Errorf("expected FormatKeep, got %v", action)
	}
	if out != "" {
		t.Errorf("expected empty string, got %q", out)
	}
}
