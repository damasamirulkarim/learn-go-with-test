package main

import "testing"

func TestRepeat(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{input: "a", want: "aaaaa"},
		{input: "ab", want: "ababababab"},
		{input: "", want: ""},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := Repeat(tt.input); got != tt.want {
				t.Errorf("Repeat(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}