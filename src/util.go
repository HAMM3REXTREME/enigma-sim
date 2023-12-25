package main

// Adds a and b and returns 0<result<=26
func addWithOverflow(a, b, modulus int) int {
	result := (a + b) % modulus
	if result <= 0 {
		result += modulus
	}
	return result
}
