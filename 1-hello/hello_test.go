package main

import "testing"

func TestHello(t *testing.T) {
	tests := []struct {
		name     string
		want     string
	}{
		{"", "Hello, World"},
		{"Bob", "Hello, Bob"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hello(tt.name); got != tt.want {
				t.Errorf("Hello(%q) = %q, want %q", tt.name, got, tt.want)
			}
		})
	}
}