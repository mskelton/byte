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

// Prompt the user to enter some text using their default editor
func (e *Editor) Prompt() ([]byte, error) {
	// Create a new temp file that can be edited
	f, err := os.CreateTemp("", e.Filename)
	if err != nil {
		return []byte{}, err
	}

	// Ensure the temp file is cleaned up
	defer func() {
		_ = os.Remove(f.Name())
	}()

	// Close the fd to prevent the editor unable to save file
	if err := f.Close(); err != nil {
		return []byte{}, err
	}

	// Overide the default editor if specified in the config
	if e.Editor != "" {
		editor = e.Editor
	}

	args := strings.Split(editor, " ")
	args = append(args, f.Name())

	// Launch the editor
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return []byte{}, err
	}

	// Read the temp file contents
	raw, err := os.ReadFile(f.Name())
	if err != nil {
		return []byte{}, err
	}

	return raw, nil
}
