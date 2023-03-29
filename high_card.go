package larvis

import "sort"

// highCard manages the comparison of hands with the HIGH_CARD rank.
type highCard struct{}

// compare compares the highest card in each hand.
func (highCard) compare(hand1, hand2 Hand) string {
	sort.Sort(hand1)
	sort.Sort(hand2)

	if hand1[0].value > hand2[0].value {
		return hand1Wins
	}
	if hand1[0].value < hand2[0].value {
		return hand2Wins
	}

	return tie
}
