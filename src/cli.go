package main

import (
	"bufio"
	"fmt"
	"os"
)

func scanLine() string {
	// Scans a line. Will count spaces.
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	//fmt.Println("Your input:", line)
	return line
}

func userEnigmaInput(machine *enigmaMachine) {
	fmt.Printf("Enter text into the ENIGMA machine: ")
	userInput := scanLine()
	enigmaOutput, err := encryptText(*machine, userInput, true)
	if err != nil {
		fmt.Printf("Sorry, some error occured.\n")
		return
	}
	fmt.Printf("ENIGMA output: %s\n", string(enigmaOutput))
}
