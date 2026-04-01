package main

import (
	"fmt"
)

func commandInspect(cfg *config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("No Pokemon name was provided!")
	}

	pokemonName := args[0]

	pokemonData, ok := cfg.Pokedex[pokemonName]
	if !ok {
		return fmt.Errorf("You have not caught this Pokemon yet!")
	}

	fmt.Printf("Name: %s\n", pokemonData.Name)
	fmt.Printf("Height: %v\n", pokemonData.Height)
	fmt.Printf("Weight: %v\n", pokemonData.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemonData.Stats {
		fmt.Printf("   - %v: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, ptype := range pokemonData.Types {
		fmt.Printf("   - %v\n", ptype.Type.Name)
	}
	return nil
}
