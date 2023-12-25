package main

import "fmt"

// -------------------Rotor spinning functions:-------------------------------------------------------------------------------------
// Pass a letter signal from one side to another considering spin offset.
// Off-by-one errors can be really hard to diagnose here.
// ---------------------------------------------------------------------------------------------------------------------------------
// letterID is is the static contact pin for that letter. (Input)
// offsetSignal is the actual contact pin activated on the rotor due to its spin offset angle. Does roll over if 0>=x>26
// finalSignal is what comes out of the rotor + corrected for spin offset since both sides of the rotor move when spun. (Output)
func (r *rotor) throughRotorF(letterID int) int { // Front to back
	offsetSignal := addWithOverflow(letterID, r.spin-1, Letters)
	finalSignal := addWithOverflow(r.mapping.forward[offsetSignal], r.spin-1, Letters)
	return finalSignal
}
func (r *rotor) throughRotorB(letterID int) int { // Back to front
	offsetSignal := addWithOverflow(letterID, -(r.spin - 1), Letters)
	finalSignal := addWithOverflow(r.mapping.backward[offsetSignal], -(r.spin - 1), Letters)
	return finalSignal
}

// Just a convenient way to make a new rotor
func newRotor(mapping []rune, spinOffset int, notch int) (*rotor, error) {
	realMap, err := runes2biMap(mapping)
	if err != nil {
		return nil, fmt.Errorf("error generating rotor (bi)mapping: %s", err)
	}
	finalRotor := &rotor{
		mapping: realMap,
		spin:    spinOffset,
		notch:   notch,
	}
	return finalRotor, nil
}
