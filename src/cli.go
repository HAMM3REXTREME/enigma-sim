package main

import (
	"fmt"
)

func userEnigmaInput(machine *enigmaMachine) {
	fmt.Printf("Enter text into the ENIGMA machine: ")
	userInput := scanLine()
	enigmaOutput := encryptText(*machine, userInput, true)
	fmt.Printf("ENIGMA output: %s\n", string(enigmaOutput))
}
