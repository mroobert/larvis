package rank

import (
	"github.com/mroobert/larvis/game"
	"github.com/mroobert/larvis/rank/freq"
)

// twoPairMatcher determines the match for TWO PAIR rank.
type twoPairMatcher struct{}

func (twoPairMatcher) match(f map[game.Card]int) bool {
	return freq.HasPairs(f, 2)
}

func (twoPairMatcher) rank() game.HandRank {
	return game.TwoPairRank
}
