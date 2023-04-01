package game_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/mroobert/larvis"
	"github.com/mroobert/larvis/game"
	"github.com/mroobert/larvis/rank"
	"github.com/mroobert/larvis/tiebreak"
)

func TestPlay_Tie(t *testing.T) {
	t.Parallel()

	ranker := rank.NewRanker()
	decider := tiebreak.NewDecider()

	type args struct {
		hand1   larvis.Hand
		hand2   larvis.Hand
		ranker  rank.Ranker
		decider tiebreak.Decider
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "full house vs full house - tie",
			args: args{
				hand1: larvis.Hand{
					larvis.Card{Value: larvis.Symbols["A"], Symbol: 'A'},
					larvis.Card{Value: larvis.Symbols["A"], Symbol: 'A'},
					larvis.Card{Value: larvis.Symbols["A"], Symbol: 'A'},
					larvis.Card{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					larvis.Card{Value: larvis.Symbols["Q"], Symbol: 'Q'},
				},
				hand2: larvis.Hand{
					larvis.Card{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					larvis.Card{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					larvis.Card{Value: larvis.Symbols["A"], Symbol: 'A'},
					larvis.Card{Value: larvis.Symbols["A"], Symbol: 'A'},
					larvis.Card{Value: larvis.Symbols["A"], Symbol: 'A'},
				},
				ranker:  ranker,
				decider: decider,
			},
			want: larvis.Tie,
		},
		{
			name: "high card vs high card - tie",
			args: args{
				hand1: larvis.Hand{
					larvis.Card{Value: larvis.Symbols["A"], Symbol: 'A'},
					larvis.Card{Value: larvis.Symbols["K"], Symbol: 'K'},
					larvis.Card{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					larvis.Card{Value: larvis.Symbols["J"], Symbol: 'J'},
					larvis.Card{Value: larvis.Symbols["T"], Symbol: 'T'},
				},
				hand2: larvis.Hand{
					larvis.Card{Value: larvis.Symbols["A"], Symbol: 'A'},
					larvis.Card{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					larvis.Card{Value: larvis.Symbols["K"], Symbol: 'K'},
					larvis.Card{Value: larvis.Symbols["T"], Symbol: 'T'},
					larvis.Card{Value: larvis.Symbols["J"], Symbol: 'J'},
				},
				ranker:  ranker,
				decider: decider,
			},
			want: larvis.Tie,
		},
		{
			name: "triple vs triple - tie",
			args: args{
				hand1: larvis.Hand{
					larvis.Card{Value: larvis.Symbols["A"], Symbol: 'A'},
					larvis.Card{Value: larvis.Symbols["A"], Symbol: 'A'},
					larvis.Card{Value: larvis.Symbols["A"], Symbol: 'A'},
					larvis.Card{Value: larvis.Symbols["K"], Symbol: 'K'},
					larvis.Card{Value: larvis.Symbols["Q"], Symbol: 'Q'},
				},
				hand2: larvis.Hand{
					larvis.Card{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					larvis.Card{Value: larvis.Symbols["A"], Symbol: 'A'},
					larvis.Card{Value: larvis.Symbols["A"], Symbol: 'A'},
					larvis.Card{Value: larvis.Symbols["A"], Symbol: 'A'},
					larvis.Card{Value: larvis.Symbols["K"], Symbol: 'K'},
				},
				ranker:  ranker,
				decider: decider,
			},
			want: larvis.Tie,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			g := game.NewGame(tt.args.hand1, tt.args.hand2, tt.args.ranker, tt.args.decider)
			got, err := g.Play()
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlay_Hand1Wins(t *testing.T) {
	t.Parallel()

	ranker := rank.NewRanker()
	decider := tiebreak.NewDecider()

	type args struct {
		hand1   larvis.Hand
		hand2   larvis.Hand
		ranker  rank.Ranker
		decider tiebreak.Decider
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "full house vs full house - hand1 wins",
			args: args{
				hand1: larvis.Hand{
					larvis.Card{Value: larvis.Symbols["8"], Symbol: '8'},
					larvis.Card{Value: larvis.Symbols["8"], Symbol: '8'},
					larvis.Card{Value: larvis.Symbols["8"], Symbol: '8'},
					larvis.Card{Value: larvis.Symbols["2"], Symbol: '2'},
					larvis.Card{Value: larvis.Symbols["2"], Symbol: '2'},
				},
				hand2: larvis.Hand{
					larvis.Card{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					larvis.Card{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					larvis.Card{Value: larvis.Symbols["7"], Symbol: '7'},
					larvis.Card{Value: larvis.Symbols["7"], Symbol: '7'},
					larvis.Card{Value: larvis.Symbols["7"], Symbol: '7'},
				},
				ranker:  ranker,
				decider: decider,
			},
			want: larvis.Hand1Wins,
		},
		{
			name: "four of a kind vs full house - hand1 wins",
			args: args{
				hand1: larvis.Hand{
					larvis.Card{Value: larvis.Symbols["A"], Symbol: 'A'},
					larvis.Card{Value: larvis.Symbols["A"], Symbol: 'A'},
					larvis.Card{Value: larvis.Symbols["A"], Symbol: 'A'},
					larvis.Card{Value: larvis.Symbols["A"], Symbol: 'A'},
					larvis.Card{Value: larvis.Symbols["Q"], Symbol: 'Q'},
				},
				hand2: larvis.Hand{
					larvis.Card{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					larvis.Card{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					larvis.Card{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					larvis.Card{Value: larvis.Symbols["A"], Symbol: 'A'},
					larvis.Card{Value: larvis.Symbols["A"], Symbol: 'A'},
				},
				ranker:  ranker,
				decider: decider,
			},
			want: larvis.Hand1Wins,
		},
		{
			name: "two pairs vs a pair - hand1 wins",
			args: args{
				hand1: larvis.Hand{
					larvis.Card{Value: larvis.Symbols["A"], Symbol: 'A'},
					larvis.Card{Value: larvis.Symbols["A"], Symbol: 'A'},
					larvis.Card{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					larvis.Card{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					larvis.Card{Value: larvis.Symbols["J"], Symbol: 'J'},
				},
				hand2: larvis.Hand{
					larvis.Card{Value: larvis.Symbols["A"], Symbol: 'A'},
					larvis.Card{Value: larvis.Symbols["A"], Symbol: 'A'},
					larvis.Card{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					larvis.Card{Value: larvis.Symbols["J"], Symbol: 'J'},
					larvis.Card{Value: larvis.Symbols["T"], Symbol: 'T'},
				},
				ranker:  ranker,
				decider: decider,
			},
			want: larvis.Hand1Wins,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			g := game.NewGame(tt.args.hand1, tt.args.hand2, tt.args.ranker, tt.args.decider)
			got, err := g.Play()
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlay_Hand2Wins(t *testing.T) {
	t.Parallel()

	ranker := rank.NewRanker()
	decider := tiebreak.NewDecider()

	type args struct {
		hand1   larvis.Hand
		hand2   larvis.Hand
		ranker  rank.Ranker
		decider tiebreak.Decider
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "full house vs full house - hand2 wins",
			args: args{
				hand1: larvis.Hand{
					larvis.Card{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					larvis.Card{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					larvis.Card{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					larvis.Card{Value: larvis.Symbols["A"], Symbol: 'A'},
					larvis.Card{Value: larvis.Symbols["A"], Symbol: 'A'},
				},
				hand2: larvis.Hand{
					larvis.Card{Value: larvis.Symbols["A"], Symbol: 'A'},
					larvis.Card{Value: larvis.Symbols["A"], Symbol: 'A'},
					larvis.Card{Value: larvis.Symbols["A"], Symbol: 'A'},
					larvis.Card{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					larvis.Card{Value: larvis.Symbols["Q"], Symbol: 'Q'},
				},
				ranker:  ranker,
				decider: decider,
			},
			want: larvis.Hand2Wins,
		},
		{
			name: "high Card vs one pair - hand2 wins",
			args: args{
				hand1: larvis.Hand{
					larvis.Card{Value: larvis.Symbols["A"], Symbol: 'A'},
					larvis.Card{Value: larvis.Symbols["K"], Symbol: 'K'},
					larvis.Card{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					larvis.Card{Value: larvis.Symbols["J"], Symbol: 'J'},
					larvis.Card{Value: larvis.Symbols["T"], Symbol: 'T'},
				},
				hand2: larvis.Hand{
					larvis.Card{Value: larvis.Symbols["A"], Symbol: 'A'},
					larvis.Card{Value: larvis.Symbols["A"], Symbol: 'A'},
					larvis.Card{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					larvis.Card{Value: larvis.Symbols["J"], Symbol: 'J'},
					larvis.Card{Value: larvis.Symbols["T"], Symbol: 'T'},
				},
				ranker:  ranker,
				decider: decider,
			},
			want: larvis.Hand2Wins,
		},
		{
			name: "two pairs vs a pair - hand2 wins",
			args: args{
				hand1: larvis.Hand{
					larvis.Card{Value: larvis.Symbols["2"], Symbol: '2'},
					larvis.Card{Value: larvis.Symbols["2"], Symbol: '2'},
					larvis.Card{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					larvis.Card{Value: larvis.Symbols["J"], Symbol: 'J'},
					larvis.Card{Value: larvis.Symbols["T"], Symbol: 'T'},
				},
				hand2: larvis.Hand{
					larvis.Card{Value: larvis.Symbols["A"], Symbol: 'A'},
					larvis.Card{Value: larvis.Symbols["A"], Symbol: 'A'},
					larvis.Card{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					larvis.Card{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					larvis.Card{Value: larvis.Symbols["T"], Symbol: 'T'},
				},
				ranker:  ranker,
				decider: decider,
			},
			want: larvis.Hand2Wins,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			g := game.NewGame(tt.args.hand1, tt.args.hand2, tt.args.ranker, tt.args.decider)
			got, err := g.Play()
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlay_ReturnsErrorWhenApplyTieBreak(t *testing.T) {
	t.Parallel()

	r := rank.NewRanker()
	d := tiebreak.Decider{} // nil tieBreakers

	h1 := larvis.Hand{
		larvis.Card{Value: larvis.Symbols["A"], Symbol: 'A'},
		larvis.Card{Value: larvis.Symbols["A"], Symbol: 'A'},
		larvis.Card{Value: larvis.Symbols["Q"], Symbol: 'Q'},
		larvis.Card{Value: larvis.Symbols["J"], Symbol: 'J'},
		larvis.Card{Value: larvis.Symbols["T"], Symbol: 'T'},
	}

	h2 := larvis.Hand{
		larvis.Card{Value: larvis.Symbols["A"], Symbol: 'A'},
		larvis.Card{Value: larvis.Symbols["A"], Symbol: 'A'},
		larvis.Card{Value: larvis.Symbols["Q"], Symbol: 'Q'},
		larvis.Card{Value: larvis.Symbols["J"], Symbol: 'J'},
		larvis.Card{Value: larvis.Symbols["T"], Symbol: 'T'},
	}

	g := game.NewGame(h1, h2, r, d)
	_, err := g.Play()
	fmt.Println(err)
	if !errors.Is(err, game.ErrApplyTieBreak) {
		t.Fatalf("wrong error: %v", err)
	}
}
