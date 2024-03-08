package main

import "fmt"

func commandHelp(cfg *config, args ...string) error {
	availableCommands := getCommands()

	fmt.Println()
	fmt.Println("Pokedex CLI")
	fmt.Println("Usage:")
	fmt.Println()

	for _, cmd := range availableCommands {
		fmt.Printf(" - %s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
