package rank

import (
	"github.com/mroobert/larvis"
	"github.com/mroobert/larvis/rank/freq"
)

// tripleMatcher determines the match for TRIPLE rank.
type tripleMatcher struct{}

func (tripleMatcher) match(f map[larvis.Card]int) bool {
	return freq.HasFreq(f, 3) && !freq.HasFreq(f, 2)
}

func (tripleMatcher) rank() larvis.HandRank {
	return larvis.TripleRank
}
