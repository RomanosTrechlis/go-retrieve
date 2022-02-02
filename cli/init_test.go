package cli

import (
	"bytes"
	"github.com/RomanosTrechlis/go-retrieve/env"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestInit(t *testing.T) {
	pwd, _ := os.Getwd()
	var output bytes.Buffer
	e := env.New(pwd, "data", "config.json", &output)

	nonZeroExit = func(c int) {
		panic("exited")
	}

	// existing config file
	executor := func() { executeInit(e, "", false) }
	assert.PanicsWithValue(t, "exited", executor, "expected to exit with code 1")

	// replacing config file with non-existing file
	executor = func() { executeInit(e, "notExisting.json", false) }
	assert.PanicsWithValue(t, "exited", executor, "expected to exit with code 1")

	// replacing config file with overwrite flag off and dummy file
	executor = func() { executeInit(e, filepath.Join("data", "dummy.json"), false) }
	assert.PanicsWithValue(t, "exited", executor, "expected to exit with code 1")

	// replacing config file with overwrite flag off and legit file
	executor = func() { executeInit(e, filepath.Join("data", "config.json"), false) }
	assert.PanicsWithValue(t, "exited", executor, "expected to exit with code 1")

	// replacing config file with overwrite flag on and legit file
	executor = func() { executeInit(e, filepath.Join("data", "config.json"), true) }
	assert.PanicsWithValue(t, "exited", executor, "expected to exit with code 1")
}
