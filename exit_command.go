package main

import (
	"fmt"
	"os"
)

func callbackExit() {
	fmt.Println("See you later!")
	os.Exit(0)
}
