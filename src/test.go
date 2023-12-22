//go:build exclude

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// TEST ENCRYPTION

/*
ENIGMA Machine Diagram:

	Bulbs
	/|\
	 |

Keyboard-----> Plugboard <-----> Rotor-1 <-----> Rotor-2 <-----> Rotor-3 <-----> Reflector
*/

func mainTest() {
	/* 	 	// Test data for biMap (Plugboard)
	plugboardData := map[int]int{
		1: 5, 2: 10, 3: 15, 4: 20, 5: 1,
		6: 6, 7: 11, 8: 16, 9: 21, 10: 2,
		11: 7, 12: 12, 13: 17, 14: 22, 15: 3,
		16: 8, 17: 13, 18: 18, 19: 23, 20: 4,
		21: 9, 22: 14, 23: 19, 24: 24, 25: 25, 26: 26,
	}
	plugboard := newBimap(plugboardData)  */

	// Test data for rotor 1
	rotor1Data := map[int]int{
		1: 5, 2: 10, 3: 15, 4: 20, 5: 1,
		6: 6, 7: 11, 8: 16, 9: 21, 10: 2,
		11: 7, 12: 12, 13: 17, 14: 22, 15: 3,
		16: 8, 17: 13, 18: 18, 19: 23, 20: 4,
		21: 9, 22: 14, 23: 19, 24: 24, 25: 25, 26: 26,
	}
	rotor1 := &rotor{
		rotorMap:        newBimap(rotor1Data),
		rotorSpinOffset: 1,
		nextRotorSpin:   5,
	}

	// Test data for rotor 2
	rotor2Data := map[int]int{
		1: 10, 2: 15, 3: 20, 4: 1, 5: 6,
		6: 11, 7: 16, 8: 21, 9: 2, 10: 7,
		11: 12, 12: 17, 13: 22, 14: 3, 15: 8,
		16: 13, 17: 18, 18: 23, 19: 4, 20: 9,
		21: 14, 22: 19, 23: 24, 24: 25, 25: 26, 26: 5,
	}
	rotor2 := &rotor{
		rotorMap:        newBimap(rotor2Data),
		rotorSpinOffset: 1,
		nextRotorSpin:   10,
	}

	// Test data for rotor 3
	rotor3Data := map[int]int{
		1: 15, 2: 20, 3: 1, 4: 6, 5: 11,
		6: 16, 7: 21, 8: 2, 9: 7, 10: 12,
		11: 17, 12: 22, 13: 3, 14: 8, 15: 13,
		16: 18, 17: 23, 18: 4, 19: 9, 20: 14,
		21: 19, 22: 24, 23: 25, 24: 26, 25: 5, 26: 10,
	}
	rotor3 := &rotor{
		rotorMap:        newBimap(rotor3Data),
		rotorSpinOffset: 1,
		nextRotorSpin:   0,
	}
	rotorArray := []*rotor{rotor1, rotor2, rotor3}

	// Very important for this to correct
	/* 	// if x maps to y then y must map to x else you gonna get bugs
	   		reflector := map[int]int{
	   		1: 5, 2: 10, 3: 15, 4: 20, 5: 1,
	   		6: 6, 7: 11, 8: 16, 9: 21, 10: 2,
	   		11: 7, 12: 12, 13: 17, 14: 22, 15: 3,
	   		16: 8, 17: 13, 18: 18, 19: 23, 20: 4,
	   		21: 9, 22: 14, 23: 19, 24: 24, 25: 25, 26: 26,
	   	}
	*/
	for {
		test(rotorArray)
		//userInput(plugboard, rotorArray, reflector)
		//debugObfuscateFull(plugboard, rotorArray, []string{"H", "I", "L", "E", "R"}, reflector)
	}

}


