package main

import "testing"

func TestEvenOrOdd(t *testing.T) {
	tests := []struct {
		n    int
		want string
	}{
		{
			n:    0,
			want: "even",
		},
		{
			n:    1,
			want: "odd",
		},
		{
			n:    2,
			want: "even",
		},
		{
			n:    3,
			want: "odd",
		},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := EvenOrOdd(tt.n); got != tt.want {
				t.Errorf("EvenOrOdd(%d) = %q, want %q", tt.n, got, tt.want)
			}
		})
	}
}
