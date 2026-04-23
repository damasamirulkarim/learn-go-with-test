# Hello Function: Add Name Parameter with Default

## Requirements

- Change the `Hello` function signature from `Hello() string` to `Hello(name string) string` in `1-hello/hello.go`
- If `name` is an empty string (`""`), default to `"World"`
- Return format: `"Hello, Name"` (with comma and space)
- Update `main()` to call `Hello("World")` to preserve existing behavior
- Create `1-hello/hello_test.go` with table-driven tests covering:
  - Empty name → `"Hello, World"`
  - Non-empty name (e.g. `"Bob"`) → `"Hello, Bob"`
