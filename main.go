package main

import (
	"bufio"
	"os"
	"fmt"
	"time"
	"github.com/luho91/poke/internal/pokecache"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	config := config {
		Next:		"https://pokeapi.co/api/v2/location-area/",
		Previous:	"",
		Cache:		pokecache.NewCache(10 * time.Second),
	}

	for {
		fmt.Print("Pokédex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())

		if len(input) == 0 {
			continue
		}

		first_input := input[0]
		args := input[1:]
		command, ok := commands[first_input]

		if ok {
			err := command.callback(&config, args)
			if err != nil {
				fmt.Printf("Error happened: %v\n", err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}
