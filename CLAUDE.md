# abstract-account — CLAUDE.md

Cosmos module providing abstract account functionality for Xion. Used as a dependency by the main `xion` chain repo.

## GitHub Workflows

### `go.yml`

**Triggered by:** Push (any branch), `workflow_dispatch`

Runs Go tests, linting, and coverage checks.

### `rust.yml`

**Triggered by:** Push (any branch)

Runs Rust check, tests, and Clippy lints.

## Upstream Triggers

None — this repo is not triggered by other repos.

## Downstream Triggers

None — changes here are consumed by `burnt-labs/xion` as a Go module dependency.

## Development

```bash
# Go
go test ./...
go vet ./...

# Rust
cargo check
cargo test
cargo clippy
```

## Note

This module is imported by `burnt-labs/xion` in `go.mod`. When making breaking changes, coordinate with the xion repo.
