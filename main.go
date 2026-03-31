package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	config := config {
		Next:		"https://pokeapi.co/api/v2/location-area/",
		Previous:	"",
	}

	for {
		fmt.Print("Pokédex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())

		if len(input) == 0 {
			continue
		}

		first_input := input[0]
		command, ok := commands[first_input]

		if ok {
			command.callback(&config)
		} else {
			fmt.Println("Unknown command")
		}
	}
}
