package util

import (
	"os"
	"path"
	"strings"
)

func ResolvePath(name string) (string, error) {
	if strings.HasPrefix(name, "~/") {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}

		return path.Join(home, name[2:]), nil
	}

	return name, nil
}
