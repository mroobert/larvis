package tiebreak

import (
	"sort"

	"github.com/mroobert/larvis/game"
)

// highCardTieBreaker manages the tie-breaking of hands with the HIGH CARD rank.
type highCardTieBreaker struct{}

// compare compares the highest card in each hand.
func (highCardTieBreaker) compare(hand1, hand2 game.Hand) string {
	sort.Sort(hand1)
	sort.Sort(hand2)

	if hand1[0].Value > hand2[0].Value {
		return game.Hand1Wins
	}
	if hand1[0].Value < hand2[0].Value {
		return game.Hand2Wins
	}

	return game.Tie
}
