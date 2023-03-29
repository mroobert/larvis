package larvis

const (
	hand1Wins = "Hand 1"
	hand2Wins = "Hand 2"
	tie       = "Tie"
)

// comparer compares two hands with the same rank.
type comparer interface {
	compare(hand1, hand2 Hand) string
}

// Game represents a game of poker.
type Game struct {
	hand1  Hand
	hand2  Hand
	ranker ranker
}

func NewGame(hand1 Hand, hand2 Hand) Game {
	return Game{
		hand1:  hand1,
		hand2:  hand2,
		ranker: ranker{},
	}
}

// Play compares the two hands and returns the game result.
func (g Game) Play() string {
	r1 := g.ranker.rankHand(g.hand1)
	r2 := g.ranker.rankHand(g.hand2)

	// apply same rank comparison logic
	if r1 == r2 {
		return g.getComparer(r1).compare(g.hand1, g.hand2)
	}

	if r1 > r2 {
		return hand1Wins
	}

	return hand2Wins
}

// getComparer returns the comparer for the given rank.
func (g Game) getComparer(r rank) comparer {

	switch r {
	case fourOfAKindRank:
		return fourOfAKind{}
	case fullHouseRank:
		return fullHouse{}
	case tripleRank:
		return triple{}
	case twoPairRank:
		return twoPair{}
	case onePairRank:
		return onePair{}
	case highCardRank:
		return highCard{}
	}

	return nil
}
