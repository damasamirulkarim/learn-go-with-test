## Overview
Rename the Go source files in `2-integer/` from `integer.go`/`integer_test.go` to `adder.go`/`adder_test.go`. No content changes are required since the package name (`main`) and function names (`Add`, `TestAdd`) remain unchanged.

## Changes

### 1. Rename integer.go to adder.go
**Files:**
- `2-integer/integer.go` → `2-integer/adder.go`

**Description:**
- Rename `2-integer/integer.go` to `2-integer/adder.go`
- No content changes; file contents remain identical:

```go
package main

func Add(a, b int) int {
	return a + b
}
```

### 2. Rename integer_test.go to adder_test.go
**Files:**
- `2-integer/integer_test.go` → `2-integer/adder_test.go`

**Description:**
- Rename `2-integer/integer_test.go` to `2-integer/adder_test.go`
- No content changes; file contents remain identical:

```go
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
```

### 3. Verify tests pass after rename
- Run `go test ./2-integer/...` to confirm all tests still pass