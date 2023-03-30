package rank

import (
	"github.com/mroobert/larvis/game"
	"github.com/mroobert/larvis/rank/freq"
)

// onePairMatcher determines the match for ONE PAIR rank.
type onePairMatcher struct{}

func (onePairMatcher) match(f map[game.Card]int) bool {
	return freq.HasPairs(f, 1) && !freq.HasFreq(f, 3)
}

func (onePairMatcher) rank() game.HandRank {
	return game.OnePairRank
}
