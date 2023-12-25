package main

import (
	"errors"
)

// Adds a and b and returns 0<result<=26
func addWithOverflow(a, b, modulus int) int {
	result := (a + b) % modulus
	if result <= 0 {
		result += modulus
	}
	return result
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
