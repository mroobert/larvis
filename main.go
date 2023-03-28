package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

// Game represents a game of poker.
type Game struct {
	Hand1 Hand
	Hand2 Hand
}

func ParseGame(line string) Game {
	var game Game
	hands := strings.Split(line, ",")

	game.Hand1 = make(Hand, 0, 5)
	for _, c := range hands[0] {
		game.Hand1 = append(game.Hand1, ParseCard(c))
	}

	game.Hand2 = make(Hand, 0, 5)
	for _, c := range hands[1] {
		game.Hand2 = append(game.Hand2, ParseCard(c))
	}

	return game
}

func ParseCard(card rune) Card {
	return Card{Symbol: string(card), Value: symbols[string(card)]}
}

// Card represents a single playing card/
type Card struct {
	Symbol string
	Value  int
}

// Hand represents a hand of cards.
type Hand []Card

func (h Hand) Len() int {
	return len(h)
}

func (h Hand) Less(i, j int) bool {
	return h[i].Value > h[j].Value
}

func (h Hand) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h Hand) String() string {
	var s strings.Builder
	for _, c := range h {
		s.WriteString(c.Symbol)
	}
	return s.String()
}

// Pair represents a pair of cards.
type Pair []Card

func (p Pair) Len() int {
	return len(p)
}

func (p Pair) Less(i, j int) bool {
	return p[i].Value > p[j].Value
}

func (p Pair) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// RestTriple represents the rest of the cards in a TRIPLE rank.
type RestTriple []Card

func (r RestTriple) Len() int {
	return len(r)
}

func (r RestTriple) Less(i, j int) bool {
	return r[i].Value > r[j].Value
}

func (r RestTriple) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

// RestPair represents the rest of the cards in a TWO_PAIR rank.
type RestPair []Card

func (r RestPair) Len() int {
	return len(r)
}

func (r RestPair) Less(i, j int) bool {
	return r[i].Value > r[j].Value
}

func (r RestPair) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

// Rank represents the rank of a hand.
type Rank int

// The available ranks from lowest to highest.
const (
	HIGH_CARD_RANK Rank = iota + 1
	PAIR_RANK
	TWO_PAIR_RANK
	TRIPLE_RANK
	FULL_HOUSE_RANK
	FOUR_OF_A_KIND_RANK
)

func (r Rank) String() string {
	switch r {
	case FOUR_OF_A_KIND_RANK:
		return "Four of a Kind"
	case FULL_HOUSE_RANK:
		return "Full House"
	case TRIPLE_RANK:
		return "Triple"
	case TWO_PAIR_RANK:
		return "Two Pair"
	case PAIR_RANK:
		return "Pair"
	case HIGH_CARD_RANK:
		return "High Card"
	}
	return ""
}

func CheckFourOfAKind(hand Hand) bool {
	freq := make(map[Card]int)
	for _, c := range hand {
		freq[c] = freq[c] + 1
	}

	for k := range freq {
		if freq[k] == 4 {
			return true
		}
	}
	return false
}

func CheckFullHouse(hand Hand) bool {
	freq := make(map[Card]int)
	for _, c := range hand {
		freq[c] = freq[c] + 1
	}

	twoOfAKind := false
	threeOfAkind := false
	// cards := make(map[string]Card)
	for k := range freq {
		switch freq[k] {
		case 2:
			twoOfAKind = true
			// cards["twoOfAKind"] = k
		case 3:
			threeOfAkind = true
			// cards["threeOfAkind"] = k
		}
	}
	return threeOfAkind && twoOfAKind
}

func CheckTriple(hand Hand) bool {
	freq := make(map[Card]int)
	for _, c := range hand {
		freq[c] = freq[c] + 1
	}

	twoOfAKind := false
	threeOfAkind := false
	for k := range freq {
		switch freq[k] {
		case 2:
			twoOfAKind = true
		case 3:
			threeOfAkind = true
		}
	}
	return threeOfAkind && !twoOfAKind
}

func CheckTwoPairs(hand Hand) bool {
	freq := make(map[Card]int)
	for _, c := range hand {
		freq[c] = freq[c] + 1
	}

	p := 0
	for k := range freq {
		if freq[k] == 2 {
			p++
		}
	}
	return p == 2
}

