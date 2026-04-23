package main

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		a    int
		b    int
		want int
	}{
		{a: 1, b: 1, want: 2},
		{a: -1, b: -1, want: -2},
		{a: -1, b: 1, want: 0},
		{a: 0, b: 0, want: 0},
		{a: 5, b: -3, want: 2},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Add(tt.a, tt.b); got != tt.want {
				t.Errorf("Add(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}