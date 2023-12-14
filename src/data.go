package main

type rotor struct {
	rotorMapF           map[int]int // Contains Rotor BiMap
	rotorMapB           map[int]int // TODO: Make a proper Bimap structure
	rotorSpinOffset     int         // Offset, or the spin of the rotor.
	nextRotorSpinLetter int         // Determines whether to increment the next rotor's spin
}

type plugboard struct {
	PlugboardMapF map[int]int // Plugboard BiMap, because a plugboard makes 1:1 connections
	PlugboardMapB map[int]int
}
