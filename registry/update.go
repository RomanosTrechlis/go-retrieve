package registry

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"

	"github.com/google/go-cmp/cmp"

	"github.com/RomanosTrechlis/go-retrieve/config"
	"github.com/RomanosTrechlis/go-retrieve/dl"
	"github.com/RomanosTrechlis/go-retrieve/env"
	"github.com/RomanosTrechlis/go-retrieve/util"
)

func UpdateRegistry(e *env.ConfigEnv) (string, error) {
	c, err := config.LoadConfig(e)
	if err != nil {
		return "", err
	}

	var b bytes.Buffer
	sources := c.Active.Sources
	for _, s := range sources {
		err := findDifferences(e, s, &b)
		if err != nil {
			return "", err
		}
	}
	return b.String(), nil
}

func findDifferences(e *env.ConfigEnv, s *config.ConfigurationSource, b *bytes.Buffer) error {
	remoteB, err := dl.Download(s.Url, s.Token)
	if err != nil {
		return err
	}

	filename := filepath.Join(e.ConfigPath(), fmt.Sprintf("%s.%s", s.Name, e.Suffix()))
	localB, err := util.LoadFile(filename)
	if err != nil {
		err = os.WriteFile(filename, remoteB, 0755)
		if err != nil {
			return err
		}
	}

	//diff := cmp.Diff(localB, remoteB)
	remote, err := Unmarshal(remoteB, e.IsJson())
	if err != nil {
		return err
	}

	local, err := Unmarshal(localB, e.IsJson())
	if err != nil {
		return err
	}
	diff := cmp.Diff(local, remote)
	if diff == "" {
		return nil
	}

	_, _ = fmt.Fprintf(b, "------------- %s -----------\n", s.Name)
	_, _ = fmt.Fprintln(b, diff)
	return nil
}
