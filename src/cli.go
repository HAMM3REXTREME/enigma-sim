package main

import (
	"fmt"
)

func userEnigmaInput(plugboard *obfuscatorMap, rotorArray []*rotor, reflector map[int]int) {
	fmt.Printf("Enter text into the ENIGMA machine: ")
	userInput := scanLine()

	//debugObfuscateFull(plugboard, rotorArray, userInputChars, reflector)
	enigmaOutput := encryptText(plugboard, rotorArray, reflector, userInput, true)

	fmt.Printf("ENIGMA output: %s\n", string(enigmaOutput))
}
