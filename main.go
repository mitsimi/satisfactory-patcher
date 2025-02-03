package main

import (
	"fmt"
	"github.com/mitsimi/satisfactory-patcher/patcher"
)

func main() {
	err := patcher.Run()
	if err != nil {
		fmt.Println(err)
	}
}
