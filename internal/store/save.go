package store

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
)

func CheckCreateDataFolder(dataFolder string) error {
	if _, err := os.Stat(dataFolder); os.IsNotExist(err) {
		if err := os.MkdirAll(dataFolder, 0755); err != nil {
			return err
		}
	}
	return nil
}

func Compact(data *[]byte) ([]byte, error) {
	cmpt := &bytes.Buffer{}
	if err := json.Compact(cmpt, *data); err != nil {
		return nil, err
	}
	return cmpt.Bytes(), nil
}

func SaveFile(name string, data []byte, folder string) error {
	target := filepath.Join(folder, name)
	if err := os.WriteFile(target, data, 0755); err != nil {
		return err
	}
	return nil
}
