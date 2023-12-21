package main

import (
	"errors"
	"unicode"
)

// The ENIGMA machine rotors were made fairly useless (weak) for large letter sets because of their cyclic nature.
// So I have not bothered to 'improve' them in any way since this is just a simulator, and I have limited it to 26 letters.

func char2num(char rune) (int, error) {
	// The original ENIGMA machine only dealt with the 26 letters of the English alphabet.
	// This function will return the English alphabet number (1 to 26) from a rune
	// ...only if the provided rune is a valid letter.
	// Useful to take user input and convert it to ENIGMA rotor positions.

	if !unicode.IsLetter(char) { // this part is case insensitive
		return 0, errors.New("input must be a valid english alphabet letter")
	}

	upperChar := unicode.ToUpper(char) // Convert the character to uppercase cause it doesn't matter for us.
	number := int(upperChar - 'A' + 1) // Calculate the corresponding numerical value (1 to 26)

	return number, nil
}

func num2char(number int) (rune, error) {
	// This function will return the UPPERCASE English alphabet letter corresponding to the provided number (1 to 26).

	// Check if the number is within the valid range of 1 to 26
	if number < 1 || number > 26 {
		return 0, errors.New("letter input must be between 1 and 26")
	}

	char := rune('A' + number - 1) // Convert the number to the corresponding uppercase English alphabet letter

	return char, nil
}
