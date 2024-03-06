package main

import (
	"fmt"
	"os"
)

func callbackExit(confg *config) error {
	fmt.Println("See you later!")
	os.Exit(0)
	return nil
}
