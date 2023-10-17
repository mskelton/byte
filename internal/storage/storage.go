package storage

import (
	"errors"
	"os"
	"path"
	"path/filepath"
	"sort"
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

	// Format the byte with Prettier
	err = utils.RunCmd(path.Dir(filename), "bun", "prettier", "--write", filename)
	if err != nil {
		return err
	}

	return nil
}

func Pull() error {
	dir, err := GetByteDir()
	if err != nil {
		return nil
	}

	err = utils.RunCmd(dir, "git", "pull")
	if err != nil {
		return err
	}

	return nil
}

func SyncByte(filename string) error {
	dir := path.Dir(filename)

	// Stage the file
	err := utils.RunCmd(dir, "git", "add", filename)
	if err != nil {
		return err
	}

	// Commit the file
	err = utils.RunCmd(dir, "git", "commit", "-m", "Add "+path.Base(filename))
	if err != nil {
		return err
	}

	// Push the commit
	err = utils.RunCmd(dir, "git", "push")
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

func SearchBytes(predicate func(byte Byte) bool) ([]Byte, error) {
	bytes, err := GetAllBytes()
	if err != nil {
		return []Byte{}, err
	}

	var filtered []Byte

	if predicate == nil {
		filtered = bytes
	} else {
		for _, byte := range bytes {
			if predicate(byte) {
				filtered = append(filtered, byte)
			}
		}
	}

	// Sort the files by name
	sort.Slice(filtered, func(i, j int) bool {
		return strings.Compare(filtered[i].Title, filtered[j].Title) > 0
	})

	return filtered, err
}
