//go:build windows

package patcher

import (
	"os"

	"gopkg.in/ini.v1"
)

func Run() error {
	dir, err := getDir()
	if err != nil {
		return err
	}

	err = os.Chdir(dir)
	if err != nil {
		return err
	}

	for _, path := range INIs {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			err = createFile(path)
			if err != nil {
				return err
			}
		}

		cfg, err := ini.LoadSources(ini.LoadOptions{SkipUnrecognizableLines: true, AllowShadows: true}, path)
		if err != nil {
			return err
		}

		err = applyPatch(cfg, path)
		if err != nil {
			return err
		}

		err = cfg.SaveTo(path)
		if err != nil {
			return err
		}
	}

	return nil
}
