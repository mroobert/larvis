// Package tiebreak provides support for tie-breaking the hands
// when both hands are of the same rank.
package tiebreak

import (
	"errors"
	"fmt"

	"github.com/mroobert/larvis"
)

var (
	ErrTieBreakerNotFound = errors.New("no tie-breaker found for rank")
	ErrTieBreakersNil     = errors.New("tie-breakers not initialized")
)

// tieBreaker compares two hands of the same rank.
type tieBreaker interface {
	compare(hand1, hand2 larvis.Hand) string
}

// Decider manages the tie-breaking of hands.
type Decider struct {
	tieBreakers map[larvis.HandRank]tieBreaker
}

func NewDecider() Decider {
	return Decider{
		tieBreakers: map[larvis.HandRank]tieBreaker{
			larvis.HighCardRank:    highCardTieBreaker{},
			larvis.OnePairRank:     onePairTieBreaker{},
			larvis.TwoPairRank:     twoPairTieBreaker{},
			larvis.TripleRank:      tripleTieBreaker{},
			larvis.FullHouseRank:   fullHouseTieBreaker{},
			larvis.FourOfAKindRank: fourOfAKindTieBreaker{},
		},
	}
}

// ApplyTieBreak compares two hands of the same rank based on
// specific tie-breaking rules associated with the rank.
//
// If a winner cannot be decided, we have a tie.
func (d Decider) ApplyTieBreak(r larvis.HandRank, hand1, hand2 larvis.Hand) (string, error) {
	if d.tieBreakers == nil {
		return "", ErrTieBreakersNil
	}

	t, ok := d.tieBreakers[r]
	if !ok {
		return "", fmt.Errorf("%w %q", ErrTieBreakerNotFound, r)
	}
	return t.compare(hand1, hand2), nil
}
