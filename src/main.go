package main

import (
	"fmt"
	"strings"
)

/*
ENIGMA Machine Diagram:

               Bulbs
               /|\
                |
Keyboard-----> Plugboard <-----> Rotor-1 <-----> Rotor-2 <-----> Rotor-3 <-----> Reflector

*/

func main() {
	// Test data for obfuscatorMap (Plugboard)
	plugboardData := map[int]int{
		1: 5, 2: 10, 3: 15, 4: 20, 5: 1,
		6: 6, 7: 11, 8: 16, 9: 21, 10: 2,
		11: 7, 12: 12, 13: 17, 14: 22, 15: 3,
		16: 8, 17: 13, 18: 18, 19: 23, 20: 4,
		21: 9, 22: 14, 23: 19, 24: 24, 25: 25, 26: 26,
	}
	plugboard := newBimap(plugboardData)

	// Test data for rotor 1
	rotor1Data := map[int]int{
		1: 5, 2: 10, 3: 15, 4: 20, 5: 1,
		6: 6, 7: 11, 8: 16, 9: 21, 10: 2,
		11: 7, 12: 12, 13: 17, 14: 22, 15: 3,
		16: 8, 17: 13, 18: 18, 19: 23, 20: 4,
		21: 9, 22: 14, 23: 19, 24: 24, 25: 25, 26: 26,
	}
	rotor1 := &rotor{
		rotorMap:        newBimap(rotor1Data),
		rotorSpinOffset: 1,
		nextRotorSpin:   5,
	}

	// Test data for rotor 2
	rotor2Data := map[int]int{
		1: 10, 2: 15, 3: 20, 4: 1, 5: 6,
		6: 11, 7: 16, 8: 21, 9: 2, 10: 7,
		11: 12, 12: 17, 13: 22, 14: 3, 15: 8,
		16: 13, 17: 18, 18: 23, 19: 4, 20: 9,
		21: 14, 22: 19, 23: 24, 24: 25, 25: 26, 26: 5,
	}
	rotor2 := &rotor{
		rotorMap:        newBimap(rotor2Data),
		rotorSpinOffset: 1,
		nextRotorSpin:   10,
	}

	// Test data for rotor 3
	rotor3Data := map[int]int{
		1: 15, 2: 20, 3: 1, 4: 6, 5: 11,
		6: 16, 7: 21, 8: 2, 9: 7, 10: 12,
		11: 17, 12: 22, 13: 3, 14: 8, 15: 13,
		16: 18, 17: 23, 18: 4, 19: 9, 20: 14,
		21: 19, 22: 24, 23: 25, 24: 26, 25: 5, 26: 10,
	}
	rotor3 := &rotor{
		rotorMap:        newBimap(rotor3Data),
		rotorSpinOffset: 10,
		nextRotorSpin:   0,
	}
	rotorArray := []*rotor{rotor1, rotor2, rotor3}

	// Very important for this to correct
	// if x maps to y then y must map to x else you gonna get bugs
	reflector := map[int]int{
		1: 2, 2: 1, 3: 4, 4: 3, 5: 6,
		6: 5, 7: 8, 8: 7, 9: 10, 10: 9,
		11: 12, 12: 11, 13: 14, 14: 13, 15: 16,
		16: 15, 17: 18, 18: 17, 19: 20, 20: 19,
		21: 22, 22: 21, 23: 24, 24: 23, 25: 26, 26: 25,
	}

	for {
		userInput(plugboard, rotorArray, reflector)
		//debugObfuscateFull(plugboard, rotorArray, []string{"H", "I", "L", "E", "R"}, reflector)
	}

}

func userInput(plugboard *obfuscatorMap, rotorArray []*rotor, reflector map[int]int) {
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

	userInput = strings.ToUpper(userInput)
	userInputChars := strings.Split(userInput, "")

	//enigmaOutput := getFullEnigma(plugboard, rotorArray, userInputChars, reflector)
	debugObfuscateFull(plugboard, rotorArray, userInputChars, reflector)

	//fmt.Printf("ENIGMA output: %s\n", strings.Join(enigmaOutput, ""))
}
