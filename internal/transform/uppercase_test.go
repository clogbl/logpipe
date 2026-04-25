package transform

import (
	"testing"
)

func TestNewUppercaser_Upper(t *testing.T) {
	u := NewUppercaser(true)
	if u == nil {
		t.Fatal("expected non-nil Uppercaser")
	}
	if !u.upper {
		t.Error("expected upper=true")
	}
}

func TestNewUppercaser_Lower(t *testing.T) {
	u := NewUppercaser(false)
	if u == nil {
		t.Fatal("expected non-nil Uppercaser")
	}
	if u.upper {
		t.Error("expected upper=false")
	}
}

func TestUppercaser_ToUpper(t *testing.T) {
	u := NewUppercaser(true)
	tests := []struct {
		input string
		want  string
	}{
		{"hello world", "HELLO WORLD"},
		{"ERROR: disk full", "ERROR: DISK FULL"},
		{"", ""},
		{"already UPPER", "ALREADY UPPER"},
	}
	for _, tt := range tests {
		got, err := u.Format(tt.input, 0)
		if err != nil {
			t.Errorf("Format(%q) unexpected error: %v", tt.input, err)
		}
		if got != tt.want {
			t.Errorf("Format(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}

func TestUppercaser_ToLower(t *testing.T) {
	u := NewUppercaser(false)
	tests := []struct {
		input string
		want  string
	}{
		{"HELLO WORLD", "hello world"},
		{"ERROR: Disk Full", "error: disk full"},
		{"", ""},
		{"already lower", "already lower"},
	}
	for _, tt := range tests {
		got, err := u.Format(tt.input, 0)
		if err != nil {
			t.Errorf("Format(%q) unexpected error: %v", tt.input, err)
		}
		if got != tt.want {
			t.Errorf("Format(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}

func TestUppercaser_IndexIgnored(t *testing.T) {
	u := NewUppercaser(true)
	got1, _ := u.Format("test", 0)
	got2, _ := u.Format("test", 99)
	if got1 != got2 {
		t.Errorf("index should not affect output: %q != %q", got1, got2)
	}
}
