package larvis

import (
	"testing"
)

// Here I just show a small sample of tests as an example for the
// Play function. The tests are written using the table driven testing pattern.
// In a production context, I would present a complete list of possible tests which
// would include the rest of the ranks.
func TestPlay_RetursnCorrectWinner(t *testing.T) {
	t.Parallel()

	type args struct {
		hand1 Hand
		hand2 Hand
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "full house - tie",
			args: args{
				hand1: Hand{
					card{value: symbols["A"], symbol: 'A'},
					card{value: symbols["A"], symbol: 'A'},
					card{value: symbols["A"], symbol: 'A'},
					card{value: symbols["Q"], symbol: 'Q'},
					card{value: symbols["Q"], symbol: 'Q'},
				},
				hand2: Hand{
					card{value: symbols["Q"], symbol: 'Q'},
					card{value: symbols["Q"], symbol: 'Q'},
					card{value: symbols["A"], symbol: 'A'},
					card{value: symbols["A"], symbol: 'A'},
					card{value: symbols["A"], symbol: 'A'},
				},
			},
			want: tie,
		},
		{
			name: "full house - hand1 wins",
			args: args{
				hand1: Hand{
					card{value: symbols["8"], symbol: '8'},
					card{value: symbols["8"], symbol: '8'},
					card{value: symbols["8"], symbol: '8'},
					card{value: symbols["2"], symbol: '2'},
					card{value: symbols["2"], symbol: '2'},
				},
				hand2: Hand{
					card{value: symbols["Q"], symbol: 'Q'},
					card{value: symbols["Q"], symbol: 'Q'},
					card{value: symbols["7"], symbol: '7'},
					card{value: symbols["7"], symbol: '7'},
					card{value: symbols["7"], symbol: '7'},
				},
			},
			want: hand1Wins,
		},

		{
			name: "full house - hand2 wins",
			args: args{
				hand1: Hand{
					card{value: symbols["Q"], symbol: 'Q'},
					card{value: symbols["Q"], symbol: 'Q'},
					card{value: symbols["Q"], symbol: 'Q'},
					card{value: symbols["A"], symbol: 'Q'},
					card{value: symbols["A"], symbol: 'A'},
				},
				hand2: Hand{
					card{value: symbols["A"], symbol: 'A'},
					card{value: symbols["A"], symbol: 'A'},
					card{value: symbols["A"], symbol: 'A'},
					card{value: symbols["Q"], symbol: 'Q'},
					card{value: symbols["Q"], symbol: 'Q'},
				},
			},
			want: hand2Wins,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			g := NewGame(tt.args.hand1, tt.args.hand2)
			if got := g.Play(); got != tt.want {
				t.Fatalf("Play() = %v, want %v", got, tt.want)
			}
		})
	}
}
