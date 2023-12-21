package main

import "fmt"

func debugObfuscateFull(plugboard *obfuscatorMap, rotors []*rotor, strList []string, reflectorMap map[int]int) {
	for i := 0; i < len(strList); i++ {
		char := strList[i]
		num, _ := char2num(char)
		incrementRotors(rotors)
		fmt.Printf("INPUT: Letter #%d is '%s' \n", num, char)

		fmt.Printf("Plugboard forward %d-->", num)
		num = plugboard.throughMapF(num)
		char, _ = num2char(num)
		fmt.Printf("%d which is %s\n", num, char)

		fmt.Printf("Rotors forward %d-->", num)

		num = throughRotorsF(rotors, num)
		char, _ = num2char(num)
		fmt.Printf("%d which is %s\n", num, char)

		fmt.Printf("Reflector %d-->", num)
		num = throughReflector(reflectorMap, num)
		char, _ = num2char(num)
		fmt.Printf("%d which is %s\n", num, char)

		fmt.Printf("Rotors reverse %d-->", num)

		num = throughRotorsB(rotors, num)
		char, _ = num2char(num)
		fmt.Printf("%d which is %s\n", num, char)

		fmt.Printf("Plugboard reverse %d-->", num)
		num = plugboard.throughMapB(num)
		char, _ = num2char(num)
		fmt.Printf("%d which is %s\n", num, char)

		fmt.Printf("FINAL: Encrypted letter is #%d which is '%s'\n", num, char)
		fmt.Println()

	}
}
