package registry

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/RomanosTrechlis/go-retrieve/env"
	"github.com/RomanosTrechlis/go-retrieve/util"
)

func Define(e *env.ConfigEnv) (string, error) {
	// todo: check if name already exists inside active profile sources
	name := util.Scan("Please insert new registry's name.")
	protocol := util.Scan("Please select one of the following protocols " +
		"(Not yet implemented)\n0. FTP\n1. HTTP\n2. Local\n>")
	p, err := strconv.Atoi(protocol)
	if err != nil {
		return "", err
	}
	if !util.Contains([]string{"1", "2", "3"}, protocol) {
		return "", fmt.Errorf("the selected protocol wasn't an option")
	}

	r := &Registry{name, Protocol(p), []RegisteredTemplate{}}
	b, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "", err
	}
	return string(b), nil
}
