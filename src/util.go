package main

import (
	"errors"
)

func addWithOverflow(a, b, modulus int) int { // Adds a and b and returns 0<result<=26
	result := (a + b) % modulus
	if result <= 0 {
		result += modulus
	}
	return result
}

func validateMap(mapping map[int]int, isReflector bool) error { // Takes a map[int]int and checks if it can be a physically possible plugboard/reflector
	// If x maps to y then y must map to x else you gonna get bugs
	for i, r := range mapping {
		if i == r && isReflector {
			return errors.New("reflector cannot map to itself") // Reflectors cannot map any letter to itself
		}

		if mapping[r] != i {
			return errors.New("plugboard/reflector must form 1:1 connections")
		}
	}

	return nil
}
