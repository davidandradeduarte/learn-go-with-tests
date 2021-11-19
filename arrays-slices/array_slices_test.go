package array_slices

import (
	"testing"
)

func TestSum(t *testing.T) {
	input := [5]int{1, 2, 3, 4, 5}
	got := Sum(input)
	want := 15

	if got != want {
		t.Errorf("got %d want %d, input: %v", got, want, input)
	}
}

func TestSumSlice(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "collection of 5 numbers",
			args: args{
				numbers: []int{1, 2, 3, 4, 5},
			},
			want: 15,
		},
		{
			name: "collection of 2 numbers",
			args: args{
				numbers: []int{10, 15},
			},
			want: 25,
		},
		{
			name: "collection of 10 numbers",
			args: args{
				numbers: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumSlice(tt.args.numbers); got != tt.want {
				t.Errorf("SumSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
