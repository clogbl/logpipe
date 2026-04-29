package transform

import (
	"testing"
)

func TestNewColumnExtractor_Valid(t *testing.T) {
	_, err := NewColumnExtractor([]int{0, 1})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestNewColumnExtractor_NoIndices(t *testing.T) {
	_, err := NewColumnExtractor(nil)
	if err == nil {
		t.Fatal("expected error for nil indices")
	}
}

func TestNewColumnExtractor_EmptyIndices(t *testing.T) {
	_, err := NewColumnExtractor([]int{})
	if err == nil {
		t.Fatal("expected error for empty indices")
	}
}

func TestColumnExtractor_DefaultSeparator(t *testing.T) {
	ext, _ := NewColumnExtractor([]int{0, 2})
	out, err := ext.Format("alpha beta gamma", 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "alpha gamma" {
		t.Errorf("expected 'alpha gamma', got %q", out)
	}
}

func TestColumnExtractor_CustomSeparator(t *testing.T) {
	ext, _ := NewColumnExtractor([]int{1, 3}, WithColumnSeparator(","))
	out, err := ext.Format("a,b,c,d", 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "b d" {
		t.Errorf("expected 'b d', got %q", out)
	}
}

func TestColumnExtractor_CustomJoiner(t *testing.T) {
	ext, _ := NewColumnExtractor([]int{0, 1}, WithColumnJoiner(","))
	out, err := ext.Format("foo bar baz", 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "foo,bar" {
		t.Errorf("expected 'foo,bar', got %q", out)
	}
}

func TestColumnExtractor_OutOfBoundsIndex(t *testing.T) {
	ext, _ := NewColumnExtractor([]int{0, 10})
	out, err := ext.Format("only three words", 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// out-of-bounds indices are silently skipped
	if out != "only" {
		t.Errorf("expected 'only', got %q", out)
	}
}

func TestColumnExtractor_EmptyLine(t *testing.T) {
	ext, _ := NewColumnExtractor([]int{0})
	out, err := ext.Format("", 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "" {
		t.Errorf("expected empty string, got %q", out)
	}
}
