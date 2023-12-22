package main

import "unicode"

func addWithOverflow(a, b, modulus int) int {
	// Adds a and b and returns 0<result<=26
	result := (a + b) % modulus
	if result <= 0 {
		result += modulus
	}
	return result
}

func isValidInput(input string) bool {
	// Only allow A to Z (case insensitive)
	for _, char := range input {
		if !unicode.IsLetter(char) {
			return false
		}
	}
	return true
}
