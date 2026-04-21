package transform

import (
	"testing"
	"time"
)

func TestNewRateLimiter_Valid(t *testing.T) {
	rl, err := NewRateLimiter(10)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if rl == nil {
		t.Fatal("expected non-nil RateLimiter")
	}
}

func TestNewRateLimiter_InvalidRate(t *testing.T) {
	for _, rate := range []float64{0, -1, -100} {
		_, err := NewRateLimiter(rate)
		if err == nil {
			t.Errorf("expected error for rate %v, got nil", rate)
		}
	}
}

func TestRateLimiter_BurstAllowed(t *testing.T) {
	// With rate=5, initial tokens=5, first 5 calls should be allowed.
	rl, _ := NewRateLimiter(5)
	fixed := time.Now()
	rl.now = func() time.Time { return fixed }
	rl.lastTick = fixed

	for i := 0; i < 5; i++ {
		if !rl.Allow() {
			t.Errorf("call %d should be allowed", i)
		}
	}
	// 6th call should be dropped
	if rl.Allow() {
		t.Error("6th call should be dropped")
	}
}

func TestRateLimiter_RefillOverTime(t *testing.T) {
	rl, _ := NewRateLimiter(10)
	fixed := time.Now()
	rl.now = func() time.Time { return fixed }
	rl.lastTick = fixed
	rl.tokens = 0 // drain tokens

	// Advance time by 1 second — should refill 10 tokens.
	rl.now = func() time.Time { return fixed.Add(time.Second) }
	for i := 0; i < 10; i++ {
		if !rl.Allow() {
			t.Errorf("call %d after refill should be allowed", i)
		}
	}
	if rl.Allow() {
		t.Error("call after exhausted refill should be dropped")
	}
}

func TestRateLimiter_Format_Passthrough(t *testing.T) {
	rl, _ := NewRateLimiter(1)
	fixed := time.Now()
	rl.now = func() time.Time { return fixed }
	rl.lastTick = fixed

	out := rl.Format("hello")
	if out != "hello" {
		t.Errorf("expected 'hello', got %q", out)
	}
}

func TestRateLimiter_Format_Drops(t *testing.T) {
	rl, _ := NewRateLimiter(1)
	fixed := time.Now()
	rl.now = func() time.Time { return fixed }
	rl.lastTick = fixed

	rl.Format("first") // consume the single token
	out := rl.Format("second")
	if out != "" {
		t.Errorf("expected empty string for dropped line, got %q", out)
	}
}
