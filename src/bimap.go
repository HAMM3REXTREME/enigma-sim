package main

func newBimap(forward map[int]int) *biMap {
	// Takes a normal map and returns a bimap
	backward := make(map[int]int)
	for k, v := range forward {
		backward[v] = k
	}
	return &biMap{
		forward:  forward,
		backward: backward,
	}
}

// BiMaps form 1:1 connections (basically like plugboards)
func (letterMap *biMap) throughMapF(letterIn int) int {
	if mappedLetter, ok := letterMap.forward[letterIn]; ok { // Front to back map
		return mappedLetter // Return the mapped letter if present (plugged in)
	}
	return letterIn // Return the same letter if not present in the plugboard
}

func (letterMap *biMap) throughMapB(letterIn int) int {
	if mappedLetter, ok := letterMap.backward[letterIn]; ok { // Back to front map
		return mappedLetter // Return the mapped letter if present (plugged in)
	}
	return letterIn // Return the same letter if not present in the plugboard
}
