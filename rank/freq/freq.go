// Package freq provides support for frequency analysis of cards in a hand.
package freq

import "github.com/mroobert/larvis"

// CardsFreq returns a map with frequency of each card in the hand.
func CardsFreq(h larvis.Hand) map[larvis.Card]int {
	freq := make(map[larvis.Card]int)
	for _, c := range h {
		freq[c] = freq[c] + 1
	}

	return freq
}

// HasFreq returns true if the frequency map has a card with frequency n.
// This means that the hand contains a card n times.
func HasFreq(f map[larvis.Card]int, n int) bool {
	for _, v := range f {
		if v == n {
			return true
		}
	}
	return false
}

// HasPairs returns true if the frequency map has n pairs.
// This means that the hand contains n pairs.
func HasPairs(f map[larvis.Card]int, n int) bool {
	p := 0
	for k := range f {
		if f[k] == 2 {
			p++
		}
	}
	return p == n
}
