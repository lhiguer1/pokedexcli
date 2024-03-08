package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config, args ...string) error {
	res, err := cfg.pokeAPIClient.ListLocations(cfg.nextLocationAreaURL)
	if err != nil {
		return err
	}

	fmt.Println("Location Areas:")
	for _, area := range res.Results {
		fmt.Printf(" - %s\n", area.Name)
	}

	cfg.nextLocationAreaURL = res.Next
	cfg.prevLocationAreaURL = res.Previous

	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if cfg.prevLocationAreaURL == nil {
		return errors.New("you are on the first page")
	}

	res, err := cfg.pokeAPIClient.ListLocations(cfg.prevLocationAreaURL)
	if err != nil {
		return err
	}

	fmt.Println("Location Areas:")
	for _, area := range res.Results {
		fmt.Printf(" - %s\n", area.Name)
	}

	cfg.nextLocationAreaURL = res.Next
	cfg.prevLocationAreaURL = res.Previous

	return nil
}
