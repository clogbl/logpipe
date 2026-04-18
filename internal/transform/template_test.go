package transform

import (
	"testing"
)

func TestNewTemplateFormatter_Valid(t *testing.T) {
	_, err := NewTemplateFormatter("[{{.Index}}] {{.Line}}")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestNewTemplateFormatter_Empty(t *testing.T) {
	_, err := NewTemplateFormatter("   ")
	if err == nil {
		t.Fatal("expected error for empty template")
	}
}

func TestNewTemplateFormatter_Invalid(t *testing.T) {
	_, err := NewTemplateFormatter("{{.Unclosed")
	if err == nil {
		t.Fatal("expected error for invalid template syntax")
	}
}

func TestTemplateFormatter_Format_Line(t *testing.T) {
	f, err := NewTemplateFormatter("msg={{.Line}}")
	if err != nil {
		t.Fatalf("setup failed: %v", err)
	}
	out, err := f.Format("hello", 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "msg=hello" {
		t.Errorf("expected 'msg=hello', got %q", out)
	}
}

func TestTemplateFormatter_Format_Index(t *testing.T) {
	f, _ := NewTemplateFormatter("{{.Index}}: {{.Line}}")
	out, err := f.Format("world", 3)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "3: world" {
		t.Errorf("expected '3: world', got %q", out)
	}
}

func TestTemplateFormatter_Format_OnlyLine(t *testing.T) {
	f, _ := NewTemplateFormatter("{{.Line}}")
	out, _ := f.Format("raw", 99)
	if out != "raw" {
		t.Errorf("expected 'raw', got %q", out)
	}
}
