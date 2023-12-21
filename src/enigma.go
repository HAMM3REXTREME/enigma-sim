package main

func stringEnigma(plugboard *obfuscatorMap, rotors []*rotor, reflectorMap map[int]int, strList []string) []string {
	enigmaList := make([]string, len(strList))
	for i := 0; i < len(strList); i++ {
		char := strList[i]
		incrementRotors(rotors)

		num, _ := char2num(char)

		num = plugboard.throughMapF(num)
		num = throughRotorsF(rotors, num)
		num = throughReflector(reflectorMap, num)
		num = throughRotorsB(rotors, num)
		num = plugboard.throughMapB(num)

		char, _ = num2char(num)
		enigmaList[i] = char

	}
	return enigmaList
}

func incrementRotors(rotorArray []*rotor) {
	// Increments rotors by modifying arg
	rotorArray[0].rotorSpinOffset = addWithOverflow(rotorArray[0].rotorSpinOffset, 1, 26) // Increment the position of the first rotor

	// Check if the notch is reached for the next rotors and spin them accordingly
	for i := 1; i < len(rotorArray)-1; i++ {
		if rotorArray[i-1].rotorSpinOffset == rotorArray[i-1].nextRotorSpin {
			// Increment the current rotor
			//fmt.Printf("i: %d Incrementing rotor because prev. spin offset: %d\n", i, rotorArray[i-1].rotorSpinOffset)
			rotorArray[i].rotorSpinOffset = addWithOverflow(rotorArray[i].rotorSpinOffset, 1, 26)
			// Check if the next rotor's notch is reached, then increment it
			if rotorArray[i].rotorSpinOffset == rotorArray[i].nextRotorSpin {
				//fmt.Printf("i: %d Incrementing next rotor because spin offset: %d matches nextRotorSpin: %d\n", i, rotorArray[i].rotorSpinOffset, rotorArray[i].nextRotorSpin)
				rotorArray[i+1].rotorSpinOffset = addWithOverflow(rotorArray[i+1].rotorSpinOffset, 1, 26)
			}
		}
	}
}

func throughRotorsF(rotorArray []*rotor, letterID int) int {
	// This function only goes through the RotorArray First to Last
	// Also do note that this does not go through the reflector.

	letterSignal := letterID
	for i := 0; i < len(rotorArray); i++ {
		letterSignal = rotorArray[i].throughRotorF(letterSignal)

	}
	return letterSignal
}

func throughRotorsB(rotorArray []*rotor, letterID int) int {
	// This function only goes through the RotorArray Last to First
	// Also do note that this does not go through the reflector.

	letterSignal := letterID
	for i := len(rotorArray) - 1; i >= 0; i-- {
		letterSignal = rotorArray[i].throughRotorB(letterSignal)

	}
	return letterSignal
}

func throughReflector(givenReflector map[int]int, rotorOut int) int {
	// Reflector is just a limited rotor, but with a limitation...
	// No letter can map to itself, a cryptographic weakness caused by the same wires being used for forwards and backwards legs.
	// Make sure your reflector slice is physically possible to avoid bugs
	return givenReflector[rotorOut]
}
