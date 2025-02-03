package main

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

func main() {
	cfg, err := ini.Load("test.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	eng := new(Engine)
	err = cfg.MapTo(eng)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("%+v\n", *eng)

	// Now, make some changes and save it
	eng.Player.ConfiguredInternetSpeed = 1234567890

	err = ini.ReflectFrom(cfg, eng)
	if err != nil {
		fmt.Printf("Fail to reflect file: %v", err)
		os.Exit(1)
	}

	fmt.Printf("%+v\n", *eng)

	cfg.SaveTo("test.ini")
}
