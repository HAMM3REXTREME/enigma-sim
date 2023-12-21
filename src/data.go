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

func (r *rotor) signalLetterF(letterID int) int {
	// Pass a letter signal forward to backward considering spin offset.
	// Off-by-one errors can be really hard to diagnose here.

	// letterID is is the static contact pin for that letter.
	// offsetSignal is the actual contact pin activated on the rotor due to its spin offset angle. Does roll over if 0>=x>26
	// finalSignal is what comes out of the rotor + corrected for spin offset since both sides of the rotor move when spun.
	offsetSignal := addWithOverflow(letterID, r.rotorSpinOffset-1, 26)
	finalSignal := addWithOverflow(r.rotorMap.forward[offsetSignal], r.rotorSpinOffset-1, 26)
	//fmt.Printf("    ROTOR: Input is %d which goes f(%d)+%d to get output %d.\n", letterID, offsetSignal, r.rotorSpinOffset-1, finalSignal)
	return finalSignal
}

func (r *rotor) signalLetterB(letterID int) int {
	// Pass a letter signal backward to forward considering spin offset.
	// Off-by-one errors can be really hard to diagnose here.

	// letterID is is the static contact pin for that letter.
	// offsetSignal is the actual contact pin activated on the rotor due to its spin offset angle. Does roll over if 0>=x>26
	// finalSignal is what comes out of the rotor + corrected for spin offset since both sides of the rotor move when spun.
	offsetSignal := addWithOverflow(letterID, -(r.rotorSpinOffset - 1), 26)
	finalSignal := addWithOverflow(r.rotorMap.backward[offsetSignal], -(r.rotorSpinOffset - 1), 26)
	//fmt.Printf("    ROTOR: Input is %d which goes b(%d)-%d to get output %d.\n", letterID, offsetSignal, r.rotorSpinOffset-1, finalSignal)
	return finalSignal
}
