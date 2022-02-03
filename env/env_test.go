package env

import (
	"os"
	"path/filepath"
	"testing"
)

func TestConfigEnv_ConfigFilePath(t *testing.T) {
	e := DefaultConfigEnv()
	h, _ := os.UserHomeDir()
	dir := filepath.Join(h, e.ConfigDir)
	if dir != e.ConfigPath() {
		t.Errorf("failed: expected %v, got %v", dir, e.ConfigPath())
	}

	file := filepath.Join(dir, e.ConfigName)
	if file != e.ConfigFilePath() {
		t.Errorf("failed: expected %v, got %v", dir, e.ConfigPath())
	}

	e = New("", e.ConfigDir, e.ConfigName, e.Writer())
	if e.ConfigDir != e.ConfigPath() {
		t.Errorf("failed: expected %v, got %v", dir, e.ConfigPath())
	}
}
