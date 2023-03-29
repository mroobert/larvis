package larvis

import "sort"

// twoPair manages the comparison of hands with the TWO_PAIR rank.
type twoPair struct{}

// compare compares the two pairs in each hand. The highest pair is compared first.
// If the two pairs are the same, the remaining card is compared.
func (t twoPair) compare(hand1, hand2 Hand) string {
	pairsHand1, remainingCard1 := t.findTwoPair(hand1)
	pairsHand2, remainingCard2 := t.findTwoPair(hand2)

	// pairs comparison
	sort.Slice(pairsHand1, func(i, j int) bool {
		return pairsHand1[i].value > pairsHand1[j].value // sort descending
	})
	sort.Slice(pairsHand2, func(i, j int) bool {
		return pairsHand2[i].value > pairsHand2[j].value // sort descending
	})
	if pairsHand1[0].value > pairsHand2[0].value {
		return hand1Wins
	}
	if pairsHand1[0].value < pairsHand2[0].value {
		return hand2Wins
	}
	if pairsHand1[1].value > pairsHand2[1].value {
		return hand1Wins
	}
	if pairsHand1[1].value < pairsHand2[1].value {
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

// findTwoPair returns the two pairs and the remaining card.
func (t twoPair) findTwoPair(hand Hand) (pairs []card, remainingCard card) {
	freq := cardsFreq(hand)
	for k := range freq {
		if freq[k] == 2 {
			pairs = append(pairs, k)
		} else {
			remainingCard = k
		}
	}
	return
}
