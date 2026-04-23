# AGENTS.md

## Project

Learn Go with Test - a Go learning repository structured by chapters. Each chapter directory contains implementation (`*.go`) and test (`*_test.go`) files.

## Commands

```bash
go test ./...        # Run all tests
go test ./1-hello   # Run tests for specific chapter
go test ./2-integer # Run tests for specific chapter
go build ./...      # Build packages (fails without main)
```

## Structure

- `1-hello/` - Chapter 1: Hello World
- `2-integer/` - Chapter 2: Integers
- No `main.go` - this is a library, not an executable

## Notes

- Go 1.26.1
- Follows "Learn Go with Test" conventions