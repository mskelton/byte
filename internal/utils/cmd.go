package utils

import "os/exec"

func RunCmd(dir string, command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Dir = dir
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
