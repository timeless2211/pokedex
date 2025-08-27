package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		input := scanner.Text()
		cleanInput := cleanInput(input)
		if len(cleanInput) > 0 && len(cleanInput[0]) > 0 {
			upperCaseFirst := strings.ToUpper(string(cleanInput[0][0]))
			restOfWord := ""
			if len(cleanInput[0]) > 1 {
				restOfWord = cleanInput[0][1:]
			}
			textWithoutFirstWord := strings.Join(cleanInput[1:], " ")
			fmt.Printf("You entered: %s%s %s\n", upperCaseFirst, restOfWord, textWithoutFirstWord)
		}
	}
}
