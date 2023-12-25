package main

// Rotor spinning functions:
// Pass a letter signal from one side to another considering spin offset.
// Off-by-one errors can be really hard to diagnose here.

// letterID is is the static contact pin for that letter.
// offsetSignal is the actual contact pin activated on the rotor due to its spin offset angle. Does roll over if 0>=x>26
// finalSignal is what comes out of the rotor + corrected for spin offset since both sides of the rotor move when spun.
func (r *rotor) throughRotorF(letterID int) int {
	offsetSignal := addWithOverflow(letterID, r.spin-1, Letters)
	finalSignal := addWithOverflow(r.mapping.forward[offsetSignal], r.spin-1, Letters)
	return finalSignal
}
func (r *rotor) throughRotorB(letterID int) int {
	offsetSignal := addWithOverflow(letterID, -(r.spin - 1), Letters)
	finalSignal := addWithOverflow(r.mapping.backward[offsetSignal], -(r.spin - 1), Letters)
	return finalSignal
}
