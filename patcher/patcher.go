package patcher

import (
	"os"
	"path/filepath"
)

var config = map[string]func() interface{}{
	EngineIni:      PatchedEngine,
	GameIni:        PatchedGame,
	ScalabilityIni: PatchedScalability,
}

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

func createFile(path string) {
	f, _ := os.Create(path)
	f.Close()
}