func CheckPair(hand Hand) bool {
	freq := make(map[Card]int)
	for _, c := range hand {
		freq[c] = freq[c] + 1
	}

	threeOfAkind := false
	p := 0
	for k := range freq {
		switch freq[k] {
		case 2:
			p++
		case 3:
			threeOfAkind = true
		}
	}
	return !threeOfAkind && p == 1
}

func RankHand(h Hand) Rank {
	r := HIGH_CARD_RANK

	ok := CheckFourOfAKind(h)
	if ok {
		r = FOUR_OF_A_KIND_RANK
	}

	ok = CheckFullHouse(h)
	if ok {
		r = FULL_HOUSE_RANK
	}

	ok = CheckTriple(h)
	if ok {
		r = TRIPLE_RANK
	}

	ok = CheckTwoPairs(h)
	if ok {
		r = TWO_PAIR_RANK
	}

	ok = CheckPair(h)
	if ok {
		r = PAIR_RANK
	}

	return r
}

func Winner(h1 Hand, h2 Hand) string {
	r1 := RankHand(h1)
	r2 := RankHand(h2)

	if r1 > r2 {
		return "Hand 1 wins"
	}
	if r1 < r2 {
		return "Hand 2 wins"
	}
	if r1 == r2 {
		return CompareSameRank(h1, h2, r1)
	}

	return "Tie"
}

