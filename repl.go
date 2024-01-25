package main

import (
	"fmt"
	"io"
	"strings"

	"github.com/chzyer/readline"
)

type consoleCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func getCommands() map[string]consoleCommand {
	return map[string]consoleCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "List some location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "List previous location areas",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore {location_area}",
			description: "List encountered pokemon in a location area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch {pokemon_name}",
			description: "Try to catch a pokemon and add to pokedex",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect {pokemon_name}",
			description: "Inspect a pokemon if it exists in the pokedex",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List all pokemon in the pokedex",
			callback:    commandPokedex,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func Start(cfg *config, in io.Reader, out io.Writer) {
	rl, err := readline.NewEx(&readline.Config{
		Prompt:      "Pokedex > ",
		HistoryFile: "/tmp/readline.tmp",
	})
	if err != nil {
		panic(err)
	}
	defer rl.Close()
	// scanner := bufio.NewScanner(in)
	for {
		line, err := rl.Readline()
		if err == readline.ErrInterrupt {
			break
		} else if err == io.EOF {
			break
		}
		// fmt.Print("Pokedex > ")
		// scanner.Scan()
		// text := cleanInput(scanner.Text())
		text := cleanInput(line)
		if len(text) == 0 {
			continue
		}
		commandName := text[0]
		args := []string{}
		if len(text) > 1 {
			args = text[1:]
		}
		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}
