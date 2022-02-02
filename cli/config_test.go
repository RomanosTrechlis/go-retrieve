package cli

import (
	"bytes"
	"github.com/RomanosTrechlis/go-retrieve/env"
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

var configJson string = `{
  "active": {
    "name": "github",
    "sources": [
      {
        "name": "registry",
        "url": "URL",
        "token": "RETEMP_TOKEN"
      }
    ]
  },
  "profiles": [
    {
      "name": "github",
      "sources": [
        {
          "name": "registry",
          "url": "URL",
          "token": "RETEMP_TOKEN"
        }
      ]
    }
  ]
}
`

func TestConfig(t *testing.T) {
	pwd, _ := os.Getwd()
	var output bytes.Buffer
	e := env.New(pwd, "data", "config.json", &output)
	executeConfig(e, false)
	if output.String() != configJson {
		t.Errorf("failed to print the right information: expected %s, got %s", output.String(), configJson)
	}

	var dumpOutput bytes.Buffer
	e = env.New(pwd, "data", "config.json", &dumpOutput)
	executeConfig(e, true)
	if !strings.HasPrefix(dumpOutput.String(), "(*config.Configuration)") {
		t.Errorf("dump failed to print the right information: expected to begin with %s, got %s",
			"(*config.Configuration)", dumpOutput.String())
	}

	e = env.New(pwd, "./", "config.json", &output)

	nonZeroExit = func(c int) {
		panic("exited")
	}

	executor := func() {
		executeConfig(e, false)
	}
	assert.PanicsWithValue(t, "exited", executor, "expected to exit with code 1")
}
