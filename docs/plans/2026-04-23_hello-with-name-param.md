## Overview

Modify the existing `Hello` function to accept a `name` parameter with a default-to-"World" behavior. Following TDD: write tests first, then implement.

The current `Hello() string` returns `"Hello, world"` — the spec changes the signature to `Hello(name string) string`, defaults empty names to `"World"`, and uses the format `"Hello, Name"` (capitalized).

## Changes

### 1. Create test file
**Files:**
- `1-hello/hello_test.go` (new)

**Description:**
- Create a table-driven test covering the two required cases: empty name and non-empty name.

```
[]struct {
    name     string
    expected string
}{
    {"", "Hello, World"},
    {"Bob", "Hello, Bob"},
}
```

### 2. Update Hello function and main
**Files:**
- `1-hello/hello.go:5-7`

**Description:**
- Change signature from `Hello() string` to `Hello(name string) string`
- If `name == ""`, default to `"World"`
- Return `fmt.Sprintf("Hello, %s", name)` (or equivalent)
- Update `main()` at line 10 to call `Hello("World")`