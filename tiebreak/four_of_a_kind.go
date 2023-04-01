package tiebreak

import (
	"github.com/mroobert/larvis"
	"github.com/mroobert/larvis/rank/freq"
)

// fourOfAKindTieBreaker manages the tie-breaking of hands with the FOUR OF A KIND rank.
type fourOfAKindTieBreaker struct{}

// compare compares the FOUR OF A KIND in each hand.
// If the FOUR OF A KIND is the same, the remaining card is compared.
func (f fourOfAKindTieBreaker) compare(hand1, hand2 larvis.Hand) string {
	fourOfAKindHand1, remainingCard1 := f.findFourOfAkind(hand1)
	fourOfAKindHand2, remainingCard2 := f.findFourOfAkind(hand2)

	// FOUR OF A KIND comparison
	if fourOfAKindHand1.Value > fourOfAKindHand2.Value {
		return larvis.Hand1Wins
	}
	if fourOfAKindHand1.Value < fourOfAKindHand2.Value {
		return larvis.Hand2Wins
	}

	// remaining card comparison
	if remainingCard1.Value > remainingCard2.Value {
		return larvis.Hand1Wins
	}
	if remainingCard1.Value < remainingCard2.Value {
		return larvis.Hand2Wins
	}

	return larvis.Tie
}

// findFourOfAkind returns the FOUR OF A KIND and the remaining card.
func (fourOfAKindTieBreaker) findFourOfAkind(hand larvis.Hand) (fourOfAKind larvis.Card, remainingCard larvis.Card) {
	f := freq.CardsFreq(hand)
	for k := range f {
		if f[k] == 4 {
			fourOfAKind = k
		} else {
			remainingCard = k
		}
	}
	return
}
