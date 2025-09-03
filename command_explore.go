package main

import (
	"fmt"
)

func commandExplore(cfg *config) error {
	fmt.Printf("Exploring %s...\n", *cfg.locationName)
	locationName := cfg.locationName
	if locationName == nil {
		return fmt.Errorf("no location name provided")
	}
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
