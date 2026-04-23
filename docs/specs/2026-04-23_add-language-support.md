# Add Language Support to Hello Function

## Requirements

- Add a `language` parameter as the second argument to `Hello(name string, language string) string`
- Default to English when `language` is empty string (`""`)
- Supported languages:
  - English (default) → `"Hello, %s"`
  - Indonesian → `"Halo, %s"`
  - Spanish → `"Hola, %s"`
- If `name` is empty, default to `"World"` (existing behavior preserved)
- Update tests to cover all three languages plus the empty/default language case