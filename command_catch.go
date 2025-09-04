package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("no pokemon name provided")
	}
	pokemonName := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}
	randNum := rand.Intn(pokemon.BaseExperience)
	if randNum > 40 {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		return nil
	}
	fmt.Printf("%s escaped!\n", pokemon.Name)
	fmt.Println("You may now inspect it with the inspect command.")

	cfg.caughtPokemon[pokemon.Name] = pokemon
	return nil
}
