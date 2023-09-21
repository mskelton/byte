package storage

import (
	"errors"
	"os"
	"path"

	"github.com/spf13/viper"
)

func GetByteDir() (string, error) {
	// Get the dir from the config and throw an error if it is not specified
	dir := viper.GetString("dir")
	if dir == "" {
		return "", errors.New("no byte dir defined. Please define one in $HOME/.config/byte/config.yaml.")
	}

	// Resolve the dir that was specified in the config
	return resolveDir(dir)
}

func GetBytePath(id string) (string, error) {
	dir, err := GetByteDir()
	if err != nil {
		return "", err
	}

	return path.Join(dir, id+".md"), nil
}

func WriteByte(filename string, data []byte) error {
	// Create the directory for the byte if necessary
	err := os.Mkdir(path.Dir(filename), os.ModePerm)
	if err != nil && !(errors.Is(err, os.ErrExist)) {
		return err
	}

	// Write the byte to the file system
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func SyncByte(filename string) error {

	return nil
}
