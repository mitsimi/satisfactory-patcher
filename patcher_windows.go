//go:build windows

package main

import (
	"io/fs"
	"os"
	"path/filepath"
)

func getDir() fs.FS {
	return os.DirFS(filepath.Join("%LOCALAPPDATA%", "FactoryGame", "Saved", "Config", "Windows"))
}
