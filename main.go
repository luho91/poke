package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokédex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())
		fmt.Println(fmt.Sprintf("Your command was: %s", input[0]))
	}
}
