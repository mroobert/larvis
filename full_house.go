package larvis

// fullHouse manages the comparison of hands with the FULL_HOUSE rank.
type fullHouse struct{}

// compare compares the three of a kind. If the three of a kind is the same,
// the two of a kind is compared.
func (f fullHouse) compare(hand1, hand2 Hand) string {
	threeOfAKindHand1, twoOfAKindHand1 := f.findTwoAndThreeOfAKind(hand1)
	threeOfAKindHand2, twoOfAKindHand2 := f.findTwoAndThreeOfAKind(hand2)

	// three of a kind comparison
	if threeOfAKindHand1.value > threeOfAKindHand2.value {
		return hand1Wins
	}
	if threeOfAKindHand1.value < threeOfAKindHand2.value {
		return hand2Wins
	}

	// two of a kind comparison
	if twoOfAKindHand1.value > twoOfAKindHand2.value {
		return hand1Wins
	}
	if twoOfAKindHand1.value < twoOfAKindHand2.value {
		return hand2Wins
	}

	return tie
}

// findTwoAndThreeOfAKind returns the three of a kind and the two of a kind.
func (f fullHouse) findTwoAndThreeOfAKind(hand Hand) (threeOfAKind card, twoOfAKind card) {
	freq := cardsFreq(hand)
	for k := range freq {
		switch freq[k] {
		case 2:
			twoOfAKind = k
		case 3:
			threeOfAKind = k
		}
	}
	return
}
