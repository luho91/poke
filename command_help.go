package main

import(
	"fmt"
)

func commandHelp(*config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")
	for _, command := range commands {
		fmt.Println(command.name, ": ", command.description)
	}
	return nil
}
