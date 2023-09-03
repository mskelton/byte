package storage

import (
	"os"
	"path"
	"strings"
)

func resolveDir(name string) (string, error) {
	if strings.HasPrefix(name, "~/") {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}

		return path.Join(home, name[2:]), nil
	}

	return name, nil
}
