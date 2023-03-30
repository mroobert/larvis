package rank

import (
	"github.com/mroobert/larvis/game"
	"github.com/mroobert/larvis/rank/freq"
)

// fullHouseMatcher determines the match for FULL HOUSE rank.
type fullHouseMatcher struct{}

func (fullHouseMatcher) match(f map[game.Card]int) bool {
	return freq.HasFreq(f, 3) && freq.HasFreq(f, 2)
}

func (fullHouseMatcher) rank() game.HandRank {
	return game.FullHouseRank
}
