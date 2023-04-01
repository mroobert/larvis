package rank

import (
	"github.com/mroobert/larvis"
	"github.com/mroobert/larvis/rank/freq"
)

// fourOfAKindMatcher determines the match for FOUR OF A KIND rank.
type fourOfAKindMatcher struct{}

func (fourOfAKindMatcher) match(f map[larvis.Card]int) bool {
	return freq.HasFreq(f, 4)
}

func (fourOfAKindMatcher) rank() larvis.HandRank {
	return larvis.FourOfAKindRank
}
