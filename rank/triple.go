package rank

import (
	"github.com/mroobert/larvis/game"
	"github.com/mroobert/larvis/rank/freq"
)

// tripleMatcher determines the match for TRIPLE rank.
type tripleMatcher struct{}

func (tripleMatcher) match(f map[game.Card]int) bool {
	return freq.HasFreq(f, 3) && !freq.HasFreq(f, 2)
}

func (tripleMatcher) rank() game.HandRank {
	return game.TripleRank
}
