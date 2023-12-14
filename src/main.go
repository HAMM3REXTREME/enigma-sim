package main

/*
ENIGMA Machine Diagram:

               Bulbs
               /|\
                |
Keyboard-----> Plugboard <-----> Rotor-1 <-----> Rotor-2 <-----> Rotor-3 <-----> Reflector

*/

import (
	"errors"
	"fmt"
	"unicode"
)

func main() {
	var letterMap plugboard
	letterMap.PlugboardMapF := map[string]string{
		"A": "N",
		"B": "O",
		"C": "P",
		"D": "Q",
		"E": "R",
		"F": "S",
		"G": "T",
		"H": "U",
		"I": "V",
		"J": "W",
		"K": "X",
		"L": "Y",
		"M": "Z",
		"N": "A",
		"O": "B",
		"P": "C",
		"Q": "D",
		"R": "E",
		"S": "F",
		"T": "G",
		"U": "H",
		"V": "I",
		"W": "J",
		"X": "K",
		"Y": "L",
		"Z": "M"}
	testStr := "A"
	number, err := getLetterNumberByChar(testStr)
	if err != nil {
		fmt.Printf("Error for '%s': %s\n", testStr, err)
	} else {
		fmt.Printf("Value for '%s': %d\n", testStr, number)
		fmt.Printf("After ENIGMA Plugboard: %s\n", getThroughPlugboard(letterMap, testStr))
	}
}

func getThroughPlugboard(letterMap plugboard, letterIn string) string {
	// This function returns a letter only if present in the provided 'plugboard'
	// Otherwise returns the same letter as provided since nothing is plugged into our hypothetical ENIGMA plugboard
	if mappedLetter, ok := letterMap.PlugboardMapF[letterIn]; ok {
		return mappedLetter // Return the mapped letter if present
	}
	return letterIn // Return the same letter if not present in the plugboard
}

func getLetterNumberByChar(char string) (int, error) {
	// The original ENIGMA machine only dealt with the 26 letters of the english alphabet.
	// This function will return the english alphabet number (1 to 26)...
	// ...only if the provided string is 1 character long and is a valid letter.
	// Useful for 'authenticity' for eg: an user interface.

	// Check if the string has length 1
	if len(char) != 1 {
		return 0, errors.New("input must be a single character")
	}
	// Check if the character is a valid English alphabet letter
	r := rune(char[0])
	if !unicode.IsLetter(r) { // this part is case insensitive
		return 0, errors.New("input must be a valid English alphabet letter")
	}

	upperChar := unicode.ToUpper(r)    // Convert the character to uppercase cause it don't matter for us.
	number := int(upperChar - 'A' + 1) // Calculate the corresponding numerical value (1 to 26)

	return number, nil
}

func updateRotors() {
	// This function basically increments the rotors position on the axle.
	// The first rotor gets spun every time this function is called, essentially every time a key is pressed on our ENIGMA machine's 'keyboard'
}

func getThroughRotorsF(rotorArray [3]rotor) string {
	// This function only goes through the Rotor A to B to C
	// Does not reflect
}

func getThroughReflector(reflector rotor) string {
	// Reflector is just a limited rotor, but with a limitation...
	// No letter can map to itself, a cryptographic weakness caused by the same wires being used for forwards and backwards legs.
}
