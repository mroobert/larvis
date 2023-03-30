// Package rank provides support for ranking the hands.
package rank

import (
	"github.com/mroobert/larvis/game"
	"github.com/mroobert/larvis/rank/freq"
)

// rankMatcher evaluates the frequency of cards in a hand
// to determine if the hand is of a particular rank.
type rankMatcher interface {
	match(f map[game.Card]int) bool
	rank() game.HandRank
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

// RankHand ranks the hand based on the frequency of cards in the hand.
func (r Ranker) RankHand(h game.Hand) game.HandRank {
	f := freq.CardsFreq(h)
	for _, evaluator := range r.matchers {
		if evaluator.match(f) {
			return evaluator.rank()
		}
	}

	return game.HighCardRank
}
