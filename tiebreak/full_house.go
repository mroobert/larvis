package tiebreak

import (
	"github.com/mroobert/larvis"
	"github.com/mroobert/larvis/rank/freq"
)

// fullHouseTieBreaker manages the tie-breaking of hands with the FULL HOUSE rank.
type fullHouseTieBreaker struct{}

// compare compares the THREE OF A KIND. If the THREE OF A KIND is the same,
// the TWO OF A KIND is compared.
func (f fullHouseTieBreaker) compare(hand1, hand2 larvis.Hand) string {
	threeOfAKindHand1, twoOfAKindHand1 := f.findTwoAndThreeOfAKind(hand1)
	threeOfAKindHand2, twoOfAKindHand2 := f.findTwoAndThreeOfAKind(hand2)

	// THREE OF A KIND comparison
	if threeOfAKindHand1.Value > threeOfAKindHand2.Value {
		return larvis.Hand1Wins
	}
	if threeOfAKindHand1.Value < threeOfAKindHand2.Value {
		return larvis.Hand2Wins
	}

	// TWO OF A KIND comparison
	if twoOfAKindHand1.Value > twoOfAKindHand2.Value {
		return larvis.Hand1Wins
	}
	if twoOfAKindHand1.Value < twoOfAKindHand2.Value {
		return larvis.Hand2Wins
	}

	return larvis.Tie
}

// findTwoAndThreeOfAKind returns the THREE OF A KIND and the TWO OF A KIND.
func (fullHouseTieBreaker) findTwoAndThreeOfAKind(hand larvis.Hand) (threeOfAKind larvis.Card, twoOfAKind larvis.Card) {
	f := freq.CardsFreq(hand)
	for k := range f {
		switch f[k] {
		case 2:
			twoOfAKind = k
		case 3:
			threeOfAKind = k
		}
	}
	return
}
