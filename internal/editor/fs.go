package editor

import "os"

func CreateTempFile(pattern string, data []byte) (string, error) {
	file, err := os.CreateTemp("", pattern)
	if err != nil {
		return "", err
	}

	// Close the file when we are done
	defer file.Close()

	// Write the data to the file
	_, err = file.Write(data)
	if err != nil {
		return "", err
	}

	return file.Name(), nil
}
