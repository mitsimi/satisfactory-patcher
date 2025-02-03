package patcher

import (
	"os"
	"path/filepath"
)

func getDir() (string, error) {
	localAppData, err := os.UserCacheDir()
	if err != nil {
		return "", err
	}

	dir := filepath.Join(localAppData, "FactoryGame", "Saved", "Config", "WindowsNoEditor")

	if _, err = os.Stat(dir); os.IsNotExist(err) {
		return "", err
	}

	return dir, nil
}

func createFile(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}

	f.Close()
	return nil
}
