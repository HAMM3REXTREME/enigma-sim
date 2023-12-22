package main

import "unicode"

// Just a convienient way to make a new enigma machine
func newEnigmaMachine(plugboard *biMap, rotors []*rotor, reflector map[int]int) *enigmaMachine {
	return &enigmaMachine{
		plugboard:  plugboard,
		rotorArray: rotors,
		reflector:  reflector,
	}
}

func encryptText(machine enigmaMachine, text string, matchCase bool) string {
	// Fancy text encryption
	encryptedText := "" // Our new encrypted text

	for _, char := range text {
		if unicode.IsLetter(char) { // Only try to encrypt letters
			encryptedChar := encryptRune(machine, char)

			if unicode.IsLower(char) && matchCase { // Our new encrypted character should be lowercase if the starting character was lowercase and matchCase = true
				encryptedText += (string(unicode.ToLower(encryptedChar)))
			} else {
				encryptedText += string(encryptedChar)
			}

		} else {
			encryptedText += string(char)
		}
	}

	return encryptedText
}

func encryptRune(machine enigmaMachine, char rune) rune {
	// Takes a rune and outputs the encrypted rune. Pure function
	var newChar rune

	incrementRotors(machine.rotorArray)
	num, _ := rune2num(char)

	num = machine.plugboard.throughMapF(num)
	num = throughRotorArrayF(machine.rotorArray, num)
	num = throughReflector(machine.reflector, num)
	num = throughRotorArrayB(machine.rotorArray, num)
	num = machine.plugboard.throughMapB(num)

	newChar, _ = num2rune(num)
	return newChar
}

//func checkReflector(mapping []rune)
