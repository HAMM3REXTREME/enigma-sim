package main

const Letters = 26 // The number of alphabets in the english language.

type biMap struct { // 2-Way aka 1:1 obfuscation device. Like a Plugboard.
	forward  map[int]int
	backward map[int]int
}

// Rotors are just fancy spinnable plugboards with notches.
type rotor struct {
	mapping *biMap // Rotor or Plugboard BiMap, because they both make 1:1 connections
	spin    int    // Basically the offset, or the spin of the rotor.
	notch   int    // Determines whether to increment the next rotor's spin, basically the 'notch'.
}

type enigmaMachine struct {
	plugboard  *biMap
	rotorArray []*rotor
	reflector  map[int]int
}
