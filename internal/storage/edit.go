package storage

import (
	"github.com/mskelton/byte/internal/editor"
	"github.com/spf13/viper"
)

func EditByte(id string) (string, error) {
	filename, err := GetBytePath(id)
	if err != nil {
		return "", err
	}

	editor := editor.Editor{
		Editor:   viper.GetString("editor"),
		Filename: filename,
	}

	_, err = editor.Edit()
	if err != nil {
		return "", err
	}

	err = FormatByte(filename)
	if err != nil {
		return "", err
	}

	return filename, nil
}
