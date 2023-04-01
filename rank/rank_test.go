package rank_test

import (
	"testing"

	"github.com/mroobert/larvis"
	"github.com/mroobert/larvis/rank"
)

func TestRankHand(t *testing.T) {
	t.Parallel()

	type args struct {
		h larvis.Hand
	}

	tests := []struct {
		name string
		args args
		want larvis.HandRank
	}{
		{
			name: "high card",
			args: args{
				h: larvis.Hand{
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["J"], Symbol: 'J'},
					{Value: larvis.Symbols["K"], Symbol: 'K'},
					{Value: larvis.Symbols["T"], Symbol: 'T'},
					{Value: larvis.Symbols["A"], Symbol: 'A'},
				},
			},
			want: larvis.HighCardRank,
		},
		{
			name: "one pair",
			args: args{
				h: larvis.Hand{
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["J"], Symbol: 'J'},
					{Value: larvis.Symbols["K"], Symbol: 'K'},
					{Value: larvis.Symbols["T"], Symbol: 'T'},
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
				},
			},
			want: larvis.OnePairRank,
		},
		{
			name: "two pairs",
			args: args{
				h: larvis.Hand{
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["J"], Symbol: 'J'},
					{Value: larvis.Symbols["K"], Symbol: 'K'},
					{Value: larvis.Symbols["J"], Symbol: 'J'},
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
				},
			},
			want: larvis.TwoPairRank,
		},
		{
			name: "triple",
			args: args{
				h: larvis.Hand{
					{Value: larvis.Symbols["4"], Symbol: '4'},
					{Value: larvis.Symbols["4"], Symbol: '4'},
					{Value: larvis.Symbols["4"], Symbol: '4'},
					{Value: larvis.Symbols["2"], Symbol: '2'},
					{Value: larvis.Symbols["1"], Symbol: '1'},
				},
			},
			want: larvis.TripleRank,
		},
		{
			name: "full house",
			args: args{
				h: larvis.Hand{
					{Value: larvis.Symbols["4"], Symbol: '4'},
					{Value: larvis.Symbols["4"], Symbol: '4'},
					{Value: larvis.Symbols["4"], Symbol: '4'},
					{Value: larvis.Symbols["2"], Symbol: '2'},
					{Value: larvis.Symbols["2"], Symbol: '2'},
				},
			},
			want: larvis.FullHouseRank,
		},
		{
			name: "four of a kind",
			args: args{
				h: larvis.Hand{
					{Value: larvis.Symbols["4"], Symbol: '4'},
					{Value: larvis.Symbols["4"], Symbol: '4'},
					{Value: larvis.Symbols["4"], Symbol: '4'},
					{Value: larvis.Symbols["4"], Symbol: '4'},
					{Value: larvis.Symbols["2"], Symbol: '2'},
				},
			},
			want: larvis.FourOfAKindRank,
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
