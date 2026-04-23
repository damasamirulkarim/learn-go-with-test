package main

import "testing"

func TestHello(t *testing.T) {
	tests := []struct {
		name     string
		language string
		want     string
	}{
		{"", "", "Hello, World"},
		{"Bob", "", "Hello, Bob"},
		{"Bob", "English", "Hello, Bob"},
		{"Bob", "Indonesian", "Halo, Bob"},
		{"Bob", "Spanish", "Hola, Bob"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hello(tt.name, tt.language); got != tt.want {
				t.Errorf("Hello(%q, %q) = %q, want %q", tt.name, tt.language, got, tt.want)
			}
		})
	}
}