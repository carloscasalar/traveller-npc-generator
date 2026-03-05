---
name: go-project-standards
description: "Write and refactor Go code with clean package boundaries, Make-driven build/test/lint workflows, and release hygiene. Use when adding core logic, CLI commands, importable library APIs, tests, changelog updates, or demo artifact refreshes."
argument-hint: "[task-or-feature]"
user-invocable: true
---

# Go Project Standards

## When to Use

- Implementing or refactoring Go features
- Adding or changing CLI entrypoints and command wiring
- Adding or changing importable library APIs
- Updating core domain logic and avoiding layer coupling
- Running quality checks before opening a PR
- Preparing release-ready changes with documentation and demos

## Architecture Rules

- Keep core logic in `internal/` and make it agnostic of delivery layers.
- Keep executable wiring and argument parsing in `cmd/`.
- Keep reusable importable APIs in `pkg/`.
- `internal/` code must not depend on `cmd/`.
- `pkg/` code must not depend on `cmd/`.
- Prefer dependency injection and interfaces at boundaries over hard-coded coupling.

## Implementation Workflow

1. Identify the target layer first (`internal/`, `pkg/`, or `cmd/`).
2. Implement behavior in core logic packages before wiring command or transport concerns.
3. Add or update tests close to changed code using `*_test.go` files.
4. Keep exported APIs intentional and stable; avoid leaking command-layer details into libraries.
5. If generated files are used in the project, regenerate them with the project tooling instead of hand-editing generated outputs.

## Build, Test, and Lint

- Prefer `make` targets when available, because projects often encode the canonical workflow there.
- Typical sequence:
  - `make build`
  - `make test`
  - `make lint`
  - `make check` (if present)
- If a target is missing, fall back to direct Go commands (`go build ./...`, `go test ./...`) and the configured linter.

## Release Checklist

1. Verify all checks pass (`make check` or equivalent lint + tests).
2. Update `CHANGELOG.md` with a new release section that follows the repository format.
3. Regenerate demo artifacts when the project tracks executable demos (for example GIFs, tapes, snapshots, or recorded outputs).
4. Ensure demo and changelog updates ship in the same change set as behavior updates.

## Output Expectations For The Agent

- Summarize what changed by layer: `internal`, `pkg`, `cmd`, tests, docs.
- Report the exact validation commands run and whether each passed.
- Explicitly call out whether changelog and demo artifacts were updated or why they were not needed.
- Ensure examples at the examples folder keeps working and modify them if needed.
- When adding new features, discuss if it will need a new example to showcase it.
