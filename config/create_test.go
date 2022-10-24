package config

import (
	"os"
	"testing"
)

func TestMore(t *testing.T) {
	scanMore(t, "input enter", "\n", false)
	scanMore(t, "YES enter", "YES", true)
	scanMore(t, "yes enter", "yes", true)
	scanMore(t, "Y enter", "yes", true)
	scanMore(t, "y enter", "yes", true)
}

func scanMore(t *testing.T, d, i string, o bool) {
	content := []byte(i)
	tmpfile, err := os.CreateTemp("", "example")
	if err != nil {
		t.Errorf("failed to create temp file")
	}

	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write(content); err != nil {
		t.Errorf("failed to create temp file")
	}

	if _, err := tmpfile.Seek(0, 0); err != nil {
		t.Errorf("failed to create temp file")
	}

	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }() // Restore original Stdin

	os.Stdin = tmpfile
	b := more("test")
	if b != o {
		t.Errorf("failed %s: expected %v, got %v", d, o, b)
	}

	_ = tmpfile.Close()
}
