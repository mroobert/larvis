// Package tiebreak provides support for tie-breaking the hands
// when both hands are of the same rank.
package tiebreak

import (
	"fmt"

	"github.com/mroobert/larvis/game"
)

var (
	ErrTieBreakerNotFound = fmt.Errorf("no tie-breaker found for rank")
	ErrTieBreakersNil     = fmt.Errorf("tie-breakers not initialized")
)

// tieBreaker compares two hands of the same rank.
type tieBreaker interface {
	compare(hand1, hand2 game.Hand) string
}

// Decider manages the tie-breaking of hands.
type Decider struct {
	tieBreakers map[game.HandRank]tieBreaker
}

func NewDecider() Decider {
	return Decider{
		tieBreakers: map[game.HandRank]tieBreaker{
			game.HighCardRank:    highCardTieBreaker{},
			game.OnePairRank:     onePairTieBreaker{},
			game.TwoPairRank:     twoPairTieBreaker{},
			game.TripleRank:      tripleTieBreaker{},
			game.FullHouseRank:   fullHouseTieBreaker{},
			game.FourOfAKindRank: fourOfAKindTieBreaker{},
		},
	}
}

// ApplyTieBreak compares two hands of the same rank based on
// specific tie-breaking rules associated with the rank.
//
// If a winner cannot be decided, we have a tie.
func (d Decider) ApplyTieBreak(r game.HandRank, hand1, hand2 game.Hand) (string, error) {
	if d.tieBreakers == nil {
		return "", ErrTieBreakersNil
	}

	t, ok := d.tieBreakers[r]
	if !ok {
		return "", fmt.Errorf("%w %q", ErrTieBreakerNotFound, r)
	}
	return t.compare(hand1, hand2), nil
}