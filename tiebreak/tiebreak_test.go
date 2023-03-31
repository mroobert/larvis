package tiebreak

import (
	"errors"
	"testing"

	"github.com/mroobert/larvis/game"
)

func TestApplyTieBreak_ReturnsAResult(t *testing.T) {
	t.Parallel()

	type args struct {
		hand1 game.Hand
		hand2 game.Hand
		r     game.HandRank
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "high card vs high card - tie",
			args: args{
				hand1: game.Hand{
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["J"], Symbol: 'J'},
					{Value: game.Symbols["K"], Symbol: 'K'},
					{Value: game.Symbols["T"], Symbol: 'T'},
					{Value: game.Symbols["A"], Symbol: 'A'},
				},
				hand2: game.Hand{
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["J"], Symbol: 'J'},
					{Value: game.Symbols["K"], Symbol: 'K'},
					{Value: game.Symbols["T"], Symbol: 'T'},
					{Value: game.Symbols["A"], Symbol: 'A'},
				},
				r: game.HighCardRank,
			},
			want: game.Tie,
		},
		{
			name: "high card vs high card - hand1 wins",
			args: args{
				hand1: game.Hand{
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["J"], Symbol: 'J'},
					{Value: game.Symbols["K"], Symbol: 'K'},
					{Value: game.Symbols["T"], Symbol: 'T'},
					{Value: game.Symbols["A"], Symbol: 'A'},
				},
				hand2: game.Hand{
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["J"], Symbol: 'J'},
					{Value: game.Symbols["K"], Symbol: 'K'},
					{Value: game.Symbols["T"], Symbol: 'T'},
					{Value: game.Symbols["9"], Symbol: '9'},
				},
				r: game.HighCardRank,
			},
			want: game.Hand1Wins,
		},
		{
			name: "high card vs high card - hand2 wins",
			args: args{
				hand1: game.Hand{
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["J"], Symbol: 'J'},
					{Value: game.Symbols["K"], Symbol: 'K'},
					{Value: game.Symbols["T"], Symbol: 'T'},
					{Value: game.Symbols["9"], Symbol: '9'},
				},
				hand2: game.Hand{
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["J"], Symbol: 'J'},
					{Value: game.Symbols["K"], Symbol: 'K'},
					{Value: game.Symbols["T"], Symbol: 'T'},
					{Value: game.Symbols["A"], Symbol: 'A'},
				},
				r: game.HighCardRank,
			},
			want: game.Hand2Wins,
		},
		{
			name: "one pair vs one pair - tie",
			args: args{
				hand1: game.Hand{
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["K"], Symbol: 'K'},
					{Value: game.Symbols["T"], Symbol: 'T'},
					{Value: game.Symbols["A"], Symbol: 'A'},
				},
				hand2: game.Hand{
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["K"], Symbol: 'K'},
					{Value: game.Symbols["T"], Symbol: 'T'},
					{Value: game.Symbols["A"], Symbol: 'A'},
				},
				r: game.OnePairRank,
			},
			want: game.Tie,
		},
		{
			name: "one pair vs one pair - hand1 wins",
			args: args{
				hand1: game.Hand{
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["K"], Symbol: 'K'},
					{Value: game.Symbols["T"], Symbol: 'T'},
					{Value: game.Symbols["A"], Symbol: 'A'},
				},
				hand2: game.Hand{
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["K"], Symbol: 'K'},
					{Value: game.Symbols["T"], Symbol: 'T'},
					{Value: game.Symbols["9"], Symbol: '9'},
				},
				r: game.OnePairRank,
			},
			want: game.Hand1Wins,
		},
		{
			name: "one pair vs one pair - hand2 wins",
			args: args{
				hand1: game.Hand{
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["K"], Symbol: 'K'},
					{Value: game.Symbols["T"], Symbol: 'T'},
					{Value: game.Symbols["9"], Symbol: '9'},
				},
				hand2: game.Hand{
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["K"], Symbol: 'K'},
					{Value: game.Symbols["T"], Symbol: 'T'},
					{Value: game.Symbols["A"], Symbol: 'A'},
				},
				r: game.OnePairRank,
			},
			want: game.Hand2Wins,
		},
		{
			name: "two pair vs two pair - tie",
			args: args{
				hand1: game.Hand{
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["K"], Symbol: 'K'},
					{Value: game.Symbols["K"], Symbol: 'K'},
					{Value: game.Symbols["A"], Symbol: 'A'},
				},
				hand2: game.Hand{
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["K"], Symbol: 'K'},
					{Value: game.Symbols["K"], Symbol: 'K'},
					{Value: game.Symbols["A"], Symbol: 'A'},
				},
				r: game.TwoPairRank,
			},
			want: game.Tie,
		},
		{
			name: "two pair vs two pair - hand1 wins",
			args: args{
				hand1: game.Hand{
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["K"], Symbol: 'K'},
					{Value: game.Symbols["K"], Symbol: 'K'},
					{Value: game.Symbols["A"], Symbol: 'A'},
				},
				hand2: game.Hand{
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["K"], Symbol: 'K'},
					{Value: game.Symbols["K"], Symbol: 'K'},
					{Value: game.Symbols["9"], Symbol: '9'},
				},
				r: game.TwoPairRank,
			},
			want: game.Hand1Wins,
		},
		{
			name: "two pair vs two pair - hand2 wins",
			args: args{
				hand1: game.Hand{
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["K"], Symbol: 'K'},
					{Value: game.Symbols["K"], Symbol: 'K'},
					{Value: game.Symbols["9"], Symbol: '9'},
				},
				hand2: game.Hand{
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["K"], Symbol: 'K'},
					{Value: game.Symbols["K"], Symbol: 'K'},
					{Value: game.Symbols["A"], Symbol: 'A'},
				},
				r: game.TwoPairRank,
			},
			want: game.Hand2Wins,
		},
		{
			name: "triple vs triple - tie",
			args: args{
				hand1: game.Hand{
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["K"], Symbol: 'K'},
					{Value: game.Symbols["A"], Symbol: 'A'},
				},
				hand2: game.Hand{
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["K"], Symbol: 'K'},
					{Value: game.Symbols["A"], Symbol: 'A'},
				},
				r: game.TripleRank,
			},
			want: game.Tie,
		},
		{
			name: "triple vs triple - hand1 wins",
			args: args{
				hand1: game.Hand{
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["K"], Symbol: 'K'},
					{Value: game.Symbols["A"], Symbol: 'A'},
				},
				hand2: game.Hand{
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["K"], Symbol: 'K'},
					{Value: game.Symbols["9"], Symbol: '9'},
				},
				r: game.TripleRank,
			},
			want: game.Hand1Wins,
		},
		{
			name: "triple vs triple - hand2 wins",
			args: args{
				hand1: game.Hand{
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["K"], Symbol: 'K'},
					{Value: game.Symbols["9"], Symbol: '9'},
				},
				hand2: game.Hand{
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["K"], Symbol: 'K'},
					{Value: game.Symbols["A"], Symbol: 'A'},
				},
				r: game.TripleRank,
			},
			want: game.Hand2Wins,
		},
		{
			name: "full house vs full house - tie",
			args: args{
				hand1: game.Hand{
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["K"], Symbol: 'K'},
					{Value: game.Symbols["K"], Symbol: 'K'},
				},
				hand2: game.Hand{
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["K"], Symbol: 'K'},
					{Value: game.Symbols["K"], Symbol: 'K'},
				},
				r: game.FullHouseRank,
			},
			want: game.Tie,
		},
		{
			name: "full house vs full house - hand1 wins",
			args: args{
				hand1: game.Hand{
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["K"], Symbol: 'K'},
					{Value: game.Symbols["K"], Symbol: 'K'},
				},
				hand2: game.Hand{
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["2"], Symbol: '2'},
					{Value: game.Symbols["2"], Symbol: '2'},
				},
				r: game.FullHouseRank,
			},
			want: game.Hand1Wins,
		},
		{
			name: "full house vs full house - hand2 wins",
			args: args{
				hand1: game.Hand{
					{Value: game.Symbols["T"], Symbol: 'T'},
					{Value: game.Symbols["T"], Symbol: 'T'},
					{Value: game.Symbols["T"], Symbol: 'T'},
					{Value: game.Symbols["K"], Symbol: 'K'},
					{Value: game.Symbols["9"], Symbol: '9'},
				},
				hand2: game.Hand{
					{Value: game.Symbols["J"], Symbol: 'J'},
					{Value: game.Symbols["J"], Symbol: 'J'},
					{Value: game.Symbols["J"], Symbol: 'J'},
					{Value: game.Symbols["K"], Symbol: 'K'},
					{Value: game.Symbols["T"], Symbol: 'T'},
				},
				r: game.FullHouseRank,
			},
			want: game.Hand2Wins,
		},
		{
			name: "four of a kind vs four of a kind - tie",
			args: args{
				hand1: game.Hand{
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["K"], Symbol: 'K'},
				},
				hand2: game.Hand{
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["K"], Symbol: 'K'},
				},
				r: game.FourOfAKindRank,
			},
			want: game.Tie,
		},
		{
			name: "four of a kind vs four of a kind - hand1 wins",
			args: args{
				hand1: game.Hand{
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["K"], Symbol: 'K'},
				},
				hand2: game.Hand{
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["2"], Symbol: '2'},
				},
				r: game.FourOfAKindRank,
			},
			want: game.Hand1Wins,
		},
		{
			name: "four of a kind vs four of a kind - hand2 wins",
			args: args{
				hand1: game.Hand{
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["1"], Symbol: '1'},
				},
				hand2: game.Hand{
					{Value: game.Symbols["A"], Symbol: 'A'},
					{Value: game.Symbols["A"], Symbol: 'A'},
					{Value: game.Symbols["A"], Symbol: 'A'},
					{Value: game.Symbols["A"], Symbol: 'A'},
					{Value: game.Symbols["T"], Symbol: 'T'},
				},
				r: game.FourOfAKindRank,
			},
			want: game.Hand2Wins,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			d := NewDecider()
			got, err := d.ApplyTieBreak(tt.args.r, tt.args.hand1, tt.args.hand2)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Fatalf("got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestApplyTieBreak_ReturnsErrorForTieBreakerNotFound(t *testing.T) {
	t.Parallel()

	type args struct {
		r     game.HandRank
		d     Decider
		hand1 game.Hand
		hand2 game.Hand
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "unknown rank",
			args: args{
				r: game.HandRank(100), // unknown rank
				d: NewDecider(),
			},
		},
		{
			name: "empty tie breakers map",
			args: args{
				r: game.HighCardRank,
				d: Decider{
					tieBreakers: map[game.HandRank]tieBreaker{},
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			_, err := tt.args.d.ApplyTieBreak(tt.args.r, tt.args.hand1, tt.args.hand2)
			if !errors.Is(err, ErrTieBreakerNotFound) {
				t.Fatalf("wrong error: %v", err)
			}
		})
	}
}

func TestApplyTieBreak_ReturnsErrorForNilTieBreakers(t *testing.T) {
	t.Parallel()
	d := Decider{} // nil tieBreakers
	_, err := d.ApplyTieBreak(game.HighCardRank, game.Hand{}, game.Hand{})
	if !errors.Is(err, ErrTieBreakersNil) {
		t.Fatalf("wrong error: %v", err)
	}
}
