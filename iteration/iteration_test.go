package iteration

import (
	"testing"
)

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat('a', 10)
	}
}

// func ExampleRepeat(t *testing.B){
// 	sum := Repeat('a', 5)
// 	fmt.Println(sum)
// 	// Output: 6
// }

func TestRepeat(t *testing.T) {
	type args struct {
		char  rune
		times int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "a 1 time",
			args: args{char: 'a', times: 1},
			want: "a",
		},
		{
			name: "a 5 times",
			args: args{char: 'a', times: 5},
			want: "aaaaa",
		},
		{
			name: "zero times",
			args: args{char: 'a', times: 0},
			want: "",
		},
		{
			name: "negative times",
			args: args{char: 'a', times: -1},
			want: "",
		},
		{
			name: "a 1 time (using decimal unicode)",
			args: args{char: 97, times: 1},
			want: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Repeat(tt.args.char, tt.args.times); got != tt.want {
				t.Errorf("got %q want %q", got, tt.want)
			}
		})
	}
}
