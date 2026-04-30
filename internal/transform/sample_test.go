package transform_test

import (
	"testing"

	"github.com/user/logpipe/internal/transform"
)

func TestNewSampler_Valid(t *testing.T) {
	s, err := transform.NewSampler(3)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if s == nil {
		t.Fatal("expected non-nil Sampler")
	}
}

func TestNewSampler_InvalidRate(t *testing.T) {
	for _, n := range []int{0, -1, -100} {
		_, err := transform.NewSampler(n)
		if err == nil {
			t.Errorf("expected error for n=%d", n)
		}
	}
}

func TestSampler_EveryLine(t *testing.T) {
	s, _ := transform.NewSampler(1)
	for i := 0; i < 5; i++ {
		out, err := s.Format("line", i)
		if err != nil {
			t.Errorf("line %d: unexpected skip", i)
		}
		if out != "line" {
			t.Errorf("line %d: expected 'line', got %q", i, out)
		}
	}
}

func TestSampler_EveryThird(t *testing.T) {
	s, _ := transform.NewSampler(3)
	results := make([]bool, 9)
	for i := 0; i < 9; i++ {
		_, err := s.Format("x", i)
		results[i] = (err == nil)
	}
	// lines 3, 6, 9 (1-indexed) should pass
	expected := []bool{false, false, true, false, false, true, false, false, true}
	for i, want := range expected {
		if results[i] != want {
			t.Errorf("index %d: got pass=%v, want %v", i, results[i], want)
		}
	}
}

func TestSampler_DropsNonMultiples(t *testing.T) {
	s, _ := transform.NewSampler(2)
	// first call: count=1, not multiple of 2 → skip
	_, err := s.Format("a", 0)
	if err != transform.ErrSkip {
		t.Errorf("expected ErrSkip on first call, got %v", err)
	}
	// second call: count=2, multiple of 2 → pass
	out, err := s.Format("b", 1)
	if err != nil {
		t.Errorf("unexpected error on second call: %v", err)
	}
	if out != "b" {
		t.Errorf("expected 'b', got %q", out)
	}
}

func TestSampler_RateOne_AllPass(t *testing.T) {
	s, _ := transform.NewSampler(1)
	lines := []string{"alpha", "beta", "gamma"}
	for i, l := range lines {
		out, err := s.Format(l, i)
		if err != nil || out != l {
			t.Errorf("index %d: expected pass-through, got out=%q err=%v", i, out, err)
		}
	}
}
