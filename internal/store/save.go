package store

import (
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

func SaveFile(name string, data []byte, folder string) error {
	target := filepath.Join(folder, name)
	if err := os.WriteFile(target, data, 0755); err != nil {
		return err
	}
	return nil
}
