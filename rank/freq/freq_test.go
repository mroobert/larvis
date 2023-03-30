package freq_test

import (
	"reflect"
	"testing"

	"github.com/mroobert/larvis/game"
	"github.com/mroobert/larvis/rank/freq"
)

func TestCardsFreq(t *testing.T) {
	t.Parallel()

	type args struct {
		h game.Hand
	}

	tests := []struct {
		name string
		args args
		want map[game.Card]int
	}{
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
			want: map[game.Card]int{
				{Value: game.Symbols["Q"], Symbol: 'Q'}: 2,
				{Value: game.Symbols["J"], Symbol: 'J'}: 1,
				{Value: game.Symbols["K"], Symbol: 'K'}: 1,
				{Value: game.Symbols["T"], Symbol: 'T'}: 1,
			},
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
			want: map[game.Card]int{
				{Value: game.Symbols["Q"], Symbol: 'Q'}: 2,
				{Value: game.Symbols["J"], Symbol: 'J'}: 2,
				{Value: game.Symbols["K"], Symbol: 'K'}: 1,
			},
		},
		{
			name: "three of a kind",
			args: args{
				h: game.Hand{
					{Value: game.Symbols["4"], Symbol: '4'},
					{Value: game.Symbols["4"], Symbol: '4'},
					{Value: game.Symbols["4"], Symbol: '4'},
					{Value: game.Symbols["2"], Symbol: '2'},
					{Value: game.Symbols["1"], Symbol: '1'},
				},
			},
			want: map[game.Card]int{
				{Value: game.Symbols["4"], Symbol: '4'}: 3,
				{Value: game.Symbols["2"], Symbol: '2'}: 1,
				{Value: game.Symbols["1"], Symbol: '1'}: 1,
			},
		},
		{
			name: "four of a kind",
			args: args{
				h: game.Hand{
					{Value: game.Symbols["4"], Symbol: '4'},
					{Value: game.Symbols["4"], Symbol: '4'},
					{Value: game.Symbols["4"], Symbol: '4'},
					{Value: game.Symbols["4"], Symbol: '4'},
					{Value: game.Symbols["1"], Symbol: '1'},
				},
			},
			want: map[game.Card]int{
				{Value: game.Symbols["4"], Symbol: '4'}: 4,
				{Value: game.Symbols["1"], Symbol: '1'}: 1,
			},
		},
		{
			name: "full house",
			args: args{
				h: game.Hand{
					{Value: game.Symbols["4"], Symbol: '4'},
					{Value: game.Symbols["4"], Symbol: '4'},
					{Value: game.Symbols["4"], Symbol: '4'},
					{Value: game.Symbols["1"], Symbol: '1'},
					{Value: game.Symbols["1"], Symbol: '1'},
				},
			},
			want: map[game.Card]int{
				{Value: game.Symbols["4"], Symbol: '4'}: 3,
				{Value: game.Symbols["1"], Symbol: '1'}: 2,
			},
		},
		{
			name: "high card",
			args: args{
				h: game.Hand{
					{Value: game.Symbols["4"], Symbol: '4'},
					{Value: game.Symbols["3"], Symbol: '3'},
					{Value: game.Symbols["2"], Symbol: '2'},
					{Value: game.Symbols["1"], Symbol: '1'},
					{Value: game.Symbols["T"], Symbol: 'T'},
				},
			},
			want: map[game.Card]int{
				{Value: game.Symbols["4"], Symbol: '4'}: 1,
				{Value: game.Symbols["3"], Symbol: '3'}: 1,
				{Value: game.Symbols["2"], Symbol: '2'}: 1,
				{Value: game.Symbols["1"], Symbol: '1'}: 1,
				{Value: game.Symbols["T"], Symbol: 'T'}: 1,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := freq.CardsFreq(tt.args.h); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CardsFreq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasPairs_ReturnsTrue(t *testing.T) {
	t.Parallel()

	type args struct {
		freq map[game.Card]int
		n    int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "two pairs and n is 2",
			args: args{
				freq: map[game.Card]int{
					{Value: game.Symbols["A"], Symbol: 'A'}: 2,
					{Value: game.Symbols["Q"], Symbol: 'Q'}: 2,
					{Value: game.Symbols["J"], Symbol: 'J'}: 1,
				},
				n: 2,
			},
			want: true,
		},
		{
			name: "one pair and n is 1",
			args: args{
				freq: map[game.Card]int{
					{Value: game.Symbols["Q"], Symbol: 'Q'}: 2,
					{Value: game.Symbols["J"], Symbol: 'J'}: 1,
					{Value: game.Symbols["K"], Symbol: 'K'}: 1,
					{Value: game.Symbols["T"], Symbol: 'T'}: 1,
				},
				n: 1,
			},
			want: true,
		},
		{
			name: "no pairs and n is 0",
			args: args{
				freq: map[game.Card]int{
					{Value: game.Symbols["A"], Symbol: 'A'}: 1,
					{Value: game.Symbols["2"], Symbol: '2'}: 1,
					{Value: game.Symbols["J"], Symbol: 'J'}: 1,
					{Value: game.Symbols["4"], Symbol: '4'}: 1,
					{Value: game.Symbols["T"], Symbol: 'T'}: 1,
				},
				n: 0,
			},
			want: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := freq.HasPairs(tt.args.freq, tt.args.n); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasPairs_ReturnsFalse(t *testing.T) {
	t.Parallel()

	type args struct {
		freq map[game.Card]int
		n    int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "two pairs and n is 2",
			args: args{
				freq: map[game.Card]int{
					{Value: game.Symbols["A"], Symbol: 'A'}: 3,
					{Value: game.Symbols["Q"], Symbol: 'Q'}: 1,
					{Value: game.Symbols["J"], Symbol: 'J'}: 1,
				},
				n: 2,
			},
			want: false,
		},
		{
			name: "two pairs and n is 1",
			args: args{
				freq: map[game.Card]int{
					{Value: game.Symbols["A"], Symbol: 'A'}: 2,
					{Value: game.Symbols["Q"], Symbol: 'Q'}: 2,
					{Value: game.Symbols["J"], Symbol: 'J'}: 1,
				},
				n: 1,
			},
			want: false,
		},
		{
			name: "two pairs and n is 0",
			args: args{
				freq: map[game.Card]int{
					{Value: game.Symbols["A"], Symbol: 'A'}: 2,
					{Value: game.Symbols["Q"], Symbol: 'Q'}: 2,
					{Value: game.Symbols["J"], Symbol: 'J'}: 1,
				},
				n: 0,
			},
			want: false,
		},
		{
			name: "one pair and n is 2",
			args: args{
				freq: map[game.Card]int{
					{Value: game.Symbols["Q"], Symbol: 'Q'}: 2,
					{Value: game.Symbols["J"], Symbol: 'J'}: 1,
					{Value: game.Symbols["K"], Symbol: 'K'}: 1,
					{Value: game.Symbols["T"], Symbol: 'T'}: 1,
				},
				n: 2,
			},
			want: false,
		},
		{
			name: "one pair and n is 1",
			args: args{
				freq: map[game.Card]int{
					{Value: game.Symbols["Q"], Symbol: 'Q'}: 1,
					{Value: game.Symbols["J"], Symbol: 'J'}: 1,
					{Value: game.Symbols["K"], Symbol: 'K'}: 1,
					{Value: game.Symbols["T"], Symbol: 'T'}: 1,
				},
				n: 1,
			},
			want: false,
		},
		{
			name: "one pair and n is 0",
			args: args{
				freq: map[game.Card]int{
					{Value: game.Symbols["Q"], Symbol: 'Q'}: 2,
					{Value: game.Symbols["J"], Symbol: 'J'}: 1,
					{Value: game.Symbols["K"], Symbol: 'K'}: 1,
					{Value: game.Symbols["T"], Symbol: 'T'}: 1,
				},
				n: 0,
			},
			want: false,
		},
		{
			name: "no pairs and n is 2",
			args: args{
				freq: map[game.Card]int{
					{Value: game.Symbols["A"], Symbol: 'A'}: 1,
					{Value: game.Symbols["Q"], Symbol: 'Q'}: 1,
					{Value: game.Symbols["J"], Symbol: 'J'}: 1,
					{Value: game.Symbols["K"], Symbol: 'K'}: 1,
					{Value: game.Symbols["T"], Symbol: 'T'}: 1,
				},
				n: 2,
			},
			want: false,
		},
		{
			name: "no pairs and n is 1",
			args: args{
				freq: map[game.Card]int{
					{Value: game.Symbols["A"], Symbol: 'A'}: 1,
					{Value: game.Symbols["2"], Symbol: '2'}: 1,
					{Value: game.Symbols["J"], Symbol: 'J'}: 1,
					{Value: game.Symbols["4"], Symbol: '4'}: 1,
					{Value: game.Symbols["T"], Symbol: 'T'}: 1,
				},
				n: 1,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := freq.HasPairs(tt.args.freq, tt.args.n); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
