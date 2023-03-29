package larvis

// Rank represents the rank of a hand.
type rank int

// The available ranks from lowest to highest.
const (
	highCardRank rank = iota + 1
	onePairRank
	twoPairRank
	tripleRank
	fullHouseRank
	fourOfAKindRank
)

// Ranker manages the ranking of a hand.
type ranker struct{}

// RankHand determines the rank of a hand.
func (r ranker) rankHand(h Hand) rank {
	freq := cardsFreq(h)

	switch {
	case r.fourOfAKind(freq):
		return fourOfAKindRank
	case r.fullHouse(freq):
		return fullHouseRank
	case r.triple(freq):
		return tripleRank
	case r.twoPairs(freq):
		return twoPairRank
	case r.pair(freq):
		return onePairRank
	default:
		return highCardRank
	}
}

// fourOfAKind returns true if the hand has four of a kind.
func (r ranker) fourOfAKind(freq map[card]int) bool {
	return r.hasFreq(freq, 4)
}

// fullHouse returns true if the hand has three of a kind and a two of a kind.
func (r ranker) fullHouse(freq map[card]int) bool {
	return r.hasFreq(freq, 3) && r.hasFreq(freq, 2)
}

// triple returns true if the hand has three of a kind and no two of a kind.
func (r ranker) triple(freq map[card]int) bool {
	return r.hasFreq(freq, 3) && !r.hasFreq(freq, 2)
}

// twoPairs returns true if the hand has two pairs.
func (r ranker) twoPairs(freq map[card]int) bool {
	return r.hasPairs(freq, 2)
}

// pair returns true if the hand has a pair and no three of a kind.
func (r ranker) pair(freq map[card]int) bool {
	return r.hasPairs(freq, 1) && !r.hasFreq(freq, 3)
}

// hasFreq returns true if the hand has a card with the given frequency.
func (r ranker) hasFreq(freq map[card]int, n int) bool {
	for _, f := range freq {
		if f == n {
			return true
		}
	}
	return false
}

// hasPairs returns true if the hand has n pairs.
func (r ranker) hasPairs(freq map[card]int, n int) bool {
	p := 0
	for k := range freq {
		if freq[k] == 2 {
			p++
		}
	}
	return p == n
}
