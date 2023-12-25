package main

import (
	"errors"
	"fmt"
	"unicode"
)

/* The ENIGMA machine rotors were made fairly useless (weak) for large letter sets because of their cyclic nature.
So I have not bothered to 'improve' them in any way since this is just a simulator, and I have limited it to 26 letters.
The original ENIGMA machine only dealt with the 26 letters of the English alphabet. */

func rune2num(char rune) (int, error) { // rune [A to Z] ---> int (1 to 26)
	if !unicode.IsLetter(char) { // Bail if rune is not a letter
		return 0, errors.New("input must be a valid english alphabet letter")
	}

	upperChar := unicode.ToUpper(char) // Uppercase/lowercase should land on the same spot from 1 to 26
	number := int(upperChar - 'A' + 1) // Calculate the corresponding numerical value (1 to 26)

	return number, nil
}
func num2rune(number int) (rune, error) { // int (1 to 26) ---> rune [A to Z] (along with any errors)
	if number < 1 || number > Letters {
		return 0, errors.New("letter input must be between 1 and 26")
	}

	char := rune('A' + number - 1) // Convert the number to the corresponding uppercase English alphabet letter
	return char, nil
}

func runes2map(runes []rune) (map[int]int, error) { // Takes a []rune and returns a map[int][int]  (ints are 1 to 26)
	finalMap := make(map[int]int)
	for i, r := range runes {
		letterID, err := rune2num(r)
		if err != nil {
			return nil, fmt.Errorf("rune2map: Error converting rune to letter ID at index %d: %v", i, err)
		}
		finalMap[i+1] = letterID
	}
	return finalMap, nil
}

func runes2biMap(runes []rune) (*biMap, error) { // Takes a []rune and returns a biMap (1 to 26)
	forward := make(map[int]int)
	backward := make(map[int]int)

	for i, r := range runes {
		letterID, err := rune2num(r)
		if err != nil {
			return nil, fmt.Errorf("rune2bimap: Error converting rune to letter ID at index %d: %v", i, err)
		}
		forward[i+1] = letterID
		backward[letterID] = i + 1
	}

	return &biMap{
		forward:  forward,
		backward: backward,
	}, nil
}
