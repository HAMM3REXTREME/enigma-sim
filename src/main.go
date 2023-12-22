package main

func main() {
	plugboardData := map[int]int{
		1: 5, 2: 10, 3: 15, 4: 20, 5: 1,
		6: 6, 7: 11, 8: 16, 9: 21, 10: 2,
		11: 7, 12: 12, 13: 17, 14: 22, 15: 3,
		16: 8, 17: 13, 18: 18, 19: 23, 20: 4,
		21: 9, 22: 14, 23: 19, 24: 24, 25: 25, 26: 26,
	}
	plugboard := newBimap(plugboardData)
	// Reflector: If x maps to y then y must map to x else you gonna get bugs
	reflector := map[int]int{
		1: 2, 2: 1, 3: 4, 4: 3, 5: 6,
		6: 5, 7: 8, 8: 7, 9: 10, 10: 9,
		11: 12, 12: 11, 13: 14, 14: 13, 15: 16,
		16: 15, 17: 18, 18: 17, 19: 20, 20: 19,
		21: 22, 22: 21, 23: 24, 24: 23, 25: 26, 26: 25,
	}
	rotor1 := newRotor([]rune("DMTWSILRUYQNKFEJCAZBPGXOHV"), 9, 5)
	rotor2 := newRotor([]rune("HQZGPJTMOBLNCIFDYAWVEUSRKX"), 9, 10)
	rotor3 := newRotor([]rune("UQNTLSZFMREHDPXKIBVYGJCWOA"), 3, 0)
	rotorArray := []*rotor{rotor1, rotor2, rotor3}
	userMachine := newEnigmaMachine(plugboard, rotorArray, reflector)

	for {
		userEnigmaInput(userMachine)
	}

}

/*
ENIGMA Machine Diagram:

               Bulbs
               /|\
                |
Keyboard-----> Plugboard <-----> Rotor-1 <-----> Rotor-2 <-----> Rotor-3 <-----> Reflector

*/
