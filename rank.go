package larvis

// HandRank represents the rank of a hand.
type HandRank int

// The available ranks from lowest to highest.
const (
	HighCardRank HandRank = iota + 1
	OnePairRank
	TwoPairRank
	TripleRank
	FullHouseRank
	FourOfAKindRank
)

func (r HandRank) String() string {
	switch r {
	case HighCardRank:
		return "High Card"
	case OnePairRank:
		return "One Pair"
	case TwoPairRank:
		return "Two Pair"
	case TripleRank:
		return "Triple"
	case FullHouseRank:
		return "Full House"
	case FourOfAKindRank:
		return "Four of a Kind"
	default:
		return "Unknown"
	}
}
