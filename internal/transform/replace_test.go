package transform_test

import (
	"testing"

	"github.com/user/logpipe/internal/transform"
)

func TestNewReplacer_Valid(t *testing.T) {
	_, err := transform.NewReplacer(`\d+`, "NUM", false)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestNewReplacer_EmptyPattern(t *testing.T) {
	_, err := transform.NewReplacer("", "x", false)
	if err == nil {
		t.Fatal("expected error for empty pattern")
	}
}

func TestNewReplacer_InvalidRegex(t *testing.T) {
	_, err := transform.NewReplacer("[invalid", "x", false)
	if err == nil {
		t.Fatal("expected error for invalid regex")
	}
}

func TestReplacer_NoMatch(t *testing.T) {
	r, _ := transform.NewReplacer(`\d+`, "NUM", false)
	out, err := r.Format("no digits here", 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "no digits here" {
		t.Errorf("expected unchanged line, got %q", out)
	}
}

func TestReplacer_SingleMatch(t *testing.T) {
	r, _ := transform.NewReplacer(`\d+`, "NUM", false)
	out, err := r.Format("error 404 occurred", 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "error NUM occurred" {
		t.Errorf("expected %q, got %q", "error NUM occurred", out)
	}
}

func TestReplacer_MultipleMatches(t *testing.T) {
	r, _ := transform.NewReplacer(`\d+`, "NUM", false)
	out, err := r.Format("line 1 and line 2", 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "line NUM and line NUM" {
		t.Errorf("expected %q, got %q", "line NUM and line NUM", out)
	}
}

func TestReplacer_LiteralMode(t *testing.T) {
	r, _ := transform.NewReplacer("foo.bar", "baz", true)
	// dot should be treated literally, not as regex wildcard
	out, err := r.Format("fooXbar and foo.bar", 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "fooXbar and baz" {
		t.Errorf("expected %q, got %q", "fooXbar and baz", out)
	}
}

func TestNewLiteralReplacer_EmptyOld(t *testing.T) {
	_, err := transform.NewLiteralReplacer("", "new")
	if err == nil {
		t.Fatal("expected error for empty old string")
	}
}

func TestNewLiteralReplacer_Valid(t *testing.T) {
	r, err := transform.NewLiteralReplacer("WARN", "WARNING")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	out, _ := r.Format("WARN: disk full", 0)
	if out != "WARNING: disk full" {
		t.Errorf("expected %q, got %q", "WARNING: disk full", out)
	}
}
