package filter

import (
	"testing"
)

func TestNewPattern_Valid(t *testing.T) {
	p, err := NewPattern("err", `ERROR`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if p.Label != "err" {
		t.Errorf("expected label 'err', got %q", p.Label)
	}
}

func TestNewPattern_Invalid(t *testing.T) {
	_, err := NewPattern("bad", `[invalid`)
	if err == nil {
		t.Fatal("expected error for invalid regex, got nil")
	}
}

func TestPattern_Match(t *testing.T) {
	p, _ := NewPattern("warn", `WARN`)
	if !p.Match("2024/01/01 WARN something happened") {
		t.Error("expected match")
	}
	if p.Match("INFO all good") {
		t.Error("expected no match")
	}
}

func TestFilter_Keep_Include(t *testing.T) {
	f, err := NewFilter([]string{`ERROR`, `WARN`}, false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !f.Keep("ERROR disk full") {
		t.Error("expected ERROR line to be kept")
	}
	if f.Keep("INFO startup complete") {
		t.Error("expected INFO line to be dropped")
	}
}

func TestFilter_Keep_Exclude(t *testing.T) {
	f, err := NewFilter([]string{`DEBUG`}, true)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if f.Keep("DEBUG verbose output") {
		t.Error("expected DEBUG line to be dropped")
	}
	if !f.Keep("ERROR something broke") {
		t.Error("expected non-DEBUG line to be kept")
	}
}

func TestFilter_Keep_NoPatterns(t *testing.T) {
	f, _ := NewFilter(nil, false)
	if !f.Keep("anything goes") {
		t.Error("expected line to be kept when no patterns defined")
	}
}
