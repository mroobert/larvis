package tiebreak

import (
	"testing"

	"github.com/mroobert/larvis/game"
)

func TestCompare(t *testing.T) {
	t.Parallel()

	type args struct {
		h1 game.Hand
		h2 game.Hand
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "hand 1 wins",
			args: args{
				h1: game.Hand{
					{Value: game.Symbols["Q"], Symbol: 'Q'},
					{Value: game.Symbols["J"], Symbol: 'J'},
					{Value: game.Symbols["K"], Symbol: 'K'},
					{Value: game.Symbols["A"], Symbol: 'A'},
					{Value: game.Symbols["1"], Symbol: '1'},
				},
				h2: game.Hand{
					{Value: game.Symbols["1"], Symbol: '1'},
					{Value: game.Symbols["2"], Symbol: '2'},
					{Value: game.Symbols["3"], Symbol: '3'},
					{Value: game.Symbols["4"], Symbol: '4'},
					{Value: game.Symbols["5"], Symbol: '5'},
				},
			},
			want: game.Hand1Wins,
		},
		{
			name: "hand 2 wins",
			args: args{
				h1: game.Hand{
					{Value: game.Symbols["2"], Symbol: '2'},
					{Value: game.Symbols["7"], Symbol: '7'},
					{Value: game.Symbols["8"], Symbol: '8'},
					{Value: game.Symbols["T"], Symbol: 'T'},
					{Value: game.Symbols["Q"], Symbol: 'Q'},
				},
				h2: game.Hand{
					{Value: game.Symbols["3"], Symbol: '3'},
					{Value: game.Symbols["A"], Symbol: 'A'},
					{Value: game.Symbols["K"], Symbol: 'K'},
					{Value: game.Symbols["J"], Symbol: 'J'},
					{Value: game.Symbols["T"], Symbol: 'T'},
				},
			},
			want: game.Hand2Wins,
		},
		{
			name: "game.Tie",
			args: args{
				h1: game.Hand{
					{Value: game.Symbols["9"], Symbol: '9'},
					{Value: game.Symbols["8"], Symbol: '8'},
					{Value: game.Symbols["7"], Symbol: '7'},
					{Value: game.Symbols["6"], Symbol: '6'},
					{Value: game.Symbols["5"], Symbol: '5'},
				},
				h2: game.Hand{
					{Value: game.Symbols["9"], Symbol: '9'},
					{Value: game.Symbols["8"], Symbol: '8'},
					{Value: game.Symbols["7"], Symbol: '7'},
					{Value: game.Symbols["6"], Symbol: '6'},
					{Value: game.Symbols["5"], Symbol: '5'},
				},
			},
			want: game.Tie,
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
