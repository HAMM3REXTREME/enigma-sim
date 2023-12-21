package main

import (
	"fmt"
)

func userEnigmaInput(plugboard *obfuscatorMap, rotorArray []*rotor, reflector map[int]int) {
	var userInput string
	validInput := false

	for !validInput {
		fmt.Print("Enter a message to encrypt with the ENIGMA machine (only A-Z or a-z allowed): ")
		fmt.Scanln(&userInput)

		// Check if the input contains only A-Z or a-z characters
		validInput = isValidInput(userInput)
		if !validInput {
			fmt.Println("Invalid input. Please enter only A-Z or a-z characters.")
		}
	}

	//debugObfuscateFull(plugboard, rotorArray, userInputChars, reflector)
	enigmaOutput := stringEnigma(plugboard, rotorArray, reflector, userInput)

	fmt.Printf("ENIGMA output: %s\n", string(enigmaOutput))
}
