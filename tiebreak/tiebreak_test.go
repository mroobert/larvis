package tiebreak

import (
	"errors"
	"testing"

	"github.com/mroobert/larvis"
)

func TestApplyTieBreak_ReturnsAResult(t *testing.T) {
	t.Parallel()

	tests := getTests()
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
		r     larvis.HandRank
		d     Decider
		hand1 larvis.Hand
		hand2 larvis.Hand
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "unknown rank",
			args: args{
				r: larvis.HandRank(100), // unknown rank
				d: NewDecider(),
			},
		},
		{
			name: "empty tie breakers map",
			args: args{
				r: larvis.HighCardRank,
				d: Decider{
					tieBreakers: map[larvis.HandRank]tieBreaker{},
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
	_, err := d.ApplyTieBreak(larvis.HighCardRank, larvis.Hand{}, larvis.Hand{})
	if !errors.Is(err, ErrTieBreakersNil) {
		t.Fatalf("wrong error: %v", err)
	}
}

type (
	args struct {
		hand1 larvis.Hand
		hand2 larvis.Hand
		r     larvis.HandRank
	}
	tests struct {
		name string
		args args
		want string
	}
)

func getTests() []tests {
	return []tests{
		{
			name: "high card vs high card - tie",
			args: args{
				hand1: larvis.Hand{
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["J"], Symbol: 'J'},
					{Value: larvis.Symbols["K"], Symbol: 'K'},
					{Value: larvis.Symbols["T"], Symbol: 'T'},
					{Value: larvis.Symbols["A"], Symbol: 'A'},
				},
				hand2: larvis.Hand{
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["J"], Symbol: 'J'},
					{Value: larvis.Symbols["K"], Symbol: 'K'},
					{Value: larvis.Symbols["T"], Symbol: 'T'},
					{Value: larvis.Symbols["A"], Symbol: 'A'},
				},
				r: larvis.HighCardRank,
			},
			want: larvis.Tie,
		},
		{
			name: "high card vs high card - hand1 wins",
			args: args{
				hand1: larvis.Hand{
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["J"], Symbol: 'J'},
					{Value: larvis.Symbols["K"], Symbol: 'K'},
					{Value: larvis.Symbols["T"], Symbol: 'T'},
					{Value: larvis.Symbols["A"], Symbol: 'A'},
				},
				hand2: larvis.Hand{
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["J"], Symbol: 'J'},
					{Value: larvis.Symbols["K"], Symbol: 'K'},
					{Value: larvis.Symbols["T"], Symbol: 'T'},
					{Value: larvis.Symbols["9"], Symbol: '9'},
				},
				r: larvis.HighCardRank,
			},
			want: larvis.Hand1Wins,
		},
		{
			name: "high card vs high card - hand2 wins",
			args: args{
				hand1: larvis.Hand{
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["J"], Symbol: 'J'},
					{Value: larvis.Symbols["K"], Symbol: 'K'},
					{Value: larvis.Symbols["T"], Symbol: 'T'},
					{Value: larvis.Symbols["9"], Symbol: '9'},
				},
				hand2: larvis.Hand{
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["J"], Symbol: 'J'},
					{Value: larvis.Symbols["K"], Symbol: 'K'},
					{Value: larvis.Symbols["T"], Symbol: 'T'},
					{Value: larvis.Symbols["A"], Symbol: 'A'},
				},
				r: larvis.HighCardRank,
			},
			want: larvis.Hand2Wins,
		},
		{
			name: "one pair vs one pair - tie",
			args: args{
				hand1: larvis.Hand{
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["K"], Symbol: 'K'},
					{Value: larvis.Symbols["T"], Symbol: 'T'},
					{Value: larvis.Symbols["A"], Symbol: 'A'},
				},
				hand2: larvis.Hand{
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["K"], Symbol: 'K'},
					{Value: larvis.Symbols["T"], Symbol: 'T'},
					{Value: larvis.Symbols["A"], Symbol: 'A'},
				},
				r: larvis.OnePairRank,
			},
			want: larvis.Tie,
		},
		{
			name: "one pair vs one pair - hand1 wins",
			args: args{
				hand1: larvis.Hand{
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["K"], Symbol: 'K'},
					{Value: larvis.Symbols["T"], Symbol: 'T'},
					{Value: larvis.Symbols["A"], Symbol: 'A'},
				},
				hand2: larvis.Hand{
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["K"], Symbol: 'K'},
					{Value: larvis.Symbols["T"], Symbol: 'T'},
					{Value: larvis.Symbols["9"], Symbol: '9'},
				},
				r: larvis.OnePairRank,
			},
			want: larvis.Hand1Wins,
		},
		{
			name: "one pair vs one pair - hand2 wins",
			args: args{
				hand1: larvis.Hand{
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["K"], Symbol: 'K'},
					{Value: larvis.Symbols["T"], Symbol: 'T'},
					{Value: larvis.Symbols["9"], Symbol: '9'},
				},
				hand2: larvis.Hand{
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["K"], Symbol: 'K'},
					{Value: larvis.Symbols["T"], Symbol: 'T'},
					{Value: larvis.Symbols["A"], Symbol: 'A'},
				},
				r: larvis.OnePairRank,
			},
			want: larvis.Hand2Wins,
		},
		{
			name: "two pair vs two pair - tie",
			args: args{
				hand1: larvis.Hand{
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["K"], Symbol: 'K'},
					{Value: larvis.Symbols["K"], Symbol: 'K'},
					{Value: larvis.Symbols["A"], Symbol: 'A'},
				},
				hand2: larvis.Hand{
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["K"], Symbol: 'K'},
					{Value: larvis.Symbols["K"], Symbol: 'K'},
					{Value: larvis.Symbols["A"], Symbol: 'A'},
				},
				r: larvis.TwoPairRank,
			},
			want: larvis.Tie,
		},
		{
			name: "two pair vs two pair - hand1 wins",
			args: args{
				hand1: larvis.Hand{
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["K"], Symbol: 'K'},
					{Value: larvis.Symbols["K"], Symbol: 'K'},
					{Value: larvis.Symbols["A"], Symbol: 'A'},
				},
				hand2: larvis.Hand{
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["K"], Symbol: 'K'},
					{Value: larvis.Symbols["K"], Symbol: 'K'},
					{Value: larvis.Symbols["9"], Symbol: '9'},
				},
				r: larvis.TwoPairRank,
			},
			want: larvis.Hand1Wins,
		},
		{
			name: "two pair vs two pair - hand2 wins",
			args: args{
				hand1: larvis.Hand{
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["K"], Symbol: 'K'},
					{Value: larvis.Symbols["K"], Symbol: 'K'},
					{Value: larvis.Symbols["9"], Symbol: '9'},
				},
				hand2: larvis.Hand{
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["K"], Symbol: 'K'},
					{Value: larvis.Symbols["K"], Symbol: 'K'},
					{Value: larvis.Symbols["A"], Symbol: 'A'},
				},
				r: larvis.TwoPairRank,
			},
			want: larvis.Hand2Wins,
		},
		{
			name: "triple vs triple - tie",
			args: args{
				hand1: larvis.Hand{
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["K"], Symbol: 'K'},
					{Value: larvis.Symbols["A"], Symbol: 'A'},
				},
				hand2: larvis.Hand{
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["K"], Symbol: 'K'},
					{Value: larvis.Symbols["A"], Symbol: 'A'},
				},
				r: larvis.TripleRank,
			},
			want: larvis.Tie,
		},
		{
			name: "triple vs triple - hand1 wins",
			args: args{
				hand1: larvis.Hand{
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["K"], Symbol: 'K'},
					{Value: larvis.Symbols["A"], Symbol: 'A'},
				},
				hand2: larvis.Hand{
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["K"], Symbol: 'K'},
					{Value: larvis.Symbols["9"], Symbol: '9'},
				},
				r: larvis.TripleRank,
			},
			want: larvis.Hand1Wins,
		},
		{
			name: "triple vs triple - hand2 wins",
			args: args{
				hand1: larvis.Hand{
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["K"], Symbol: 'K'},
					{Value: larvis.Symbols["9"], Symbol: '9'},
				},
				hand2: larvis.Hand{
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["K"], Symbol: 'K'},
					{Value: larvis.Symbols["A"], Symbol: 'A'},
				},
				r: larvis.TripleRank,
			},
			want: larvis.Hand2Wins,
		},
		{
			name: "full house vs full house - tie",
			args: args{
				hand1: larvis.Hand{
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["K"], Symbol: 'K'},
					{Value: larvis.Symbols["K"], Symbol: 'K'},
				},
				hand2: larvis.Hand{
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["K"], Symbol: 'K'},
					{Value: larvis.Symbols["K"], Symbol: 'K'},
				},
				r: larvis.FullHouseRank,
			},
			want: larvis.Tie,
		},
		{
			name: "full house vs full house - hand1 wins",
			args: args{
				hand1: larvis.Hand{
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["K"], Symbol: 'K'},
					{Value: larvis.Symbols["K"], Symbol: 'K'},
				},
				hand2: larvis.Hand{
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["2"], Symbol: '2'},
					{Value: larvis.Symbols["2"], Symbol: '2'},
				},
				r: larvis.FullHouseRank,
			},
			want: larvis.Hand1Wins,
		},
		{
			name: "full house vs full house - hand2 wins",
			args: args{
				hand1: larvis.Hand{
					{Value: larvis.Symbols["T"], Symbol: 'T'},
					{Value: larvis.Symbols["T"], Symbol: 'T'},
					{Value: larvis.Symbols["T"], Symbol: 'T'},
					{Value: larvis.Symbols["K"], Symbol: 'K'},
					{Value: larvis.Symbols["9"], Symbol: '9'},
				},
				hand2: larvis.Hand{
					{Value: larvis.Symbols["J"], Symbol: 'J'},
					{Value: larvis.Symbols["J"], Symbol: 'J'},
					{Value: larvis.Symbols["J"], Symbol: 'J'},
					{Value: larvis.Symbols["K"], Symbol: 'K'},
					{Value: larvis.Symbols["T"], Symbol: 'T'},
				},
				r: larvis.FullHouseRank,
			},
			want: larvis.Hand2Wins,
		},
		{
			name: "four of a kind vs four of a kind - tie",
			args: args{
				hand1: larvis.Hand{
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["K"], Symbol: 'K'},
				},
				hand2: larvis.Hand{
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["K"], Symbol: 'K'},
				},
				r: larvis.FourOfAKindRank,
			},
			want: larvis.Tie,
		},
		{
			name: "four of a kind vs four of a kind - hand1 wins",
			args: args{
				hand1: larvis.Hand{
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["K"], Symbol: 'K'},
				},
				hand2: larvis.Hand{
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["2"], Symbol: '2'},
				},
				r: larvis.FourOfAKindRank,
			},
			want: larvis.Hand1Wins,
		},
		{
			name: "four of a kind vs four of a kind - hand2 wins",
			args: args{
				hand1: larvis.Hand{
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["1"], Symbol: '1'},
				},
				hand2: larvis.Hand{
					{Value: larvis.Symbols["A"], Symbol: 'A'},
					{Value: larvis.Symbols["A"], Symbol: 'A'},
					{Value: larvis.Symbols["A"], Symbol: 'A'},
					{Value: larvis.Symbols["A"], Symbol: 'A'},
					{Value: larvis.Symbols["T"], Symbol: 'T'},
				},
				r: larvis.FourOfAKindRank,
			},
			want: larvis.Hand2Wins,
		},
	}
}
