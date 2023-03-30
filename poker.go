package poker

import (
	"errors"
	"fmt"

	"github.com/mroobert/larvis/game"
	"github.com/mroobert/larvis/rank"
	"github.com/mroobert/larvis/tiebreak"
)

var ErrApplyTieBreak = errors.New("failed to apply tie-break")

// Poker represents a game of poker.
type Poker struct {
	hand1   game.Hand
	hand2   game.Hand
	ranker  rank.Ranker
	decider tiebreak.Decider
}

func NewPoker(
	hand1 game.Hand,
	hand2 game.Hand,
	ranker rank.Ranker,
	decider tiebreak.Decider,
) Poker {
	return Poker{
		hand1:   hand1,
		hand2:   hand2,
		ranker:  ranker,
		decider: decider,
	}
}

// Play compares the two hands and returns the game result.
func (p Poker) Play() (string, error) {
	r1 := p.ranker.RankHand(p.hand1)
	r2 := p.ranker.RankHand(p.hand2)

	// apply tie-breaking rules
	if r1 == r2 {
		res, err := p.decider.ApplyTieBreak(r1, p.hand1, p.hand2)
		if err != nil {
			return "", fmt.Errorf("%w: %v", ErrApplyTieBreak, err)
		}
		return res, nil
	}

	if r1 > r2 {
		return game.Hand1Wins, nil
	}

	return game.Hand2Wins, nil
}
