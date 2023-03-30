package rank_test

import (
	"testing"

	"github.com/mroobert/larvis/game"
	"github.com/mroobert/larvis/rank"
)

func TestRankHand(t *testing.T) {
	t.Parallel()

	type args struct {
		h game.Hand
	}

	tests := []struct {
		name string
		args args
		want game.HandRank
	}{
		{
			name: "high card",
			args: args{
				h: game.Hand{
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["J"], Symbol: 'J'},
					{Value: game.Symbols["K"], Symbol: 'K'},
					{Value: game.Symbols["T"], Symbol: 'T'},
					{Value: game.Symbols["A"], Symbol: 'A'},
				},
			},
			want: game.HighCardRank,
		},
		{
			name: "one pair",
			args: args{
				h: game.Hand{
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["J"], Symbol: 'J'},
					{Value: game.Symbols["K"], Symbol: 'K'},
					{Value: game.Symbols["T"], Symbol: 'T'},
					{Value: game.Symbols["Q"], Symbol: 'Q'},
				},
			},
			want: game.OnePairRank,
		},
		{
			name: "two pairs",
			args: args{
				h: game.Hand{
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["J"], Symbol: 'J'},
					{Value: game.Symbols["K"], Symbol: 'K'},
					{Value: game.Symbols["J"], Symbol: 'J'},
					{Value: game.Symbols["Q"], Symbol: 'Q'},
				},
			},
			want: game.TwoPairRank,
		},
		{
			name: "triple",
			args: args{
				h: game.Hand{
					{Value: game.Symbols["4"], Symbol: '4'},
					{Value: game.Symbols["4"], Symbol: '4'},
					{Value: game.Symbols["4"], Symbol: '4'},
					{Value: game.Symbols["2"], Symbol: '2'},
					{Value: game.Symbols["1"], Symbol: '1'},
				},
			},
			want: game.TripleRank,
		},
		{
			name: "full house",
			args: args{
				h: game.Hand{
					{Value: game.Symbols["4"], Symbol: '4'},
					{Value: game.Symbols["4"], Symbol: '4'},
					{Value: game.Symbols["4"], Symbol: '4'},
					{Value: game.Symbols["2"], Symbol: '2'},
					{Value: game.Symbols["2"], Symbol: '2'},
				},
			},
			want: game.FullHouseRank,
		},
		{
			name: "four of a kind",
			args: args{
				h: game.Hand{
					{Value: game.Symbols["4"], Symbol: '4'},
					{Value: game.Symbols["4"], Symbol: '4'},
					{Value: game.Symbols["4"], Symbol: '4'},
					{Value: game.Symbols["4"], Symbol: '4'},
					{Value: game.Symbols["2"], Symbol: '2'},
				},
			},
			want: game.FourOfAKindRank,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			r := rank.NewRanker()
			if got := r.RankHand(tt.args.h); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}

}
