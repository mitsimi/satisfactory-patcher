//go:build !windows

package patcher

import (
	"fmt"
	"os"
	"runtime"
)

func Run() error {
	fmt.Printf("%s is not supported", runtime.GOOS)
	os.Exit(1)
}
