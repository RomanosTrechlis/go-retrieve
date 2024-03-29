package cli

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/RomanosTrechlis/go-retrieve/env"
)

func TestBackup(t *testing.T) {
	pwd, _ := os.Getwd()
	var output bytes.Buffer

	NonZeroExit = func(c int) {
		panic("exited")
	}

	bu := RegisterBackupCmd()
	if bu.Name() != "backup" {
		t.Errorf("expected 'backup', got '%s'", bu.Name())
	}

	// non-existing config file
	e := env.New(pwd, "./data", "nonExisting.json", &output)
	executor := func() {
		executeBackup(e, "")
	}
	assert.PanicsWithValue(t, "exited", executor, "expected to exit with code 1")

	// json file of a different format
	e = env.New(pwd, "./data", "dummy.json", &output)
	executor = func() {
		executeBackup(e, "")
	}
	assert.PanicsWithValue(t, "exited", executor, "expected to exit with code 1")

	// successfully backup file to pwd
	e = env.New(pwd, "./data", "config.json", &output)
	executeBackup(e, "")
	_, err := os.ReadFile(e.ConfigName)
	if err != nil {
		t.Errorf("failed to backup config file")
	}
	_ = os.Remove(e.ConfigName)

	// successfully backup file to specific location
	path := filepath.Join("./data", "backup.json")
	executeBackup(e, path)
	_, err = os.ReadFile(path)
	if err != nil {
		t.Errorf("failed to backup config file")
	}
	_ = os.Remove(path)
}
