package rank

import (
	"github.com/mroobert/larvis"
	"github.com/mroobert/larvis/rank/freq"
)

// fullHouseMatcher determines the match for FULL HOUSE rank.
type fullHouseMatcher struct{}

func (fullHouseMatcher) match(f map[larvis.Card]int) bool {
	return freq.HasFreq(f, 3) && freq.HasFreq(f, 2)
}

func (fullHouseMatcher) rank() larvis.HandRank {
	return larvis.FullHouseRank
}
