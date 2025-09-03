package main

import (
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("no location name provided")
	}
	locationName := args[0]
	fmt.Printf("Exploring %s...\n", locationName)
	location, err := cfg.pokeapiClient.GetLocation(locationName)
	if err != nil {
		return err
	}
	fmt.Print("Found Pokemon:\n")
	for _, pokemon := range location.Encounters {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}
	fmt.Println("")
	return nil
}
