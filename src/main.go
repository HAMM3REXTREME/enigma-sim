package main

import "fmt"

func main() {
	plugboard := runes2biMap([]rune("DMTWSILRUYQNKFEJCAZBPGXOHV"))
	reflector := runes2map([]rune("BADCFEHGJILKNMPORQTSVUXWZY"))

	if validateMap(reflector, true) != nil {
		fmt.Println("ENIGMA Error")
		return
	}

	rotor1 := &rotor{mapping: runes2biMap([]rune("DMTWSILRUYQNKFEJCAZBPGXOHV")), spin: 9, notch: 5}
	rotor2 := &rotor{mapping: runes2biMap([]rune("HQZGPJTMOBLNCIFDYAWVEUSRKX")), spin: 9, notch: 10}
	rotor3 := &rotor{mapping: runes2biMap([]rune("UQNTLSZFMREHDPXKIBVYGJCWOA")), spin: 3, notch: 0}
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
