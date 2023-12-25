package main

import (
	"errors"
	"unicode"
)

// Just a convenient way to make a new enigma machine
func newEnigmaMachine(plugboard *biMap, rotors []*rotor, reflector map[int]int) (*enigmaMachine, error) {

	if err := validateReflector(reflector); err != nil {
		return nil, err
	}

	return &enigmaMachine{
		plugboard:  plugboard,
		rotorArray: rotors,
		reflector:  reflector,
	}, nil
}

// Fancy text encryption
func encryptText(machine enigmaMachine, text string, matchCase bool) (string, error) {
	encryptedText := "" // Our new encrypted text

	for _, char := range text {
		if unicode.IsLetter(char) { // Only try to encrypt letters
			encryptedChar, err := encryptRune(machine, char)
			if err != nil {
				return "", err
			}

			if unicode.IsLower(char) && matchCase { // Our new encrypted character should be lowercase if the starting character was lowercase and matchCase = true
				encryptedText += (string(unicode.ToLower(encryptedChar)))
			} else {
				encryptedText += string(encryptedChar)
			}

		} else {
			encryptedText += string(char)
		}
	}

	return encryptedText, nil
}

func encryptRune(machine enigmaMachine, char rune) (rune, error) {
	// Takes a rune and outputs the encrypted rune. Pure function
	var newChar rune

	incrementRotors(machine.rotorArray)

	num, err := rune2num(char)
	if err != nil {
		return 0, err
	}

	num = machine.plugboard.throughMapF(num)
	num = throughRotorArrayF(machine.rotorArray, num)
	num = throughReflector(machine.reflector, num)
	num = throughRotorArrayB(machine.rotorArray, num)
	num = machine.plugboard.throughMapB(num)

	newChar, err = num2rune(num)
	if err != nil {
		return 0, err
	}

	return newChar, nil
}

// Takes a map[int]int and checks if it can be a physically reflector
// If x maps to y then y must map to x else the machine exhibits buggy behavior
func validateReflector(mapping map[int]int) error {
	for i, r := range mapping {
		if i == r {
			return errors.New("reflector cannot map a letter to itself")
		}
		if mapping[r] != i {
			return errors.New("reflector must form 1:1 connections")
		}
	}
	return nil
}
