package main

import "fmt"

func debugObfuscateFull(plugboard *obfuscatorMap, rotors [3]*rotor, strList []string, reflectorMap map[int]int) {
	for i := 0; i < len(strList); i++ {
		char := strList[i]
		num, _ := getLetterNumberByChar(char)
		incrementRotors(rotors)
		fmt.Printf("INPUT: Letter #%d is '%s' \n", num, char)

		fmt.Printf("Plugboard forward %d-->", num)
		num = getThroughPlugboardF(plugboard, num)
		char, _ = getCharByNumber(num)
		fmt.Printf("%d which is %s\n", num, char)

		fmt.Printf("Rotors forward %d-->", num)

		num = getThroughRotorsF(rotors, num)
		char, _ = getCharByNumber(num)
		fmt.Printf("%d which is %s\n", num, char)

		fmt.Printf("Reflector %d-->", num)
		num = getThroughReflector(reflectorMap, num)
		char, _ = getCharByNumber(num)
		fmt.Printf("%d which is %s\n", num, char)

		fmt.Printf("Rotors reverse %d-->", num)

		num = getThroughRotorsB(rotors, num)
		char, _ = getCharByNumber(num)
		fmt.Printf("%d which is %s\n", num, char)

		fmt.Printf("Plugboard reverse %d-->", num)
		num = getThroughPlugboardB(plugboard, num)
		char, _ = getCharByNumber(num)
		fmt.Printf("%d which is %s\n", num, char)

		fmt.Printf("FINAL: Encrypted letter is #%d which is '%s'\n", num, char)
		fmt.Println()

	}
}
