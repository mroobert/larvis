package game

import (
	"fmt"
	"strings"
)

const (
	Hand1Wins = "Hand 1"
	Hand2Wins = "Hand 2"
	Tie       = "Tie"
)

// The available symbols and their values.
var Symbols = map[string]int{
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
type Card struct {
	Symbol rune
	Value  int
}

func createCard(symbol rune) (Card, error) {
	if _, ok := Symbols[string(symbol)]; !ok {
		return Card{}, fmt.Errorf("invalid card symbol %q", string(symbol))
	}
	return Card{
		Symbol: symbol,
		Value:  Symbols[string(symbol)],
	}, nil
}

func (c Card) String() string {
	return string(c.Symbol)
}

// Hand represents a hand of cards.
type Hand []Card

func CreateHand(cards []rune) (Hand, error) {
	if len(cards) != 5 {
		return Hand{}, fmt.Errorf("invalid number of cards: %d", len(cards))
	}

	h := make(Hand, 0, 5)
	symbols := make(map[rune]int)
	for _, c := range cards {
		card, err := createCard(c)
		if err != nil {
			return Hand{}, err
		}
		symbols[c]++
		if symbols[c] > 4 {
			return Hand{}, fmt.Errorf("the %q card repeats more than 4 times", string(c))
		}
		h = append(h, card)
	}
	return h, nil
}

func (h Hand) String() string {
	var s strings.Builder
	for _, c := range h {
		s.WriteString(string(c.Symbol))
	}
	return s.String()
}

// implement sort.Interface
func (h Hand) Len() int           { return len(h) }
func (h Hand) Less(i, j int) bool { return h[i].Value > h[j].Value } // Descending order.
func (h Hand) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
