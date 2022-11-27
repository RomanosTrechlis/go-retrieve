package cli

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/RomanosTrechlis/go-retrieve/env"
)

var configJson string = `active:
    name: github
    sources:
        - name: registry
          url: URL
          token: ${RT_TOKEN}
profiles:
    - name: github
      sources:
        - name: registry
          url: URL
          token: ${RT_TOKEN}

`

func TestConfig(t *testing.T) {
	pwd, _ := os.Getwd()
	var output bytes.Buffer

	cfg := RegisterConfig(nil)
	if cfg.Name() != "config" {
		t.Errorf("expected 'config', got '%s'", cfg.Name())
	}

	e := env.New(pwd, "./data", "config.yml", &output)
	ExecuteConfig(e, false)
	if output.String() != configJson {
		t.Errorf("failed to print the right information: expected '%s', got '%s'", output.String(), configJson)
	}

	var dumpOutput bytes.Buffer
	e = env.New(pwd, "./data", "config.yml", &dumpOutput)
	ExecuteConfig(e, true)
	if !strings.HasPrefix(dumpOutput.String(), "(*config.Configuration)") {
		t.Errorf("dump failed to print the right information: expected to begin with %s, got %s",
			"(*config.Configuration)", dumpOutput.String())
	}

	e = env.New(pwd, "./", "config.yml", &output)

	NonZeroExit = func(c int) {
		panic("exited")
	}

	executor := func() {
		ExecuteConfig(e, false)
	}
	assert.PanicsWithValue(t, "exited", executor, "expected to exit with code 1")
}
