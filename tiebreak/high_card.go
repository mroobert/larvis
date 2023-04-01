package tiebreak

import (
	"sort"

	"github.com/mroobert/larvis"
)

// highCardTieBreaker manages the tie-breaking of hands with the HIGH CARD rank.
type highCardTieBreaker struct{}

// compare compares the highest card in each hand.
func (highCardTieBreaker) compare(hand1, hand2 larvis.Hand) string {
	sort.Sort(hand1)
	sort.Sort(hand2)

	if hand1[0].Value > hand2[0].Value {
		return larvis.Hand1Wins
	}
	if hand1[0].Value < hand2[0].Value {
		return larvis.Hand2Wins
	}

	return larvis.Tie
}
