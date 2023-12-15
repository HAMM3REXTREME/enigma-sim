package main

import "fmt"

func addWithOverflow(a, b, modulus int) int {
	// Adds a and b
	// Returns 0<result<=26
	result := (a + b) % modulus
	if result <= 0 {
		result += modulus
	}
	return result
}

func updateRotors(rotorArray [3]*rotor) {
	// Increment the position of the first rotor
	rotorArray[0].rotorSpinOffset = addWithOverflow(rotorArray[0].rotorSpinOffset, 1, 26)

	// Check if the notch is reached for the next rotors and spin them accordingly
	for i := 1; i < len(rotorArray); i++ {
		if rotorArray[i-1].rotorSpinOffset == rotorArray[i-1].nextRotorSpin+1 {
			// Spin the next rotor
			rotorArray[i].rotorSpinOffset = addWithOverflow(rotorArray[i].rotorSpinOffset, 1, 26)
		}
	}

}

func getThroughRotorsF(rotorArray [3]*rotor, letterID int) int {
	// This function only goes through the RotorArray First to Last
	// Also do note that this does not go through the reflector.

	letterSignal := letterID
	for i := 0; i < len(rotorArray); i++ {
		// Adjust the rotor position to consider rotorSpinOffset
		adjustedSignal := addWithOverflow(letterSignal, rotorArray[i].rotorSpinOffset-1, 26)
		fmt.Printf("    ROTOR #%d: Input is letter at #%d plus offset %d--> Really: %d\n", i, letterSignal, rotorArray[i].rotorSpinOffset-1, adjustedSignal)
		letterSignal = rotorArray[i].rotorMap.forward[adjustedSignal]
		fmt.Printf("    ROTOR #%d : Output is letter #%d\n", i, letterSignal)

		// Check for rotor rotation
		if i == len(rotorArray)-1 {
			rotorArray[i].rotate()
		}
	}
	fmt.Println()
	return letterSignal
}

func getThroughRotorsB(rotorArray [3]*rotor, letterID int) int {
	// This function only goes through the RotorArray Last to First
	// Also do note that this does not go through the reflector.

	letterSignal := letterID
	for i := len(rotorArray) - 1; i >= 0; i-- {
		// Adjust the rotor position to consider rotorSpinOffset
		adjustedSignal := addWithOverflow(letterSignal, -(rotorArray[i].rotorSpinOffset - 1), 26)
		fmt.Printf("    ROTOR #%d: Input is letter at #%d plus offset %d--> Really: %d\n", i, letterSignal, rotorArray[i].rotorSpinOffset-1, adjustedSignal)
		letterSignal = rotorArray[i].rotorMap.backward[adjustedSignal]
		fmt.Printf("    ROTOR #%d : Output is letter #%d\n", i, letterSignal)

		// Check for rotor rotation
		if i == 0 {
			rotorArray[i].rotate()
		}
	}
	fmt.Println()
	return letterSignal
}

// Add a rotate method to the rotor struct
func (r *rotor) rotate() {
	r.rotorSpinOffset = addWithOverflow(r.rotorSpinOffset, 1, 26)
	r.nextRotorSpin = addWithOverflow(r.nextRotorSpin, 1, 26)
}

func getThroughPlugboardF(letterMap *obfuscatorMap, letterIn int) int {
	// This function returns a letter num only if present in the provided 'plugboard'
	// Otherwise returns the same letter as provided since nothing is plugged into our hypothetical ENIGMA plugboard
	if mappedLetter, ok := letterMap.forward[letterIn]; ok {
		return mappedLetter // Return the mapped letter if present
	}
	return letterIn // Return the same letter if not present in the plugboard
}

func getThroughPlugboardB(letterMap *obfuscatorMap, letterIn int) int {
	// This function returns a letter num only if present in the provided 'plugboard'
	// Otherwise returns the same letter as provided since nothing is plugged into our hypothetical ENIGMA plugboard
	if mappedLetter, ok := letterMap.backward[letterIn]; ok {
		return mappedLetter // Return the mapped letter if present
	}
	return letterIn // Return the same letter if not present in the plugboard
}

func getThroughReflector(givenReflector map[int]int, rotorOut int) int {
	// Reflector is just a limited rotor, but with a limitation...
	// No letter can map to itself, a cryptographic weakness caused by the same wires being used for forwards and backwards legs.
	return givenReflector[rotorOut]
}
