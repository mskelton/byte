package storage

import (
	"errors"
	"path"

	"github.com/spf13/viper"
)

func GetByteDir() (string, error) {
	// Get the dir from the config and throw an error if it is not specified
	dir := viper.GetString("dir")
	if dir == "" {
		return "", errors.New("no bytetel dir defined. Please define one in $HOME/.config/byte/config.yaml.")
	}

	// Resolve the dir that was specified in the config
	return resolveDir(dir)
}

func GetBytePath(slug string) (string, error) {
	dir, err := GetByteDir()
	if err != nil {
		return "", err
	}

	return path.Join(dir, slug+".md"), nil
}
