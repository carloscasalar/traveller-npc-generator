# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Build & Test Commands
- Build: `make build`
- Run: `make run`
- Test all: `make test`
- Test single package: `go test -v ./package/path`
- Test specific test: `go test -v ./package/path -run TestName`
- Lint: `make lint` (uses revive)
- All checks: `make check`

## Code Style Guidelines
- Format: Standard Go formatting (`gofmt`)
- Imports: Group stdlib first, then external dependencies, finally internal packages
- Testing: Use testify for assertions (`assert` and `require` packages)
- Error handling: Return explicit errors with descriptive messages
- Naming: Use Go idioms (CamelCase for exported, camelCase for unexported)
- Types: Prefer strong typing with defined constants for enums
- Builder pattern for constructing complex objects
- Package structure: cmd/, internal/, pkg/ for different visibility levels
- Documentation: Comment exported functions, types, and constants