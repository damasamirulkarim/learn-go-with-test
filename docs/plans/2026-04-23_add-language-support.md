## Overview

Add a `language` parameter to `Hello(name string, language string) string` so it can return greetings in English, Indonesian, and Spanish. Defaults to English when `language` is empty. Following TDD: update tests first, then implement.

## Changes

### 1. Update tests for language support
**Files:**
- `1-hello/hello_test.go:1-21`

**Description:**
- Add `language` field to the table-driven test struct
- Cover all three languages plus the empty/default case

```
[]struct {
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
```

### 2. Update Hello function
**Files:**
- `1-hello/hello.go:5-9`

**Description:**
- Change signature to `Hello(name string, language string) string`
- Add a `greeting` variable that selects the format string based on `language`
- Use a switch on `language`: `"Indonesian"` → `"Halo, %s"`, `"Spanish"` → `"Hola, %s"`, default (including `""`) → `"Hello, %s"`
- Preserve existing empty-name default: if `name == ""`, default to `"World"`
- Return `fmt.Sprintf(greeting, name)`

### 3. Update main
**Files:**
- `1-hello/hello.go:13`

**Description:**
- Update `main()` to call `Hello("World", "")`