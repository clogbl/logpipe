package transform

import (
	"testing"
)

func TestNewJSONFieldExtractor_Valid(t *testing.T) {
	e, err := NewJSONFieldExtractor([]string{"level", "msg"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if e == nil {
		t.Fatal("expected non-nil extractor")
	}
}

func TestNewJSONFieldExtractor_NoFields(t *testing.T) {
	_, err := NewJSONFieldExtractor([]string{})
	if err == nil {
		t.Fatal("expected error for empty fields")
	}
}

func TestJSONFieldExtractor_ValidLine(t *testing.T) {
	e, _ := NewJSONFieldExtractor([]string{"level", "msg"})
	out, err := e.Format(`{"level":"info","msg":"started","ts":1234}`, 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "info started" {
		t.Errorf("got %q, want %q", out, "info started")
	}
}

func TestJSONFieldExtractor_CustomSeparator(t *testing.T) {
	e, _ := NewJSONFieldExtractor([]string{"level", "msg"}, WithJSONSeparator(" | "))
	out, _ := e.Format(`{"level":"warn","msg":"retrying"}`, 0)
	if out != "warn | retrying" {
		t.Errorf("got %q, want %q", out, "warn | retrying")
	}
}

func TestJSONFieldExtractor_InvalidJSON(t *testing.T) {
	e, _ := NewJSONFieldExtractor([]string{"level"})
	raw := "not json at all"
	out, err := e.Format(raw, 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != raw {
		t.Errorf("expected raw line passthrough, got %q", out)
	}
}

func TestJSONFieldExtractor_MissingField(t *testing.T) {
	e, _ := NewJSONFieldExtractor([]string{"level", "missing"})
	raw := `{"level":"debug"}`
	out, err := e.Format(raw, 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != raw {
		t.Errorf("expected raw line passthrough, got %q", out)
	}
}

func TestJSONFieldExtractor_SingleField(t *testing.T) {
	e, _ := NewJSONFieldExtractor([]string{"msg"})
	out, _ := e.Format(`{"msg":"hello world","code":42}`, 0)
	if out != "hello world" {
		t.Errorf("got %q, want %q", out, "hello world")
	}
}
