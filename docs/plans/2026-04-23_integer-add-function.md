## Overview

Create a new `2-integer/` package with an `Add` function that sums two integers, following TDD: write the table-driven test first, then implement.

## Changes

### 1. Create test file
**Files:**
- `2-integer/integer_test.go` (new)

**Description:**
- Create a table-driven test covering positive, negative, and zero cases using the same patterns as `1-hello/hello_test.go`.

```
[]struct {
    a    int
    b    int
    want int
}{
    {1, 1, 2},
    {-1, -1, -2},
    {-1, 1, 0},
    {0, 0, 0},
    {5, -3, 2},
}
```

### 2. Implement Add function
**Files:**
- `2-integer/integer.go` (new)

**Description:**
- Package `main` with `func Add(a, b int) int` that returns `a + b`.