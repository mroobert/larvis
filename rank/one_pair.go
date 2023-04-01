package rank

import (
	"github.com/mroobert/larvis"
	"github.com/mroobert/larvis/rank/freq"
)

// onePairMatcher determines the match for ONE PAIR rank.
type onePairMatcher struct{}

func (onePairMatcher) match(f map[larvis.Card]int) bool {
	return freq.HasPairs(f, 1) && !freq.HasFreq(f, 3)
}

func (onePairMatcher) rank() larvis.HandRank {
	return larvis.OnePairRank
}
