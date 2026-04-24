package transform

import (
	"testing"
)

func TestNewIndenter_Valid(t *testing.T) {
	_, err := NewIndenter("  ")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestNewIndenter_Empty(t *testing.T) {
	_, err := NewIndenter("")
	if err == nil {
		t.Fatal("expected error for empty indent, got nil")
	}
}

func TestNewSpaceIndenter_Valid(t *testing.T) {
	_, err := NewSpaceIndenter(4)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestNewSpaceIndenter_ZeroSpaces(t *testing.T) {
	_, err := NewSpaceIndenter(0)
	if err == nil {
		t.Fatal("expected error for zero spaces, got nil")
	}
}

func TestNewSpaceIndenter_NegativeSpaces(t *testing.T) {
	_, err := NewSpaceIndenter(-3)
	if err == nil {
		t.Fatal("expected error for negative spaces, got nil")
	}
}

func TestIndenter_Format_Tab(t *testing.T) {
	ind, _ := NewIndenter("\t")
	out, err := ind.Format("hello", 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "\thello" {
		t.Errorf("expected '\\thello', got %q", out)
	}
}

func TestIndenter_Format_Spaces(t *testing.T) {
	ind, _ := NewSpaceIndenter(2)
	out, err := ind.Format("world", 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "  world" {
		t.Errorf("expected '  world', got %q", out)
	}
}

func TestIndenter_Format_EmptyLine(t *testing.T) {
	ind, _ := NewIndenter(">> ")
	out, err := ind.Format("", 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != ">> " {
		t.Errorf("expected '>> ', got %q", out)
	}
}

func TestIndenter_Format_IndexIgnored(t *testing.T) {
	ind, _ := NewIndenter("--")
	out0, _ := ind.Format("line", 0)
	out5, _ := ind.Format("line", 5)
	if out0 != out5 {
		t.Errorf("index should be ignored: got %q and %q", out0, out5)
	}
}
