package main

import (
	"errors"
	"unicode"
)

// The ENIGMA machine rotors were made fairly useless (weak) for large letter sets because of their cyclic nature.
// So I have not bothered to 'improve' them in any way since this is just a simulator, and I have limited it to 26 letters.

func getLetterNumberByChar(char string) (int, error) {
	// The original ENIGMA machine only dealt with the 26 letters of the english alphabet.
	// This function will return the english alphabet number (1 to 26)...
	// ...only if the provided string is 1 character long and is a valid letter.
	// Useful to take user input and convert it to ENIGMA rotor positions.

	if len(char) != 1 { // Check if the string has length 1
		return 0, errors.New("Input must be a single character")
	}
	// Check if the character is a valid English alphabet letter
	r := rune(char[0])
	if !unicode.IsLetter(r) { // this part is case insensitive
		return 0, errors.New("Input must be a valid English alphabet letter")
	}

	upperChar := unicode.ToUpper(r)    // Convert the character to uppercase cause it don't matter for us.
	number := int(upperChar - 'A' + 1) // Calculate the corresponding numerical value (1 to 26)

	return number, nil
}

func getCharByNumber(number int) (string, error) {
	// Basically the reverse of getLetterNumberByChar()
	// This function will return the UPPERCASE English alphabet letter corresponding to the provided number (1 to 26).

	// Check if the number is within the valid range of 1 to 26
	if number < 1 || number > 26 {
		return "", errors.New("Letter Input must be between 1 and 26")
	}

	char := string('A' + number - 1) // Convert the number to the corresponding uppercase English alphabet letter

	return char, nil
}
