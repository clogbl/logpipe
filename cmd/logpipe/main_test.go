package main

import (
	"bytes"
	"os/exec"
	"strings"
	"testing"
)

func buildBinary(t *testing.T) string {
	t.Helper()
	tmp := t.TempDir()
	bin := tmp + "/logpipe"
	cmd := exec.Command("go", "build", "-o", bin, ".")
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("build failed: %v\n%s", err, out)
	}
	return bin
}

func runBinary(t *testing.T, bin, input string, args ...string) string {
	t.Helper()
	cmd := exec.Command(bin, args...)
	cmd.Stdin = strings.NewReader(input)
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		t.Fatalf("run failed: %v", err)
	}
	return out.String()
}

func TestMain_PassThrough(t *testing.T) {
	bin := buildBinary(t)
	input := "hello world\nfoo bar\n"
	out := runBinary(t, bin, input)
	if out != input {
		t.Errorf("expected %q, got %q", input, out)
	}
}

func TestMain_IncludeFilter(t *testing.T) {
	bin := buildBinary(t)
	input := "hello world\nfoo bar\n"
	out := runBinary(t, bin, input, "--include=hello")
	if !strings.Contains(out, "hello") {
		t.Errorf("expected hello in output, got %q", out)
	}
	if strings.Contains(out, "foo") {
		t.Errorf("unexpected foo in output, got %q", out)
	}
}

func TestMain_ExcludeFilter(t *testing.T) {
	bin := buildBinary(t)
	input := "hello world\nfoo bar\n"
	out := runBinary(t, bin, input, "--exclude=foo")
	if strings.Contains(out, "foo") {
		t.Errorf("unexpected foo in output, got %q", out)
	}
	if !strings.Contains(out, "hello") {
		t.Errorf("expected hello in output, got %q", out)
	}
}

func TestMain_Prefix(t *testing.T) {
	bin := buildBinary(t)
	input := "hello\n"
	out := runBinary(t, bin, input, "--prefix=[LOG]")
	if !strings.HasPrefix(out, "[LOG]") {
		t.Errorf("expected prefix [LOG], got %q", out)
	}
}
