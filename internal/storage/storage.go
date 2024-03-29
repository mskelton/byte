package storage

import (
	"errors"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/mskelton/byte/internal/utils"
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

func ReadByte(id string) ([]byte, error) {
	filename, err := GetBytePath(id)
	if err != nil {
		return []byte{}, err
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		return []byte{}, err
	}

	return data, nil
}

// Format the byte with Prettier
func FormatByte(filename string) error {
	return utils.RunCmd(path.Dir(filename), "bun", "prettier", "--write", filename)
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

	err = FormatByte(filename)
	if err != nil {
		return err
	}

	return nil
}

func GetAllBytes() ([]Byte, error) {
	dir, err := GetByteDir()
	if err != nil {
		return []Byte{}, err
	}

	var bytes []Byte
	err = filepath.Walk(dir, func(name string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			content, err := os.ReadFile(name)
			if err != nil {
				return err
			}

			// Get the id from the filename
			id := strings.TrimSuffix(path.Base(name), path.Ext(name))

			byte, err := ParseByte(id, content)
			if err != nil {
				return err
			}

			bytes = append(bytes, byte)
		}

		return nil
	})

	return bytes, err

}
