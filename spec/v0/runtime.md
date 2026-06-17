# Runtime Model v0

- Execution is single-pass over statements
- Variables are stored in a global environment (map[string]int)
- There is no scope (everything is global)
- There is no heap, no stack model yet
- Program ends after last statement