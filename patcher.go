//go:build !windows

package main

import (
	"fmt"
	"os"
	"runtime"
)

func getDir() {
	fmt.Printf("%s is not supported", runtime.GOOS)
	os.Exit(1)
}