func main() {
	// Test data for biMap (Plugboard)
	plugboardData := map[int]int{
		1: 5, 2: 10, 3: 15, 4: 20, 5: 1,
		6: 6, 7: 11, 8: 16, 9: 21, 10: 2,
		11: 7, 12: 12, 13: 17, 14: 22, 15: 3,
		16: 8, 17: 13, 18: 18, 19: 23, 20: 4,
		21: 9, 22: 14, 23: 19, 24: 24, 25: 25, 26: 26,
	}
	plugboard := newBimap(plugboardData)

	rotor1Data := map[int]int{
		1: 5, 2: 10, 3: 15, 4: 20, 5: 1,
		6: 6, 7: 11, 8: 16, 9: 21, 10: 2,
		11: 7, 12: 12, 13: 17, 14: 22, 15: 3,
		16: 8, 17: 13, 18: 18, 19: 23, 20: 4,
		21: 9, 22: 14, 23: 19, 24: 24, 25: 25, 26: 26,
	}
	rotor1 := &rotor{
		rotorMap:        newBimap(rotor1Data),
		rotorSpinOffset: 9,
		nextRotorSpin:   5,
	}

	rotor2Data := map[int]int{
		1: 10, 2: 15, 3: 20, 4: 1, 5: 6,
		6: 11, 7: 16, 8: 21, 9: 2, 10: 7,
		11: 12, 12: 17, 13: 22, 14: 3, 15: 8,
		16: 13, 17: 18, 18: 23, 19: 4, 20: 9,
		21: 14, 22: 19, 23: 24, 24: 25, 25: 26, 26: 5,
	}
	rotor2 := &rotor{
		rotorMap:        newBimap(rotor2Data),
		rotorSpinOffset: 9,
		nextRotorSpin:   10,
	}

	rotor3Data := map[int]int{
		1: 15, 2: 20, 3: 1, 4: 6, 5: 11,
		6: 16, 7: 21, 8: 2, 9: 7, 10: 12,
		11: 17, 12: 22, 13: 3, 14: 8, 15: 13,
		16: 18, 17: 23, 18: 4, 19: 9, 20: 14,
		21: 19, 22: 24, 23: 25, 24: 26, 25: 5, 26: 10,
	}
	rotor3 := &rotor{
		rotorMap:        newBimap(rotor3Data),
		rotorSpinOffset: 3,
		nextRotorSpin:   0,
	}
	//rotor1.rotorMap = runes2biMap([]rune("ASDFGHJKLQWERTYUIOPZXCVBNM"))
	rotorArray := []*rotor{rotor1, rotor2, rotor3}

	// Very important for this to be correct. If x maps to y then y must map to x else you gonna get bugs
	reflector := map[int]int{
		1: 2, 2: 1, 3: 4, 4: 3, 5: 6,
		6: 5, 7: 8, 8: 7, 9: 10, 10: 9,
		11: 12, 12: 11, 13: 14, 14: 13, 15: 16,
		16: 15, 17: 18, 18: 17, 19: 20, 20: 19,
		21: 22, 22: 21, 23: 24, 24: 23, 25: 26, 26: 25,
	}

	for {
		userEnigmaInput(plugboard, rotorArray, reflector)
		//debugObfuscateFull(plugboard, rotorArray, []string{"H", "I", "L", "E", "R"}, reflector)
	}

}


func testAdding() {
	// Passed
	rand.Seed(time.Now().UnixNano())
	a := rand.Intn(26) + 1
	b := rand.Intn(26) + 1
	ans := addWithOverflow(a, b, 26)
	fmt.Printf("Overflow test: %d + %d = %d...\n", a, b, ans)
	if ans <= 0 || ans > 26 {
		fmt.Printf("---------------------------SOME BUG DETECTED\n")
	}
}

func test(testArray []*rotor) {
	incrementRotors(testArray)
	//fmt.Printf("Rotors Dump: %s", testArray)
	for i := 0; i < len(testArray); i++ {
		fmt.Printf("ROTOR #%d Spin: %d    |    ", i+1, testArray[i].rotorSpinOffset)
	}
	fmt.Println()

}