func CompareSameRank(h1 Hand, h2 Hand, r Rank) string {
	if r == HIGH_CARD_RANK {
		sort.Sort(h1)
		sort.Sort(h2)

		if h1[0].Value > h2[0].Value {
			return "Hand 1 wins"
		}
		if h1[0].Value < h2[0].Value {
			return "Hand 2 wins"
		}
	}

	if r == FOUR_OF_A_KIND_RANK {
		freqH1 := make(map[Card]int)
		for _, c := range h1 {
			freqH1[c] = freqH1[c] + 1
		}

		var lastH1, fourOfAKindH1 Card
		for k := range freqH1 {
			if freqH1[k] == 4 {
				fourOfAKindH1 = k
			} else {
				lastH1 = k
			}
		}

		freqH2 := make(map[Card]int)
		for _, c := range h2 {
			freqH2[c] = freqH2[c] + 1
		}

		var lastH2, fourOfAKindH2 Card
		for k := range freqH2 {
			if freqH2[k] == 4 {
				fourOfAKindH2 = k
			} else {
				lastH2 = k
			}
		}

		if fourOfAKindH1.Value > fourOfAKindH2.Value {
			return "Hand 1 wins"
		}
		if fourOfAKindH1.Value < fourOfAKindH2.Value {
			return "Hand 2 wins"
		}
		if lastH1.Value > lastH2.Value {
			return "Hand 1 wins"
		}
		if lastH1.Value < lastH2.Value {
			return "Hand 2 wins"
		}

	}

	if r == FULL_HOUSE_RANK {
		freqH1 := make(map[Card]int)
		for _, c := range h1 {
			freqH1[c] = freqH1[c] + 1
		}

		cardsH1 := make(map[string]Card)
		for k := range freqH1 {
			switch freqH1[k] {
			case 2:
				cardsH1["twoOfAKind"] = k
			case 3:
				cardsH1["threeOfAkind"] = k
			}
		}

		freqH2 := make(map[Card]int)
		for _, c := range h2 {
			freqH2[c] = freqH2[c] + 1
		}

		cardsH2 := make(map[string]Card)
		for k := range freqH2 {
			switch freqH2[k] {
			case 2:
				cardsH2["twoOfAKind"] = k
			case 3:
				cardsH2["threeOfAkind"] = k
			}
		}

		if cardsH1["threeOfAkind"].Value > cardsH2["threeOfAkind"].Value {
			return "Hand 1 wins"
		}
		if cardsH1["threeOfAkind"].Value < cardsH2["threeOfAkind"].Value {
			return "Hand 2 wins"
		}
		if cardsH1["twoOfAKind"].Value > cardsH2["twoOfAKind"].Value {
			return "Hand 1 wins"
		}
		if cardsH1["twoOfAKind"].Value < cardsH2["twoOfAKind"].Value {
			return "Hand 2 wins"
		}
	}

	if r == TRIPLE_RANK {
		freqH1 := make(map[Card]int)
		for _, c := range h1 {
			freqH1[c] = freqH1[c] + 1
		}

		var tripleH1 Card
		restH1 := make(RestTriple, 0, 2)
		for k := range freqH1 {
			if freqH1[k] == 3 {
				tripleH1 = k
			} else {
				restH1 = append(restH1, k)
			}
		}

		freqH2 := make(map[Card]int)
		for _, c := range h2 {
			freqH2[c] = freqH2[c] + 1
		}

		var tripleH2 Card
		restH2 := make(RestTriple, 0, 2)
		for k := range freqH2 {
			if freqH2[k] == 3 {
				tripleH2 = k
			} else {
				restH2 = append(restH2, k)
			}
		}

		if tripleH1.Value > tripleH2.Value {
			return "Hand 1 wins"
		}
		if tripleH1.Value < tripleH2.Value {
			return "Hand 2 wins"
		}
		sort.Sort(restH1)
		sort.Sort(restH2)
		if restH1[0].Value > restH2[0].Value {
			return "Hand 1 wins"
		}
		if restH1[0].Value < restH2[0].Value {
			return "Hand 2 wins"
		}
		if restH1[1].Value > restH2[1].Value {
			return "Hand 1 wins"
		}
		if restH1[1].Value < restH2[1].Value {
			return "Hand 2 wins"
		}
	}

	if r == TWO_PAIR_RANK {
		freqH1 := make(map[Card]int)
		for _, c := range h1 {
			freqH1[c] = freqH1[c] + 1
		}

		freqH2 := make(map[Card]int)
		for _, c := range h2 {
			freqH2[c] = freqH2[c] + 1
		}

		var lastH1 Card
		pairsH1 := make(Pair, 0, 2)
		for k := range freqH1 {
			if freqH1[k] == 2 {
				pairsH1 = append(pairsH1, k)
			} else {
				lastH1 = k
			}
		}

		var lastH2 Card
		pairsH2 := make(Pair, 0, 2)
		for k := range freqH2 {
			if freqH2[k] == 2 {
				pairsH2 = append(pairsH2, k)
			} else {
				lastH2 = k
			}
		}

		sort.Sort(pairsH1)
		sort.Sort(pairsH2)

		if pairsH1[0].Value > pairsH2[0].Value {
			return "Hand 1 wins"
		}
		if pairsH1[0].Value < pairsH2[0].Value {
			return "Hand 2 wins"
		}
		if lastH1.Value > lastH2.Value {
			return "Hand 1 wins"
		}
		if lastH1.Value < lastH2.Value {
			return "Hand 2 wins"
		}
	}

	if r == PAIR_RANK {
		freqH1 := make(map[Card]int)
		for _, c := range h1 {
			freqH1[c] = freqH1[c] + 1
		}

		var pairH1 Card
		restH1 := make(RestPair, 0, 3)
		for k := range freqH1 {
			if freqH1[k] == 2 {
				pairH1 = k
			} else {
				restH1 = append(restH1, k)
			}
		}

		freqH2 := make(map[Card]int)
		for _, c := range h2 {
			freqH2[c] = freqH2[c] + 1
		}

		var pairH2 Card
		restH2 := make(RestPair, 0, 3)
		for k := range freqH2 {
			if freqH2[k] == 2 {
				pairH2 = k
			} else {
				restH2 = append(restH2, k)
			}
		}

		if pairH1.Value > pairH2.Value {
			return "Hand 1 wins"
		}
		if pairH1.Value < pairH2.Value {
			return "Hand 2 wins"
		}
		sort.Sort(restH1)
		sort.Sort(restH2)
		if restH1[0].Value > restH2[0].Value {
			return "Hand 1 wins"
		}
		if restH1[0].Value < restH2[0].Value {
			return "Hand 2 wins"
		}
		if restH1[1].Value > restH2[1].Value {
			return "Hand 1 wins"
		}
		if restH1[1].Value < restH2[1].Value {
			return "Hand 2 wins"
		}
		if restH1[2].Value > restH2[2].Value {
			return "Hand 1 wins"
		}
		if restH1[2].Value < restH2[2].Value {
			return "Hand 2 wins"
		}

	}

	return "Tie"
}

func main() {
	file, err := os.Open("games.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan() // skip header
	for scanner.Scan() {
		line := scanner.Text()
		game := ParseGame(line)
		fmt.Printf("Hand1: %s; Hand2: %s => %s\n", game.Hand1, game.Hand2, Winner(game.Hand1, game.Hand2))
	}
}
