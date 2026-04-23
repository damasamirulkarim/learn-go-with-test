## Overview
Create the `3-iteration/` chapter following existing conventions. The `Repeat` function repeats a string 5 times using a simple `for` loop. Tests use the same table-driven pattern as chapters 1-2.

## Changes

### 1. Implementation
**Files:**
- `3-iteration/repeat.go` (new)

**Description:**
- `package main`, exported `Repeat(s string) string`
- Hardcoded repeat count of 5 via `for` loop, concatenate into result string
- No imports needed (strings package unnecessary for simple concatenation)

```go
package main

func Repeat(s string) string {
	var repeated string
	for i := 0; i < 5; i++ {
		repeated += s
	}
	return repeated
}
```

### 2. Tests
**Files:**
- `3-iteration/repeat_test.go` (new)

**Description:**
- `package main`, table-driven `TestRepeat` following `TestAdd` pattern
- Cases: single char (`"a"` → `"aaaaa"`), multi-char (`"ab"` → `"ababababab"`), empty string (`""` → `""`)

```go
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
```