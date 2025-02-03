//go:build windows

package patcher

import (
	"gopkg.in/ini.v1"
	"os"
)

func Run() error {
	dir, err := getDir()
	if err != nil {
		return err
	}

	dir = "."

	err = os.Chdir(dir)
	if err != nil {
		return err
	}

	f, err := ini.LoadSources(ini.LoadOptions{SkipUnrecognizableLines: true}, EngineIni)
	if err != nil {
		return err
	}
	v := 104857600
	err = f.Append(Engine{
		Player: EnginePlayer{
			ConfiguredInternetSpeed: v,
			ConfiguredLanSpeed:      v,
		},
		IpNetDriver: EngineIpNetDriver{
			MaxClientRate:         v,
			MaxInternetClientRate: v,
		},
		EpicNetDriver: EngineEpicNetDriver{
			MaxClientRate:         v,
			MaxInternetClientRate: v,
		},
	})
	if err != nil {
		return err
	}

	err = f.SaveTo(EngineIni)

	/*for file, cfg := range config {
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

		err = f.SaveTo(fullPath)
	}*/

	return nil
}
