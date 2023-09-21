package storage

import (
	"errors"
	"os"
	"os/exec"
	"path"

	"github.com/mskelton/byte/internal/storage"
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

func cmd(dir string, command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Dir = dir
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func Pull() error {
	dir, err := storage.GetByteDir()
	if err != nil {
		return nil
	}

	err = cmd(dir, "git", "pull")
	if err != nil {
		return err
	}

	return nil
}

func SyncByte(filename string) error {
	dir := path.Dir(filename)

	// Stage the file
	err := cmd(dir, "git", "add", filename)
	if err != nil {
		return err
	}

	// Commit the file
	err = cmd(dir, "git", "commit", "-m", "Add "+path.Base(filename))
	if err != nil {
		return err
	}

	// Push the commit
	err = cmd(dir, "git", "push")
	if err != nil {
		return err
	}

	return nil
}
