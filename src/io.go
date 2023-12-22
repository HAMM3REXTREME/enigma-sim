package main

import (
	"bufio"
	"errors"
	"os"
	"unicode"
)

// The ENIGMA machine rotors were made fairly useless (weak) for large letter sets because of their cyclic nature.
// So I have not bothered to 'improve' them in any way since this is just a simulator, and I have limited it to 26 letters.
// The original ENIGMA machine only dealt with the 26 letters of the English alphabet.

func rune2num(char rune) (int, error) {
	// rune [A to Z] ---> int (1 to 26)
	if !unicode.IsLetter(char) { // Bail if rune is not a letter
		return 0, errors.New("input must be a valid english alphabet letter")
	}

	upperChar := unicode.ToUpper(char) // Uppercase/lowercase should land on the same spot from 1 to 26
	number := int(upperChar - 'A' + 1) // Calculate the corresponding numerical value (1 to 26)

	return number, nil
}

func num2rune(number int) (rune, error) {
	// int (1 to 26) ---> rune [A to Z]
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

func runes2map(runes []rune) map[int]int{
	// Takes a []rune and returns a map[int][int]  (ints are 1 to 26)
	mapping := make(map[int]int)
    letterList := make([]int,len(runes)) // List of 1 to 26
	for x := 0;x<len(runes);x++{ // Convert from []rune{'A','Z'} to []int{1,26}
		letterList[x],_=rune2num(runes[x])
	}
	for i, r := range letterList { // Make map from int letter list
		mapping[i+1] = int(r)
	}
	return mapping
}


func runes2biMap(runes []rune) *biMap {
	// Take a []rune and return a biMap (1 to 26)
	// Convert A-Z to 1-26 and convert to bimap
	forward := make(map[int]int)
	backward := make(map[int]int)
    letterList := make([]int,len(runes)) // List of 1 to 26
	for x := 0;x<len(runes);x++{ // Convert from []rune{'A','Z'} to []int{1,26}
		letterList[x],_=rune2num(runes[x])
	}

	for i, r := range letterList { // Make bimap from int letter list
		forward[i+1] = int(r)
		backward[int(r)] = i + 1
	}

	return &biMap{
		forward:  forward,
		backward: backward,
	}
}