package tiebreak

import (
	"sort"

	"github.com/mroobert/larvis/game"
	"github.com/mroobert/larvis/rank/freq"
)

// tripleTieBreaker manages the tie-breaking of hands with the TRIPLE rank.
type tripleTieBreaker struct{}

// compare compares the triple in each hand. If the triple is the same, the
// remaining cards are compared.
func (t tripleTieBreaker) compare(hand1, hand2 game.Hand) string {
	tripleHand1, remainingCards1 := t.findTriple(hand1)
	tripleHand2, remainingCards2 := t.findTriple(hand2)

	// triple comparison
	if tripleHand1.Value > tripleHand2.Value {
		return game.Hand1Wins
	}
	if tripleHand1.Value < tripleHand2.Value {
		return game.Hand2Wins
	}

	// remaining cards comparison
	sort.Slice(remainingCards1, func(i, j int) bool {
		return remainingCards1[i].Value > remainingCards1[j].Value // sort descending
	})
	sort.Slice(remainingCards2, func(i, j int) bool {
		return remainingCards2[i].Value > remainingCards2[j].Value // sort descending
	})
	if remainingCards1[0].Value > remainingCards2[0].Value {
		return game.Hand1Wins
	}
	if remainingCards1[0].Value < remainingCards2[0].Value {
		return game.Hand2Wins
	}
	if remainingCards1[1].Value > remainingCards2[1].Value {
		return game.Hand1Wins
	}
	if remainingCards1[1].Value < remainingCards2[1].Value {
		return game.Hand2Wins
	}

	return game.Tie
}

// findTriple returns the triple and the remaining cards.
func (tripleTieBreaker) findTriple(hand game.Hand) (triple game.Card, remainingCards []game.Card) {
	f := freq.CardsFreq(hand)
	for k := range f {
		if f[k] == 3 {
			triple = k
		} else {
			remainingCards = append(remainingCards, k)
		}
	}
	return
}
