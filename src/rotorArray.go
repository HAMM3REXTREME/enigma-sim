package main

func incrementRotors(rotorArray []*rotor) {
	// Increments rotors when notch aligns.
	rotorArray[0].spin = addWithOverflow(rotorArray[0].spin, 1, Letters)

	for i := 1; i < len(rotorArray)-1; i++ {
		if rotorArray[i-1].spin == rotorArray[i-1].notch { // If previous rotor's notch is matching,
			rotorArray[i].spin = addWithOverflow(rotorArray[i].spin, 1, Letters) // Increment this rotor then
			if rotorArray[i].spin == rotorArray[i].notch {                       // and if this rotor's notch matches, increment the next rotor too
				rotorArray[i+1].spin = addWithOverflow(rotorArray[i+1].spin, 1, Letters)
			}
		}
	}
}

// These 2 functions go through the whole array of rotors
func throughRotorArrayF(rotorArray []*rotor, letterID int) int {
	// This only goes through the RotorArray First to Last, does not go through the reflector.
	letterSignal := letterID
	for i := 0; i < len(rotorArray); i++ {
		letterSignal = rotorArray[i].throughRotorF(letterSignal)
	}
	return letterSignal
}
func throughRotorArrayB(rotorArray []*rotor, letterID int) int {
	// This only goes through the RotorArray Last to First, does not go through the reflector.
	letterSignal := letterID
	for i := len(rotorArray) - 1; i >= 0; i-- {
		letterSignal = rotorArray[i].throughRotorB(letterSignal)
	}
	return letterSignal
}

func throughReflector(reflector map[int]int, rotorOut int) int {
	// Reflector is kinda a non-spinning rotor (but only has one side, same wires used for fw and bw legs so no letter can map to itself)
	return reflector[rotorOut] // Make sure your reflector slice is physically possible to avoid bugs
}
