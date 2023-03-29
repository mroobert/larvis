package larvis

// cardsFreq returns a map of cards and their frequency.
func cardsFreq(h Hand) map[card]int {
	freq := make(map[card]int)
	for _, c := range h {
		freq[c] = freq[c] + 1
	}

	return freq
}
