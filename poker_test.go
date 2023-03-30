package poker_test

import (
	"errors"
	"testing"

	poker "github.com/mroobert/larvis"
	"github.com/mroobert/larvis/game"
	"github.com/mroobert/larvis/rank"
	"github.com/mroobert/larvis/tiebreak"
)

func TestPlay_Tie(t *testing.T) {
	t.Parallel()

	ranker := rank.NewRanker()
	decider := tiebreak.NewDecider()

	type args struct {
		hand1   game.Hand
		hand2   game.Hand
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
				hand1: game.Hand{
					game.Card{Value: game.Symbols["A"], Symbol: 'A'},
					game.Card{Value: game.Symbols["A"], Symbol: 'A'},
					game.Card{Value: game.Symbols["A"], Symbol: 'A'},
					game.Card{Value: game.Symbols["Q"], Symbol: 'Q'},
					game.Card{Value: game.Symbols["Q"], Symbol: 'Q'},
				},
				hand2: game.Hand{
					game.Card{Value: game.Symbols["Q"], Symbol: 'Q'},
					game.Card{Value: game.Symbols["Q"], Symbol: 'Q'},
					game.Card{Value: game.Symbols["A"], Symbol: 'A'},
					game.Card{Value: game.Symbols["A"], Symbol: 'A'},
					game.Card{Value: game.Symbols["A"], Symbol: 'A'},
				},
				ranker:  ranker,
				decider: decider,
			},
			want: game.Tie,
		},
		{
			name: "high card vs high card - tie",
			args: args{
				hand1: game.Hand{
					game.Card{Value: game.Symbols["A"], Symbol: 'A'},
					game.Card{Value: game.Symbols["K"], Symbol: 'K'},
					game.Card{Value: game.Symbols["Q"], Symbol: 'Q'},
					game.Card{Value: game.Symbols["J"], Symbol: 'J'},
					game.Card{Value: game.Symbols["T"], Symbol: 'T'},
				},
				hand2: game.Hand{
					game.Card{Value: game.Symbols["A"], Symbol: 'A'},
					game.Card{Value: game.Symbols["Q"], Symbol: 'Q'},
					game.Card{Value: game.Symbols["K"], Symbol: 'K'},
					game.Card{Value: game.Symbols["T"], Symbol: 'T'},
					game.Card{Value: game.Symbols["J"], Symbol: 'J'},
				},
				ranker:  ranker,
				decider: decider,
			},
			want: game.Tie,
		},
		{
			name: "triple vs triple - tie",
			args: args{
				hand1: game.Hand{
					game.Card{Value: game.Symbols["A"], Symbol: 'A'},
					game.Card{Value: game.Symbols["A"], Symbol: 'A'},
					game.Card{Value: game.Symbols["A"], Symbol: 'A'},
					game.Card{Value: game.Symbols["K"], Symbol: 'K'},
					game.Card{Value: game.Symbols["Q"], Symbol: 'Q'},
				},
				hand2: game.Hand{
					game.Card{Value: game.Symbols["Q"], Symbol: 'Q'},
					game.Card{Value: game.Symbols["A"], Symbol: 'A'},
					game.Card{Value: game.Symbols["A"], Symbol: 'A'},
					game.Card{Value: game.Symbols["A"], Symbol: 'A'},
					game.Card{Value: game.Symbols["K"], Symbol: 'K'},
				},
				ranker:  ranker,
				decider: decider,
			},
			want: game.Tie,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			p := poker.NewPoker(tt.args.hand1, tt.args.hand2, tt.args.ranker, tt.args.decider)
			got, err := p.Play()
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
		hand1   game.Hand
		hand2   game.Hand
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
				hand1: game.Hand{
					game.Card{Value: game.Symbols["8"], Symbol: '8'},
					game.Card{Value: game.Symbols["8"], Symbol: '8'},
					game.Card{Value: game.Symbols["8"], Symbol: '8'},
					game.Card{Value: game.Symbols["2"], Symbol: '2'},
					game.Card{Value: game.Symbols["2"], Symbol: '2'},
				},
				hand2: game.Hand{
					game.Card{Value: game.Symbols["Q"], Symbol: 'Q'},
					game.Card{Value: game.Symbols["Q"], Symbol: 'Q'},
					game.Card{Value: game.Symbols["7"], Symbol: '7'},
					game.Card{Value: game.Symbols["7"], Symbol: '7'},
					game.Card{Value: game.Symbols["7"], Symbol: '7'},
				},
				ranker:  ranker,
				decider: decider,
			},
			want: game.Hand1Wins,
		},
		{
			name: "four of a kind vs full house - hand1 wins",
			args: args{
				hand1: game.Hand{
					game.Card{Value: game.Symbols["A"], Symbol: 'A'},
					game.Card{Value: game.Symbols["A"], Symbol: 'A'},
					game.Card{Value: game.Symbols["A"], Symbol: 'A'},
					game.Card{Value: game.Symbols["A"], Symbol: 'A'},
					game.Card{Value: game.Symbols["Q"], Symbol: 'Q'},
				},
				hand2: game.Hand{
					game.Card{Value: game.Symbols["Q"], Symbol: 'Q'},
					game.Card{Value: game.Symbols["Q"], Symbol: 'Q'},
					game.Card{Value: game.Symbols["Q"], Symbol: 'Q'},
					game.Card{Value: game.Symbols["A"], Symbol: 'A'},
					game.Card{Value: game.Symbols["A"], Symbol: 'A'},
				},
				ranker:  ranker,
				decider: decider,
			},
			want: game.Hand1Wins,
		},
		{
			name: "two pairs vs a pair - hand1 wins",
			args: args{
				hand1: game.Hand{
					game.Card{Value: game.Symbols["A"], Symbol: 'A'},
					game.Card{Value: game.Symbols["A"], Symbol: 'A'},
					game.Card{Value: game.Symbols["Q"], Symbol: 'Q'},
					game.Card{Value: game.Symbols["Q"], Symbol: 'Q'},
					game.Card{Value: game.Symbols["J"], Symbol: 'J'},
				},
				hand2: game.Hand{
					game.Card{Value: game.Symbols["A"], Symbol: 'A'},
					game.Card{Value: game.Symbols["A"], Symbol: 'A'},
					game.Card{Value: game.Symbols["Q"], Symbol: 'Q'},
					game.Card{Value: game.Symbols["J"], Symbol: 'J'},
					game.Card{Value: game.Symbols["T"], Symbol: 'T'},
				},
				ranker:  ranker,
				decider: decider,
			},
			want: game.Hand1Wins,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			p := poker.NewPoker(tt.args.hand1, tt.args.hand2, tt.args.ranker, tt.args.decider)
			got, err := p.Play()
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
		hand1   game.Hand
		hand2   game.Hand
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
				hand1: game.Hand{
					game.Card{Value: game.Symbols["Q"], Symbol: 'Q'},
					game.Card{Value: game.Symbols["Q"], Symbol: 'Q'},
					game.Card{Value: game.Symbols["Q"], Symbol: 'Q'},
					game.Card{Value: game.Symbols["A"], Symbol: 'A'},
					game.Card{Value: game.Symbols["A"], Symbol: 'A'},
				},
				hand2: game.Hand{
					game.Card{Value: game.Symbols["A"], Symbol: 'A'},
					game.Card{Value: game.Symbols["A"], Symbol: 'A'},
					game.Card{Value: game.Symbols["A"], Symbol: 'A'},
					game.Card{Value: game.Symbols["Q"], Symbol: 'Q'},
					game.Card{Value: game.Symbols["Q"], Symbol: 'Q'},
				},
				ranker:  ranker,
				decider: decider,
			},
			want: game.Hand2Wins,
		},
		{
			name: "high Card vs one pair - hand2 wins",
			args: args{
				hand1: game.Hand{
					game.Card{Value: game.Symbols["A"], Symbol: 'A'},
					game.Card{Value: game.Symbols["K"], Symbol: 'K'},
					game.Card{Value: game.Symbols["Q"], Symbol: 'Q'},
					game.Card{Value: game.Symbols["J"], Symbol: 'J'},
					game.Card{Value: game.Symbols["T"], Symbol: 'T'},
				},
				hand2: game.Hand{
					game.Card{Value: game.Symbols["A"], Symbol: 'A'},
					game.Card{Value: game.Symbols["A"], Symbol: 'A'},
					game.Card{Value: game.Symbols["Q"], Symbol: 'Q'},
					game.Card{Value: game.Symbols["J"], Symbol: 'J'},
					game.Card{Value: game.Symbols["T"], Symbol: 'T'},
				},
				ranker:  ranker,
				decider: decider,
			},
			want: game.Hand2Wins,
		},
		{
			name: "two pairs vs a pair - hand2 wins",
			args: args{
				hand1: game.Hand{
					game.Card{Value: game.Symbols["2"], Symbol: '2'},
					game.Card{Value: game.Symbols["2"], Symbol: '2'},
					game.Card{Value: game.Symbols["Q"], Symbol: 'Q'},
					game.Card{Value: game.Symbols["J"], Symbol: 'J'},
					game.Card{Value: game.Symbols["T"], Symbol: 'T'},
				},
				hand2: game.Hand{
					game.Card{Value: game.Symbols["A"], Symbol: 'A'},
					game.Card{Value: game.Symbols["A"], Symbol: 'A'},
					game.Card{Value: game.Symbols["Q"], Symbol: 'Q'},
					game.Card{Value: game.Symbols["Q"], Symbol: 'Q'},
					game.Card{Value: game.Symbols["T"], Symbol: 'T'},
				},
				ranker:  ranker,
				decider: decider,
			},
			want: game.Hand2Wins,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			p := poker.NewPoker(tt.args.hand1, tt.args.hand2, tt.args.ranker, tt.args.decider)
			got, err := p.Play()
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlay_ReturnsErrorWhenApplyTieBreakWithNilTieBreakers(t *testing.T) {
	t.Parallel()

	r := rank.NewRanker()
	d := tiebreak.Decider{} // nil tiebreakers

	h1 := game.Hand{
		game.Card{Value: game.Symbols["A"], Symbol: 'A'},
		game.Card{Value: game.Symbols["A"], Symbol: 'A'},
		game.Card{Value: game.Symbols["Q"], Symbol: 'Q'},
		game.Card{Value: game.Symbols["J"], Symbol: 'J'},
		game.Card{Value: game.Symbols["T"], Symbol: 'T'},
	}

	h2 := game.Hand{
		game.Card{Value: game.Symbols["A"], Symbol: 'A'},
		game.Card{Value: game.Symbols["A"], Symbol: 'A'},
		game.Card{Value: game.Symbols["Q"], Symbol: 'Q'},
		game.Card{Value: game.Symbols["J"], Symbol: 'J'},
		game.Card{Value: game.Symbols["T"], Symbol: 'T'},
	}

	p := poker.NewPoker(h1, h2, r, d)
	_, err := p.Play()
	if !errors.Is(err, poker.ErrApplyTieBreak) {
		t.Fatalf("wrong error: %v", err)
	}
}
