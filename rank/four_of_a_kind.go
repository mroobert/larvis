package rank

import (
	"github.com/mroobert/larvis/game"
	"github.com/mroobert/larvis/rank/freq"
)

// fourOfAKindMatcher determines the match for FOUR OF A KIND rank.
type fourOfAKindMatcher struct{}

func (fourOfAKindMatcher) match(f map[game.Card]int) bool {
	return freq.HasFreq(f, 4)
}

func (fourOfAKindMatcher) rank() game.HandRank {
	return game.FourOfAKindRank
}
