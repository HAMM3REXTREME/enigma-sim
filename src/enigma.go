package main

func getFullEnigma(plugboard *obfuscatorMap, rotors []*rotor, strList []string, reflectorMap map[int]int) []string {
	enigmaList := make([]string, len(strList))
	for i := 0; i < len(strList); i++ {
		char := strList[i]
		incrementRotors(rotors)

		num, _ := getLetterNumberByChar(char)

		num = getThroughPlugboardF(plugboard, num)
		num = getThroughRotorsF(rotors, num)
		num = getThroughReflector(reflectorMap, num)
		num = getThroughRotorsB(rotors, num)
		num = getThroughPlugboardB(plugboard, num)

		char, _ = getCharByNumber(num)
		enigmaList[i] = char

	}
	return enigmaList
}

func incrementRotors(rotorArray []*rotor) {
	// Increment the position of the first rotor
	//fmt.Println("Increment 1st rotor.")
	rotorArray[0].rotorSpinOffset = addWithOverflow(rotorArray[0].rotorSpinOffset, 1, 26)

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

func getThroughRotorsF(rotorArray []*rotor, letterID int) int {
	// This function only goes through the RotorArray First to Last
	// Also do note that this does not go through the reflector.

	letterSignal := letterID
	for i := 0; i < len(rotorArray); i++ {
		letterSignal = rotorArray[i].letterF(letterSignal)

	}
	return letterSignal
}

func getThroughRotorsB(rotorArray []*rotor, letterID int) int {
	// This function only goes through the RotorArray Last to First
	// Also do note that this does not go through the reflector.

	letterSignal := letterID
	for i := len(rotorArray) - 1; i >= 0; i-- {
		letterSignal = rotorArray[i].letterB(letterSignal)

	}
	return letterSignal
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
	// Make sure your reflector slice is physically possible to avoid bugs
	return givenReflector[rotorOut]
}
