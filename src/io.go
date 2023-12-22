package main

import (
	"bufio"
	"errors"
	"os"
	"unicode"
)

// The ENIGMA machine rotors were made fairly useless (weak) for large letter sets because of their cyclic nature.
// So I have not bothered to 'improve' them in any way since this is just a simulator, and I have limited it to 26 letters.

func rune2num(char rune) (int, error) {
	// This function will return the English alphabet number (1 to 26) from a valid rune.
	// The original ENIGMA machine only dealt with the 26 letters of the English alphabet.
	if !unicode.IsLetter(char) { // Bail if rune is not a letter
		return 0, errors.New("input must be a valid english alphabet letter")
	}

	upperChar := unicode.ToUpper(char) // Uppercase/lowercase should land on the same spot from 1 to 26
	number := int(upperChar - 'A' + 1) // Calculate the corresponding numerical value (1 to 26)

	return number, nil
}

func num2rune(number int) (rune, error) {
	// This function will return the UPPERCASE English alphabet letter corresponding to the provided number (1 to 26).
	// You can compare the case the of the original rune/string if you insist on having the same case (though it is not historically accurate)

	// Check if the number is within the valid range of 1 to 26
	if number < 1 || number > Letters {
		return 0, errors.New("letter input must be between 1 and 26")
	}

	char := rune('A' + number - 1) // Convert the number to the corresponding uppercase English alphabet letter

	return char, nil
}

func scanLine() string {
	// Scans a line. Will count spaces.
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	//fmt.Println("Your input:", line)
	return line
}
