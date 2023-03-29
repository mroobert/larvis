package larvis

// fourOfAKind manages the comparison of hands with the FOUR_OF_A_KIND rank.
type fourOfAKind struct{}

// compare compares the four of a kind in each hand.
// If the four of a kind is the same, the remaining card is compared.
func (f fourOfAKind) compare(hand1, hand2 Hand) string {
	fourOfAKindHand1, remainingCard1 := f.findFourOfAkind(hand1)
	fourOfAKindHand2, remainingCard2 := f.findFourOfAkind(hand2)

	// four of a kind comparison
	if fourOfAKindHand1.value > fourOfAKindHand2.value {
		return hand1Wins
	}
	if fourOfAKindHand1.value < fourOfAKindHand2.value {
		return hand2Wins
	}

	// remaining card comparison
	if remainingCard1.value > remainingCard2.value {
		return hand1Wins
	}
	if remainingCard1.value < remainingCard2.value {
		return hand2Wins
	}

	return tie
}

// findFourOfAkind returns the four of a kind and the remaining card.
func (f fourOfAKind) findFourOfAkind(hand Hand) (fourOfAKind card, remainingCard card) {
	freq := cardsFreq(hand)
	for k := range freq {
		if freq[k] == 4 {
			fourOfAKind = k
		} else {
			remainingCard = k
		}
	}
	return
}
