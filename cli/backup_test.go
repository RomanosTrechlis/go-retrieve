package cli

import (
	"bytes"
	"github.com/RomanosTrechlis/go-retrieve/env"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestBackup(t *testing.T) {
	pwd, _ := os.Getwd()
	var output bytes.Buffer

	nonZeroExit = func(c int) {
		panic("exited")
	}

	// non-existing config file
	e := env.New(pwd, "data", "nonExisting.json", &output)
	executor := func() {
		executeBackup(e, "")
	}
	assert.PanicsWithValue(t, "exited", executor, "expected to exit with code 1")

	// json file of a different format
	e = env.New(pwd, "data", "dummy.json", &output)
	executor = func() {
		executeBackup(e, "")
	}
	assert.PanicsWithValue(t, "exited", executor, "expected to exit with code 1")

	// successfully backup file to pwd
	e = env.New(pwd, "data", "config.json", &output)
	executeBackup(e, "")
	_, err := ioutil.ReadFile(e.ConfigName)
	if err != nil {
		t.Errorf("failed to backup config file")
	}
	_ = os.Remove(e.ConfigName)

	// successfully backup file to specific location
	path := filepath.Join("data", "backup.json")
	executeBackup(e, path)
	_, err = ioutil.ReadFile(path)
	if err != nil {
		t.Errorf("failed to backup config file")
	}
	_ = os.Remove(path)
}
