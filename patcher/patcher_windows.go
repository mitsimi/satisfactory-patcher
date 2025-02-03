//go:build windows

package patcher

import (
	"gopkg.in/ini.v1"
	"os"
	"path/filepath"
)

func Run() error {
	dir, err := getDir()
	if err != nil {
		return err
	}

	dir = "."

	for file, cfg := range config {

		fullPath := filepath.Join(dir, file)
		if _, err := os.Stat(fullPath); os.IsNotExist(err) {
			createFile(fullPath)
		}

		f, err := ini.Load(fullPath)
		if err != nil {
			return err
		}

		err = ini.ReflectFrom(f, cfg)
		if err != nil {
			return err
		}
	}

	return nil
}
