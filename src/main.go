package main

import "fmt"

func main() {
	plugboard, err := runes2biMap([]rune("DMTWSILRUYQNKFEJCAZBPGXOHV"))
	if err != nil {
		fmt.Printf("Error creating plugboard: %s\n", err)
		return
	}

	reflector, err := runes2map([]rune("BADCFEHGJILKNMPORQTSVUXWZY"))
	if err != nil {
		fmt.Printf("Error creating reflector: %s\n", err)
		return
	}

	rotor1, err := newRotor([]rune("DMTWSILRUYQNKFEJCAZBPGXOHV"), 9, 5)
	if err != nil {
		fmt.Printf("Error creating rotor1: %s\n", err)
		return
	}

	rotor2, err := newRotor([]rune("HQZGPJTMOBLNCIFDYAWVEUSRKX"), 9, 10)
	if err != nil {
		fmt.Printf("Error creating rotor2: %s\n", err)
		return
	}

	rotor3, err := newRotor([]rune("UQNTLSZFMREHDPXKIBVYGJCWOA"), 3, 0)
	if err != nil {
		fmt.Printf("Error creating rotor3: %s\n", err)
		return
	}

	rotorArray := []*rotor{rotor1, rotor2, rotor3}

	userMachine, err := newEnigmaMachine(plugboard, rotorArray, reflector)
	if err != nil {
		fmt.Printf("Error creating ENIGMA machine. How though, seeing we passed the previous check?!: %s\n", err)
		return
	}

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
