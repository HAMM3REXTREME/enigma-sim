package main

import "unicode"

func encryptText(plugboard *obfuscatorMap, rotors []*rotor, reflectorMap map[int]int, text string, matchCase bool) string {
	encryptedText := "" // Our new encrypted text

	for _, char := range text {
		if unicode.IsLetter(char) { // Only try to encrypt letters
			encryptedChar := encryptRune(plugboard, rotors, reflectorMap, char)

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

func encryptRune(plugboard *obfuscatorMap, rotors []*rotor, reflectorMap map[int]int, char rune) rune {
	// Takes a rune and outputs the encrypted rune. Pure function
	var newChar rune

	incrementRotors(rotors)
	num, _ := rune2num(char)

	num = plugboard.throughMapF(num)
	num = throughRotorArrayF(rotors, num)
	num = throughReflector(reflectorMap, num)
	num = throughRotorArrayB(rotors, num)
	num = plugboard.throughMapB(num)

	newChar, _ = num2rune(num)
	return newChar
}

func incrementRotors(rotorArray []*rotor) {
	// Increments rotors by modifying arg
	rotorArray[0].rotorSpinOffset = addWithOverflow(rotorArray[0].rotorSpinOffset, 1, Letters) // Increment the position of the first rotor

	// Check if the notch is reached for the next rotors and spin them accordingly
	for i := 1; i < len(rotorArray)-1; i++ {
		if rotorArray[i-1].rotorSpinOffset == rotorArray[i-1].nextRotorSpin {
			// Increment the current rotor
			//fmt.Printf("i: %d Incrementing rotor because prev. spin offset: %d\n", i, rotorArray[i-1].rotorSpinOffset)
			rotorArray[i].rotorSpinOffset = addWithOverflow(rotorArray[i].rotorSpinOffset, 1, Letters)
			// Check if the next rotor's notch is reached, then increment it
			if rotorArray[i].rotorSpinOffset == rotorArray[i].nextRotorSpin {
				//fmt.Printf("i: %d Incrementing next rotor because spin offset: %d matches nextRotorSpin: %d\n", i, rotorArray[i].rotorSpinOffset, rotorArray[i].nextRotorSpin)
				rotorArray[i+1].rotorSpinOffset = addWithOverflow(rotorArray[i+1].rotorSpinOffset, 1, Letters)
			}
		}
	}
}

func throughRotorArrayF(rotorArray []*rotor, letterID int) int {
	// This function only goes through the RotorArray First to Last
	// Also do note that this does not go through the reflector.

	letterSignal := letterID
	for i := 0; i < len(rotorArray); i++ {
		letterSignal = rotorArray[i].throughRotorF(letterSignal)

	}
	return letterSignal
}

func throughRotorArrayB(rotorArray []*rotor, letterID int) int {
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
