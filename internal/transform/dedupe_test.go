package transform

import "testing"

func TestDeduplicator_FirstLine(t *testing.T) {
	d := NewDeduplicator()
	out, ok := d.Format("hello")
	if !ok {
		t.Fatal("expected first line to be kept")
	}
	if out != "hello" {
		t.Fatalf("expected %q, got %q", "hello", out)
	}
}

func TestDeduplicator_ConsecutiveDuplicate(t *testing.T) {
	d := NewDeduplicator()
	d.Format("hello")
	_, ok := d.Format("hello")
	if ok {
		t.Fatal("expected duplicate line to be suppressed")
	}
}

func TestDeduplicator_NonConsecutiveDuplicate(t *testing.T) {
	d := NewDeduplicator()
	d.Format("hello")
	d.Format("world")
	out, ok := d.Format("hello")
	if !ok {
		t.Fatal("expected non-consecutive duplicate to be kept")
	}
	if out != "hello" {
		t.Fatalf("expected %q, got %q", "hello", out)
	}
}

func TestDeduplicator_MultipleConsecutive(t *testing.T) {
	d := NewDeduplicator()
	lines := []string{"a", "a", "a", "b", "b", "c"}
	expected := []string{"a", "b", "c"}
	var got []string
	for _, l := range lines {
		if out, ok := d.Format(l); ok {
			got = append(got, out)
		}
	}
	if len(got) != len(expected) {
		t.Fatalf("expected %v, got %v", expected, got)
	}
	for i := range expected {
		if got[i] != expected[i] {
			t.Errorf("index %d: expected %q, got %q", i, expected[i], got[i])
		}
	}
}

func TestDeduplicator_Reset(t *testing.T) {
	d := NewDeduplicator()
	d.Format("hello")
	d.Reset()
	_, ok := d.Format("hello")
	if !ok {
		t.Fatal("expected line to be kept after reset")
	}
}

func TestDeduplicator_EmptyLines(t *testing.T) {
	d := NewDeduplicator()
	d.Format("")
	_, ok := d.Format("")
	if ok {
		t.Fatal("expected consecutive empty lines to be suppressed")
	}
}

func TestDeduplicator_ResetRestoresEmptyState(t *testing.T) {
	// Verify that Reset allows any line (including the previously seen one)
	// to be treated as new, and that subsequent duplicates are still suppressed.
	d := NewDeduplicator()
	d.Format("hello")
	d.Format("world")
	d.Reset()

	// After reset, "world" should be accepted again.
	if _, ok := d.Format("world"); !ok {
		t.Fatal("expected line to be kept immediately after reset")
	}
	// A consecutive duplicate following reset should still be suppressed.
	if _, ok := d.Format("world"); ok {
		t.Fatal("expected duplicate to be suppressed after reset")
	}
}
