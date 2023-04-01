package tiebreak

import (
	"testing"

	"github.com/mroobert/larvis"
)

func TestCompare(t *testing.T) {
	t.Parallel()

	type args struct {
		h1 larvis.Hand
		h2 larvis.Hand
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "hand 1 wins",
			args: args{
				h1: larvis.Hand{
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
					{Value: larvis.Symbols["J"], Symbol: 'J'},
					{Value: larvis.Symbols["K"], Symbol: 'K'},
					{Value: larvis.Symbols["A"], Symbol: 'A'},
					{Value: larvis.Symbols["1"], Symbol: '1'},
				},
				h2: larvis.Hand{
					{Value: larvis.Symbols["1"], Symbol: '1'},
					{Value: larvis.Symbols["2"], Symbol: '2'},
					{Value: larvis.Symbols["3"], Symbol: '3'},
					{Value: larvis.Symbols["4"], Symbol: '4'},
					{Value: larvis.Symbols["5"], Symbol: '5'},
				},
			},
			want: larvis.Hand1Wins,
		},
		{
			name: "hand 2 wins",
			args: args{
				h1: larvis.Hand{
					{Value: larvis.Symbols["2"], Symbol: '2'},
					{Value: larvis.Symbols["7"], Symbol: '7'},
					{Value: larvis.Symbols["8"], Symbol: '8'},
					{Value: larvis.Symbols["T"], Symbol: 'T'},
					{Value: larvis.Symbols["Q"], Symbol: 'Q'},
				},
				h2: larvis.Hand{
					{Value: larvis.Symbols["3"], Symbol: '3'},
					{Value: larvis.Symbols["A"], Symbol: 'A'},
					{Value: larvis.Symbols["K"], Symbol: 'K'},
					{Value: larvis.Symbols["J"], Symbol: 'J'},
					{Value: larvis.Symbols["T"], Symbol: 'T'},
				},
			},
			want: larvis.Hand2Wins,
		},
		{
			name: "larvis.Tie",
			args: args{
				h1: larvis.Hand{
					{Value: larvis.Symbols["9"], Symbol: '9'},
					{Value: larvis.Symbols["8"], Symbol: '8'},
					{Value: larvis.Symbols["7"], Symbol: '7'},
					{Value: larvis.Symbols["6"], Symbol: '6'},
					{Value: larvis.Symbols["5"], Symbol: '5'},
				},
				h2: larvis.Hand{
					{Value: larvis.Symbols["9"], Symbol: '9'},
					{Value: larvis.Symbols["8"], Symbol: '8'},
					{Value: larvis.Symbols["7"], Symbol: '7'},
					{Value: larvis.Symbols["6"], Symbol: '6'},
					{Value: larvis.Symbols["5"], Symbol: '5'},
				},
			},
			want: larvis.Tie,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			hc := highCardTieBreaker{}
			if got := hc.compare(tt.args.h1, tt.args.h2); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
