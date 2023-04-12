// Package rank provides support for ranking the hands.
package rank

import (
	"github.com/mroobert/larvis"
	"github.com/mroobert/larvis/rank/freq"
)

// rankMatcher evaluates the frequency of cards in hand
// to match a rank.
type rankMatcher interface {
	match(f map[larvis.Card]int) bool
	rank() larvis.HandRank
}

// Ranker manages the ranking of hands.
type Ranker struct {
	matchers []rankMatcher
}

func NewRanker() Ranker {
	matchers := []rankMatcher{
		onePairMatcher{},
		twoPairMatcher{},
		tripleMatcher{},
		fullHouseMatcher{},
		fourOfAKindMatcher{},
	}

	return Ranker{
		matchers: matchers,
	}
}

// RankHand determines the rank based on the frequency of cards in hand.
func (r Ranker) RankHand(h larvis.Hand) larvis.HandRank {
	f := freq.CardsFreq(h)
	for _, m := range r.matchers {
		if m.match(f) {
			return m.rank()
		}
	}

	return larvis.HighCardRank
}
