package tiebreak

import (
	"github.com/mroobert/larvis/game"
	"github.com/mroobert/larvis/rank/freq"
)

// fourOfAKindTieBreaker manages the tie-breaking of hands with the FOUR OF A KIND rank.
type fourOfAKindTieBreaker struct{}

// compare compares the FOUR OF A KIND in each hand.
// If the FOUR OF A KIND is the same, the remaining game.Card is compared.
func (f fourOfAKindTieBreaker) compare(hand1, hand2 game.Hand) string {
	fourOfAKindHand1, remainingCard1 := f.findFourOfAkind(hand1)
	fourOfAKindHand2, remainingCard2 := f.findFourOfAkind(hand2)

	// FOUR OF A KIND comparison
	if fourOfAKindHand1.Value > fourOfAKindHand2.Value {
		return game.Hand1Wins
	}
	if fourOfAKindHand1.Value < fourOfAKindHand2.Value {
		return game.Hand2Wins
	}

	// remaining game.Card comparison
	if remainingCard1.Value > remainingCard2.Value {
		return game.Hand1Wins
	}
	if remainingCard1.Value < remainingCard2.Value {
		return game.Hand2Wins
	}

	return game.Tie
}

// findFourOfAkind returns the FOUR OF A KIND and the remaining game.Card.
func (fourOfAKindTieBreaker) findFourOfAkind(hand game.Hand) (fourOfAKind game.Card, remainingCard game.Card) {
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
