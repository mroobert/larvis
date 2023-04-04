// Package game provides support for playing a game of poker.
package game

import (
	"errors"
	"fmt"

	"github.com/mroobert/larvis"
	"github.com/mroobert/larvis/rank"
	"github.com/mroobert/larvis/tiebreak"
)

var ErrApplyTieBreak = errors.New("failed to apply tie-break")

// Game represents a game of poker.
type Game struct {
	hand1   larvis.Hand
	hand2   larvis.Hand
	ranker  rank.Ranker
	decider tiebreak.Decider
}

func NewGame(
	hand1 larvis.Hand,
	hand2 larvis.Hand,
	ranker rank.Ranker,
	decider tiebreak.Decider,
) Game {
	return Game{
		hand1:   hand1,
		hand2:   hand2,
		ranker:  ranker,
		decider: decider,
	}
}

// Play compares the two hands and returns the game result.
func (g Game) Play() (string, error) {
	r1 := g.ranker.RankHand(g.hand1)
	r2 := g.ranker.RankHand(g.hand2)

	// apply tie-breaking rules
	if r1 == r2 {
		res, err := g.decider.ApplyTieBreak(r1, g.hand1, g.hand2)
		if err != nil {
			return "", fmt.Errorf("%w: %v", ErrApplyTieBreak, err)
		}
		return res, nil
	}

	if r1 > r2 {
		return larvis.Hand1Wins, nil
	}

	return larvis.Hand2Wins, nil
}
