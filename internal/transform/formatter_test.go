package transform

import (
	"strings"
	"testing"
)

func TestNewFormatter_Valid(t *testing.T) {
	for _, format := range []string{"raw", "json", "prefix", "RAW", "JSON"} {
		_, err := NewFormatter(format)
		if err != nil {
			t.Errorf("expected no error for format %q, got %v", format, err)
		}
	}
}

func TestNewFormatter_Invalid(t *testing.T) {
	_, err := NewFormatter("xml")
	if err == nil {
		t.Error("expected error for unknown format, got nil")
	}
}

func TestFormatter_Raw(t *testing.T) {
	f, _ := NewFormatter("raw")
	got := f.Format("hello world")
	if got != "hello world" {
		t.Errorf("raw: expected %q, got %q", "hello world", got)
	}
}

func TestFormatter_Prefix(t *testing.T) {
	f, _ := NewFormatter("prefix", WithPrefix("[APP] "))
	got := f.Format("started")
	if !strings.HasPrefix(got, "[APP] started") {
		t.Errorf("prefix: unexpected output %q", got)
	}
}

func TestFormatter_JSON(t *testing.T) {
	f, _ := NewFormatter("json")
	got := f.Format("something happened")
	if !strings.HasPrefix(got, `{"ts":`) {
		t.Errorf("json: expected JSON output, got %q", got)
	}
	if !strings.Contains(got, `"msg":`) {
		t.Errorf("json: missing msg field in %q", got)
	}
}

func TestFormatter_Timestamp(t *testing.T) {
	f, _ := NewFormatter("raw", WithTimestamp(true))
	got := f.Format("line")
	// timestamp prefix should contain a digit (year)
	if len(got) < 5 || got[0] < '0' || got[0] > '9' {
		t.Errorf("timestamp: expected timestamp prefix in %q", got)
	}
}

func TestFormatter_JSONEscaping(t *testing.T) {
	f, _ := NewFormatter("json")
	got := f.Format(`say "hello"`)
	if strings.Count(got, `\"`) < 2 {
		t.Errorf("json escaping: expected escaped quotes in %q", got)
	}
}

func TestFormatter_PrefixEmpty(t *testing.T) {
	// WithPrefix("") should behave the same as no prefix option.
	f, _ := NewFormatter("prefix", WithPrefix(""))
	got := f.Format("hello")
	if got != "hello" {
		t.Errorf("prefix empty: expected %q, got %q", "hello", got)
	}
}
