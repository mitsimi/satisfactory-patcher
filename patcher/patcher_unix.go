//go:build !windows

package patcher

import (
	"fmt"
	"runtime"
)

func Run() error {
	fmt.Printf("%s is not supported\n", runtime.GOOS)

	return nil
}
