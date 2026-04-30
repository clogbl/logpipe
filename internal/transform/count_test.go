package transform

import (
	"testing"
)

func TestNewCounter_Valid(t *testing.T) {
	c, err := NewCounter(false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if c == nil {
		t.Fatal("expected non-nil Counter")
	}
}

func TestCounter_IncrementsSilently(t *testing.T) {
	c, _ := NewCounter(false)
	for i := 0; i < 5; i++ {
		out, err := c.Format("hello", i)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if out != "hello" {
			t.Errorf("expected %q, got %q", "hello", out)
		}
	}
	if c.Count() != 5 {
		t.Errorf("expected count 5, got %d", c.Count())
	}
}

func TestCounter_AppendCount(t *testing.T) {
	c, _ := NewCounter(true)
	tests := []struct {
		line string
		want string
	}{
		{"foo", "foo [1]"},
		{"bar", "bar [2]"},
		{"baz", "baz [3]"},
	}
	for i, tt := range tests {
		out, err := c.Format(tt.line, i)
		if err != nil {
			t.Fatalf("row %d unexpected error: %v", i, err)
		}
		if out != tt.want {
			t.Errorf("row %d: expected %q, got %q", i, tt.want, out)
		}
	}
}

func TestCounter_Reset(t *testing.T) {
	c, _ := NewCounter(true)
	c.Format("a", 0)
	c.Format("b", 1)
	if c.Count() != 2 {
		t.Fatalf("expected 2, got %d", c.Count())
	}
	c.Reset()
	if c.Count() != 0 {
		t.Errorf("expected 0 after reset, got %d", c.Count())
	}
	out, _ := c.Format("x", 0)
	if out != "x [1]" {
		t.Errorf("expected %q after reset, got %q", "x [1]", out)
	}
}

func TestCounter_EmptyLine(t *testing.T) {
	c, _ := NewCounter(true)
	out, err := c.Format("", 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != " [1]" {
		t.Errorf("expected %q, got %q", " [1]", out)
	}
}
