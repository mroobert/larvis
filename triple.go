package larvis

import "sort"

// triple manages the comparison of hands with the TRIPLE rank.
type triple struct{}

// compare compares the triple in each hand. If the triple is the same, the
// remaining cards are compared.
func (t triple) compare(hand1, hand2 Hand) string {
	tripleHand1, remainingCards1 := t.findTriple(hand1)
	tripleHand2, remainingCards2 := t.findTriple(hand2)

	// triple comparison
	if tripleHand1.value > tripleHand2.value {
		return hand1Wins
	}
	if tripleHand1.value < tripleHand2.value {
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

	return tie
}

// findTriple returns the triple and the remaining cards.
func (t triple) findTriple(hand Hand) (triple card, remainingCards []card) {
	freq := cardsFreq(hand)
	for k := range freq {
		if freq[k] == 3 {
			triple = k
		} else {
			remainingCards = append(remainingCards, k)
		}
	}
	return
}
