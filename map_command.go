package main

import (
	"errors"
	"fmt"
)

func callbackMap(confg *config) error {
	resp, err := confg.pokeapiClient.ListLocationArea(confg.nextLocationAreaURL)
	if err != nil {
		return err
	}

	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}

	confg.nextLocationAreaURL = resp.Next
	confg.previousLocationAreaURL = resp.Previous
	return nil
}

func callbackMapb(confg *config) error {
	if confg.previousLocationAreaURL == nil {
		return errors.New("you're on the first page")
	}
	resp, err := confg.pokeapiClient.ListLocationArea(confg.previousLocationAreaURL)
	if err != nil {
		return err
	}

	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}

	confg.nextLocationAreaURL = resp.Next
	confg.previousLocationAreaURL = resp.Previous
	return nil
}
