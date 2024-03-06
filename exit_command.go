package main

import (
	"fmt"
	"os"
)

func callbackExit(confg *config, args ...string) error {
	fmt.Println("See you later!")
	os.Exit(0)
	return nil
}
