package larvis

import (
	"fmt"
	"strings"
)

// The available symbols and their values.
var symbols = map[string]int{
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
	"J": 11,
	"Q": 12,
	"K": 13,
	"A": 14,
}

// Card represents a single playing card.
type card struct {
	symbol rune
	value  int
}

func createCard(symbol rune) (card, error) {
	if _, ok := symbols[string(symbol)]; !ok {
		return card{}, fmt.Errorf("invalid card symbol %q", string(symbol))
	}
	return card{
		symbol: symbol,
		value:  symbols[string(symbol)],
	}, nil
}

type Hand []card

func CreateHand(cards []rune) (Hand, error) {
	if len(cards) != 5 {
		return Hand{}, fmt.Errorf("invalid number of cards %q", len(cards))
	}

	h := make(Hand, 0, 5)
	for _, c := range cards {
		card, err := createCard(c)
		if err != nil {
			return Hand{}, err
		}
		h = append(h, card)
	}
	return h, nil
}

func (h Hand) String() string {
	var s strings.Builder
	for _, c := range h {
		s.WriteString(string(c.symbol))
	}
	return s.String()
}

// implement sort.Interface
func (h Hand) Len() int           { return len(h) }
func (h Hand) Less(i, j int) bool { return h[i].value > h[j].value } // Descending order.
func (h Hand) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
