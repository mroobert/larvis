package tiebreak

import (
	"sort"

	"github.com/mroobert/larvis"
	"github.com/mroobert/larvis/rank/freq"
)

// twoPairTieBreaker manages the tie-breaking of hands with the TWO PAIR rank.
type twoPairTieBreaker struct{}

// compare compares the two pairs in each hand. The highest pair is compared first.
// If the two pairs are the same, the remaining card is compared.
func (t twoPairTieBreaker) compare(hand1, hand2 larvis.Hand) string {
	pairsHand1, remainingCard1 := t.findTwoPair(hand1)
	pairsHand2, remainingCard2 := t.findTwoPair(hand2)

	// pairs comparison
	sort.Slice(pairsHand1, func(i, j int) bool {
		return pairsHand1[i].Value > pairsHand1[j].Value // sort descending
	})
	sort.Slice(pairsHand2, func(i, j int) bool {
		return pairsHand2[i].Value > pairsHand2[j].Value // sort descending
	})
	if pairsHand1[0].Value > pairsHand2[0].Value {
		return larvis.Hand1Wins
	}
	if pairsHand1[0].Value < pairsHand2[0].Value {
		return larvis.Hand2Wins
	}
	if pairsHand1[1].Value > pairsHand2[1].Value {
		return larvis.Hand1Wins
	}
	if pairsHand1[1].Value < pairsHand2[1].Value {
		return larvis.Hand2Wins
	}

	// remaining larvis.Card comparison
	if remainingCard1.Value > remainingCard2.Value {
		return larvis.Hand1Wins
	}
	if remainingCard1.Value < remainingCard2.Value {
		return larvis.Hand2Wins
	}

	return larvis.Tie
}

// findTwoPair returns the two pairs and the remaining card.
func (twoPairTieBreaker) findTwoPair(hand larvis.Hand) (pairs []larvis.Card, remainingCard larvis.Card) {
	f := freq.CardsFreq(hand)
	for k := range f {
		if f[k] == 2 {
			pairs = append(pairs, k)
		} else {
			remainingCard = k
		}
	}
	return
}
