package main

// 2-Way aka 1:1 obfuscation device. Like a Plugboard.
type obfuscatorMap struct {
	forward  map[int]int
	backward map[int]int
}

func newBimap(forward map[int]int) *obfuscatorMap {
	backward := make(map[int]int)
	for k, v := range forward {
		backward[v] = k
	}
	return &obfuscatorMap{
		forward:  forward,
		backward: backward,
	}
}

// Rotors are just fancy spinnable plugboards with notches.
type rotor struct {
	rotorMap        *obfuscatorMap // Rotor or Plugboard BiMap, because they both make 1:1 connections
	rotorSpinOffset int            // Basically the offset, or the spin of the rotor.
	nextRotorSpin   int            // Determines whether to increment the next rotor's spin, basically the 'notch'.
}
