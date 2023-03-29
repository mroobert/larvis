package larvis

import "sort"

// onePair manages the comparison of hands with the ONE_PAIR rank.
type onePair struct{}

// compare compares the pair in each hand.
// If the pair is the same, the remaining cards are compared.
func (o onePair) compare(hand1, hand2 Hand) string {
	pairHand1, remainingCards1 := o.findPair(hand1)
	pairHand2, remainingCards2 := o.findPair(hand2)

	// pair comparison
	if pairHand1.value > pairHand2.value {
		return hand1Wins
	}
	if pairHand1.value < pairHand2.value {
		return hand2Wins
	}

	// remaining cards comparison
	sort.Slice(remainingCards1, func(i, j int) bool {
		return remainingCards1[i].value > remainingCards1[j].value // sort descending
	})
	sort.Slice(remainingCards2, func(i, j int) bool {
		return remainingCards2[i].value > remainingCards2[j].value // sort descending
	})
	if remainingCards1[0].value > remainingCards2[0].value {
		return hand1Wins
	}
	if remainingCards1[0].value < remainingCards2[0].value {
		return hand2Wins
	}
	if remainingCards1[1].value > remainingCards2[1].value {
		return hand1Wins
	}
	if remainingCards1[1].value < remainingCards2[1].value {
		return hand2Wins
	}
	if remainingCards1[2].value > remainingCards2[2].value {
		return hand1Wins
	}
	if remainingCards1[2].value < remainingCards2[2].value {
		return hand2Wins
	}

	return tie
}

// findPair returns the pair and the remaining cards.
func (o onePair) findPair(hand Hand) (pair card, remainingCards []card) {
	freq := cardsFreq(hand)
	for k := range freq {
		if freq[k] == 2 {
			pair = k
		} else {
			remainingCards = append(remainingCards, k)
		}
	}
	return
}
