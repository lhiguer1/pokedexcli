package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		scanner.Scan()

		cleanedInput := cleanInput(scanner.Text())
		args := []string{}
		if len(cleanedInput) > 1 {
			args = cleanedInput[1:]
		}

		if len(cleanedInput) == 0 {
			continue
		}

		commandName := cleanedInput[0]
		availableCommands := getCommands()

		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("invalid command")
			continue
		}

		if err := command.callback(cfg, args...); err != nil {
			fmt.Println(err)
		}

	}

}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Attempt to catch a pokemon",
			callback:    commandCatch,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "Explore a location",
			callback:    commandExplore,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func cleanInput(str string) []string {
	return strings.Fields(strings.ToLower(str))
}
