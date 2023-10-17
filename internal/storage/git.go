package storage

import (
	"path"

	"github.com/mskelton/byte/internal/utils"
)

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

func SyncByte(filename string, commitPrefix string) error {
	dir := path.Dir(filename)

	// Stage the file
	err := utils.RunCmd(dir, "git", "add", filename)
	if err != nil {
		return err
	}

	// Commit the file
	err = utils.RunCmd(dir, "git", "commit", "-m", commitPrefix+" "+path.Base(filename))
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
