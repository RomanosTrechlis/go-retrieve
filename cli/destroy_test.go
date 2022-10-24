package cli

import (
	"bytes"
	"os"
	"testing"

	"github.com/otiai10/copy"
	"github.com/stretchr/testify/assert"

	"github.com/RomanosTrechlis/go-retrieve/env"
)

func TestDestroy(t *testing.T) {
	pwd, _ := os.Getwd()
	var output bytes.Buffer

	nonZeroExit = func(c int) {
		panic("exited")
	}

	setup(t)

	// destroying without config.json file
	e := env.New(pwd, "data_to_delete", "nonExisting.json", &output)
	executor := func() {
		executeDestroy(e)
	}
	assert.PanicsWithValue(t, "exited", executor, "expected to exit with code 1")

	setup(t)

	// successfully destroying env
	e = env.New(pwd, "data_to_delete", "config.json", &output)
	executeDestroy(e)
	_, err := os.ReadFile(e.ConfigFilePath())
	if err == nil {
		t.Errorf("failed to destroy env")
	}
}

func setup(t *testing.T) {
	err := copy.Copy("data", "data_to_delete")
	if err != nil {
		t.Errorf("failed to setup test: %v", err)
	}
}
