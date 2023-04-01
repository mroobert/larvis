package tiebreak

import (
	"sort"

	"github.com/mroobert/larvis"
	"github.com/mroobert/larvis/rank/freq"
)

// onePairTieBreaker manages the tie-breaking of hands with the ONE PAIR rank.
type onePairTieBreaker struct{}

// compare compares the pair in each hand.
// If the pair is the same, the remaining cards are compared.
func (o onePairTieBreaker) compare(hand1, hand2 larvis.Hand) string {
	pairHand1, remainingCards1 := o.findPair(hand1)
	pairHand2, remainingCards2 := o.findPair(hand2)

	// pair comparison
	if pairHand1.Value > pairHand2.Value {
		return larvis.Hand1Wins
	}
	if pairHand1.Value < pairHand2.Value {
		return larvis.Hand2Wins
	}

	// remaining cards comparison
	sort.Slice(remainingCards1, func(i, j int) bool {
		return remainingCards1[i].Value > remainingCards1[j].Value // sort descending
	})
	sort.Slice(remainingCards2, func(i, j int) bool {
		return remainingCards2[i].Value > remainingCards2[j].Value // sort descending
	})
	if remainingCards1[0].Value > remainingCards2[0].Value {
		return larvis.Hand1Wins
	}
	if remainingCards1[0].Value < remainingCards2[0].Value {
		return larvis.Hand2Wins
	}
	if remainingCards1[1].Value > remainingCards2[1].Value {
		return larvis.Hand1Wins
	}
	if remainingCards1[1].Value < remainingCards2[1].Value {
		return larvis.Hand2Wins
	}
	if remainingCards1[2].Value > remainingCards2[2].Value {
		return larvis.Hand1Wins
	}
	if remainingCards1[2].Value < remainingCards2[2].Value {
		return larvis.Hand2Wins
	}

	return larvis.Tie
}

// findPair returns the pair and the remaining cards.
func (onePairTieBreaker) findPair(hand larvis.Hand) (pair larvis.Card, remainingCards []larvis.Card) {
	f := freq.CardsFreq(hand)
	for k := range f {
		if f[k] == 2 {
			pair = k
		} else {
			remainingCards = append(remainingCards, k)
		}
	}
	return
}
