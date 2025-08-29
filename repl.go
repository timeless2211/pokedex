package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type config struct {
	Next     bool
	Previous bool
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

var commands = []cliCommand{
	{
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    callbackExit,
	},
	{
		name:        "help",
		description: "Show help information",
		callback:    callbackHelp,
	},
	{
		name:        "map",
		description: "Show the map of all Pokemons",
		callback:    callbackGetMap,
	},
	{
		name:        "mapb",
		description: "Show the previous map of all Pokemons",
		callback:    callbackGetMap,
	},
}

func callbackExit(*config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func callbackHelp(*config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	fmt.Println()
	return nil
}

func callbackGetMap(*config) error {
	locationAreas, error := fetchLocationAreas(&config{})
	if error != nil {
		return error
	}
	for _, area := range locationAreas.Results {
		fmt.Println(area.(map[string]any)["name"])
	}
	return nil
}

func startRepl() {
	cfg := &config{}
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		text := cleanInput(reader.Text())
		if len(text) == 0 {
			continue
		}
		textStr := strings.Join(text, " ")

		for _, cmd := range commands {
			if cmd.name == textStr {
				if cmd.name == "map" {
					cfg.Next = true
				}
				if cmd.name == "mapb" {
					cfg.Previous = true
				}
				cmd.callback(cfg)
			}
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
