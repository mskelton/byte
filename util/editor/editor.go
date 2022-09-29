package editor

import (
	"os"
	"os/exec"
	"runtime"
	"strings"
)

var (
	editor = "vim"
)

func init() {
	if runtime.GOOS == "windows" {
		editor = "notepad"
	}

	if v := os.Getenv("VISUAL"); v != "" {
		editor = v
	} else if e := os.Getenv("EDITOR"); e != "" {
		editor = e
	}
}

type Editor struct {
	Editor   string
	Filename string
}

func (e *Editor) Edit() ([]byte, error) {
	// Overide the default editor if specified in the config
	if e.Editor != "" {
		editor = e.Editor
	}

	args := strings.Split(editor, " ")
	args = append(args, e.Filename)

	// Launch the editor
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return []byte{}, err
	}

	// Read the file contents
	raw, err := os.ReadFile(e.Filename)
	if err != nil {
		return []byte{}, err
	}

	return raw, nil
}
