package main

import (
	"github.com/luho91/poke/internal/pokecache"
)

type cliCommand struct {
	name		string
	description	string
	callback	func(*config) error
}

type config struct {
	Next		string
	Previous	string
	Cache		*pokecache.Cache
}

var commands map[string]cliCommand

func init()	{
	commands = map[string]cliCommand {
		"exit": {
			name:			"exit",
			description:	"Exit the Pokedex",
			callback:		commandExit,
		},
		"help": {
			name:			"help",
			description:	"Displays a help message",
			callback:		commandHelp,
		},
		"map": {
			name:			"map",
			description:	"Displays the names of the next 20 locations in the Pokemon world",
			callback:		commandMapNext,
		},
		"mapb": {
			name:			"mapb",
			description:	"Displays the names of the previous 20 locations in the Pokemon world",
			callback:		commandMapPrevious,
		},
	}
}
